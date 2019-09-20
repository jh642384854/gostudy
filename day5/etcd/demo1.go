package main

import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

/**
	etcd连接测试
	需要注意的是，如果远程连接这个服务，需要开启这个端口。或是在测试的时候，直接把防火墙关闭。这样就不会出现连接不上的问题
 */

func main() {
	var (
		config clientv3.Config
		client *clientv3.Client
		err    error
	)
	//客户端链接配置
	config = clientv3.Config{
		Endpoints:   []string{"192.168.20.99:2379"}, //集群列表
		DialTimeout: 5 * time.Second,                //建立连接的超时时间
	}

	//建立链接
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("connect successful")
	client = client
}
