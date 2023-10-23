package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	for true {
		cli, err := clientv3.New(clientv3.Config{
			Endpoints:   []string{"etcd.etcd.svc.cluster.local:2379"},
			DialTimeout: 5 * time.Second,
			Username:    "root",
			Password:    "pJs3YaGdp4",
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		defer cli.Close()

		// 存储键值对数据
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		_, err = cli.Put(ctx, "mykey", "hello world")
		cancel()
		if err != nil {
			fmt.Println(err)
			return
		}

		// 获取键值对数据
		ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
		resp, err := cli.Get(ctx, "mykey")
		cancel()
		if err != nil {
			fmt.Println(err)
			return
		}

		// 输出获取到的值
		for _, ev := range resp.Kvs {
			fmt.Printf("%s : %s\n", ev.Key, ev.Value)
		}

		time.Sleep(time.Second * 5)
	}
}
