package main

import (
	"fmt"
	"go_learning/logagent/config"
	"go_learning/logagent/etcd"
	"go_learning/logagent/kafka"
	"go_learning/logagent/taillog"
	"go_learning/logagent/utils"
	"sync"

	"gopkg.in/ini.v1"

	//"go_learning/logagent/taillog"
	"time"
)

var (
	cfg = new(config.AppConfig)
)

// logagent入口
func main() {
	// 加载配置文件
	err := ini.MapTo(cfg, "./config/config.ini")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 初始化kafka连接
	err = kafka.Init([]string{cfg.KafkaConfig.Address}, cfg.ChanMaxSize)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 初始化etcd
	err = etcd.Init(cfg.EtcdConfig.Address, time.Duration(cfg.EtcdConfig.Timeout)*time.Second)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("etc init success")
	ipStr, err := utils.GetOutboundIP()
	if err != nil {
		panic(err)
	}
	etcdConfKey := fmt.Sprintf(cfg.EtcdConfig.Key, ipStr)
	// 从etcd获取日志收集项的配置
	logEntryConf, err := etcd.GetConf(etcdConfKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, value := range logEntryConf {
		fmt.Println(value)
	}
	// 打开日志文件收集日志
	taillog.Init(logEntryConf)
	// 设置哨兵
	var newConfChan = taillog.NewConfChan()
	var wg sync.WaitGroup
	wg.Add(1)
	go etcd.WatchConf(cfg.EtcdConfig.Key, newConfChan)
	wg.Wait()
	/* err = taillog.Init(cfg.FileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	run() */

}

/* func run() {
	// 读取日志
	for {
		select {
		case line := <-taillog.ReadChan():
			// 发送到kafka
			kafka.SendToKafka("web_log", line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}*/
