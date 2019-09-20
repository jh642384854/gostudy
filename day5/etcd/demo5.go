package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
	"time"
)

/**
	删除key
 */
func main() {
	var (
		config      clientv3.Config
		client      *clientv3.Client
		err         error
		kv          clientv3.KV
		delRes      *clientv3.DeleteResponse
		delKeyVlaue *mvccpb.KeyValue
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
	if delRes, err = kv.Delete(context.TODO(), "/crontab/jobs/job2", clientv3.WithPrevKV()); err != nil {
		fmt.Println(err)
		return
	}
	//判断确实有被删除的key
	if len(delRes.PrevKvs) > 0 {
		//得到被删除的key的对象信息
		for _, delKeyVlaue = range delRes.PrevKvs {
			fmt.Println(string(delKeyVlaue.Key), ":", string(delKeyVlaue.Value))
		}
	}
}
