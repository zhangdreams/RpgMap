package common

import (
	"fmt"
	"rpgMap/global"
	"time"
)

// 当前为给其他线程注册轮训的线程。
// 注册后会定期往对应的线程发送 "loop" 消息
// RegisterTimer / UnRegisterTimer

type timerSet struct {
	Pid      *Pid  // 用来向各个注册的goroutine发送消息
	interval int   // 间隔时长
	lastTime int64 // 上次轮训时间
}

var timerMaps = make(map[string]timerSet)

const TimerServer = "timer_server"

type timerOp struct {
	op       string // 操作 add or del
	name     string // 线程名
	Pid      *Pid   // 线程Pid
	interval int    // 间隔
}

func StartTimer() {
	go Init()
}

func Init() {
	// 创建两个通道
	input := make(chan interface{}, 1000)
	output := make(chan interface{})
	Register(TimerServer, &Pid{In: input, Out: output}) // 注册到线程通道map内

	// 设置定时器
	pollInterval := time.Millisecond * 50 // 通常最小定时器间隔为100ms，为了减小误差和性能考虑这里定时器设置50ms
	pollTicker := time.NewTicker(pollInterval)
	defer pollTicker.Stop()

	for {
		select {
		case data, ok := <-input:
			if !ok {
				return
			}
			// 处理接收到的数据
			doEvent(data)
		case <-pollTicker.C:
			// 执行轮询逻辑
			doLoop()
		}
	}
}

func doEvent(event interface{}) interface{} {
	switch event.(type) {
	case global.Exit: // 通知退出
		doClose()
		return false
	case timerOp:
		doTimerOp(event.(timerOp))
	default:
		fmt.Println(TimerServer, "unhandle event", event)
	}
	return nil
}

func doClose() {
	fmt.Println(TimerServer, " receive exit")
	if server := WhereIs(TimerServer); server != nil {
		UnRegister(TimerServer)
		close(server.In)
		close(server.Out)
	}
}

func doTimerOp(op timerOp) {
	if op.op == "add" {
		timerMaps[op.name] = timerSet{Pid: op.Pid, interval: op.interval}
	} else if op.op == "del" {
		delete(timerMaps, op.name)
	}
}

func doLoop() {
	Now2 := Now2()
	for _, timer := range timerMaps {
		if Now2 >= timer.lastTime+int64(timer.interval) {
			timer.Pid.In <- "Loop"
		}
	}
}

// RegisterTimer 注册一个timer轮训
func RegisterTimer(name string, Pid *Pid, interval int) {
	if server := WhereIs(TimerServer); server != nil {
		server.In <- timerOp{op: "add", name: name, Pid: Pid, interval: interval}
	}
}

// UnRegisterTimer 取消一个timer轮训
func UnRegisterTimer(name string) {
	if server := WhereIs(TimerServer); server != nil {
		server.In <- timerOp{op: "del", name: name}
	}
}
