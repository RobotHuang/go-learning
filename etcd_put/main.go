package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

// etcd client put/get demo
// use etcd/clientv3
// 在阿里云上搭建etcd要用私网ip
func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"118.31.64.83:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
    fmt.Println("connect to etcd success")
	defer cli.Close()
	// put
	ctx, cancel := context.WithTimeout(context.Background(), 5* time.Second)
	//value := `[{"path":"/tmp/nginx.log","topic":"web_log"},{"path":"/tmp/redis.log","topic":"redis_log"}]`
	value := `[{"path":"/tmp/nginx.log","topic":"web_log"}]`
	_, err = cli.Put(ctx, "/logagent/collect_config", value)
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}
	/* // get
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "foo")
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	} */
}