package main

import (
	"bufio"
	"fmt"
	"os"
	"rpgMap/common"
	"rpgMap/config"
	"rpgMap/global"
	"rpgMap/maps"
	"strings"
	"time"
)

func main() {
	//args := os.Args
	//fmt.Println(args, len(args))
	var rpgMapChan = make(chan int32, 1)
	go stopCheck(&rpgMapChan)
	config.InitConfig()   //加载配置
	common.StartTimer()   //启动定时器线程
	maps.StartMapServer() // 地图管理进程

	ids := config.GetMapIDs()
	for _, mid := range ids {
		name := common.GetMapName(mid)
		_, ret := common.CallName(maps.MapServer, global.CreateMap{ID: mid, Name: name})
		fmt.Println("create map result,", ret)
	}

	time.Sleep(time.Second) // 等待1s再输出
	maps.ShowAllMap()
	//todo

	// 保持在线
	<-rpgMapChan
	stopAll()
	time.Sleep(time.Second * 5)
	fmt.Println("main stopped")
}

func stopCheck(exitChan *chan int32) {
	fmt.Println("Press 'c' and Enter to exit.")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		if input == "c" {
			close(*exitChan)
			return
		}
	}
}

func stopAll() {
	for _, Pid := range common.Pids() {
		Pid.(*common.Pid).In <- global.Exit{}
	}
}
