package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

/**
	通过前缀的方式来查找指定的key
 */

func main() {
	var (
		config  clientv3.Config
		client  *clientv3.Client
		err     error
		kv      clientv3.KV
		getResp *clientv3.GetResponse
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
	if getResp, err = kv.Get(context.TODO(), "/crontab/jobs", clientv3.WithPrefix()); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("get result")
		fmt.Println(getResp.Kvs)
		//for range来遍历多个值
		for _, etcdVal := range getResp.Kvs {
			fmt.Println(string(etcdVal.Key), "：", string(etcdVal.Value))
		}
	}
}
