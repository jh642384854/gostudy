package main

import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"golang.org/x/net/context"
	"time"
)

/**
	key的自动过期和租约机制。
	lease的应用
 */

func main() {
	var (
		config        clientv3.Config
		client        *clientv3.Client
		err           error
		kv            clientv3.KV
		lease         clientv3.Lease
		leaseRes      *clientv3.LeaseGrantResponse
		leaseID       clientv3.LeaseID
		putRes        *clientv3.PutResponse
		getRes        *clientv3.GetResponse
		leaseKeepChan <-chan *clientv3.LeaseKeepAliveResponse
		leaseKeepRes  *clientv3.LeaseKeepAliveResponse
	)
	//客户端链接配置。
	config = clientv3.Config{
		Endpoints:   []string{"192.168.20.99:2379"}, //集群列表
		DialTimeout: 5 * time.Second,                //建立连接的超时时间
	}
	//建立链接
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

	//下面的来操作KV。首先是利用client(客户端链接)来生成一个clientv3.KV对象，通过这个对象来进行KV操作   ,clientv3.WithPrevKV()
	kv = clientv3.NewKV(client)

	/**
		下面的逻辑是这样的：
		首先在创建一个key的时候，我们为这个key绑定一个过期时间的设置
		然后我们有为这个过期时间专门做了自动续期的处理。
	 */
	// 申请一个租约(相当于就是租赁一个房子)
	lease = clientv3.NewLease(client)
	//给这个房子设置租约时间为10秒.Grant()函数的第二个参数就是时间单位，单位是秒
	if leaseRes, err = lease.Grant(context.TODO(), 10); err != nil {
		fmt.Println(err)
	}
	//拿到租约ID(相当于就是合同编号)
	leaseID = leaseRes.ID

	//设置自动续租功能
	if leaseKeepChan, err = lease.KeepAlive(context.TODO(), leaseID); err != nil {
		fmt.Println(err)
	}
	//处理自动续租的逻辑
	go func() {
		for {
			select {
			case leaseKeepRes = <-leaseKeepChan:
				if leaseKeepRes == nil {
					fmt.Println("租约已经失效了")
					goto END
				} else {
					fmt.Println("自动续租成功，续租id：", leaseKeepRes.ID)
				}
			}
		}
	END:
	}()
	//然后在创建key的时候，就带上这个租约信息
	if putRes, err = kv.Put(context.TODO(), "/crontab/lock/job1", "{...lock job1}", clientv3.WithLease(leaseID)); err != nil {
		fmt.Println(err)
	}
	fmt.Println("set key Revision：", putRes.Header.Revision)

	//读取设置的key
	for {
		if getRes, err = kv.Get(context.TODO(), "/crontab/lock/job1"); err != nil {
			fmt.Println(err)
		}
		if getRes.Count == 0 {
			fmt.Println("租约到期了")
			break
		}
		fmt.Println(getRes.Kvs)
		time.Sleep(2 * time.Second)
	}
}
