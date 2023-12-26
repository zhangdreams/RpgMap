package common

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 保存个goroutine的input、output channel
// 用于向各个goroutine发送消息
var chanMap sync.Map

// Pid 线程pid，暂时只使用两个通道，后续可根据需求添加属性
// 这里暂时只考虑了单节点的情况，多节点需要在pid中附带节点信息
type Pid struct {
	In  chan interface{}
	Out chan interface{}
}

var callIndex int64 = 1

type CallInfo struct {
	ID      int64            // call Msg 的唯一ID，需要跟CallBack消息对比这个id
	RetChan chan interface{} // 消息返回的通道
	Msg     interface{}      // 发送的消息
}
type CallBack struct {
	ID     int64
	Result interface{}
}

func Register(name string, Pid *Pid) {
	chanMap.Store(name, Pid)
}

func UnRegister(name string) {
	chanMap.Delete(name)
}

func WhereIs(name string) *Pid {
	if pid, ok := chanMap.Load(name); ok {
		return pid.(*Pid)
	}
	return nil
}

// Pids 返回所有的线程chan
func Pids() map[interface{}]interface{} {
	maps := make(map[interface{}]interface{})
	chanMap.Range(func(key, value interface{}) bool {
		maps[key] = value
		return true
	})
	return maps
}

// Cast 单节点的cast
func Cast(pid *Pid, msg interface{}) {
	pid.In <- msg
}

// Call 单节点的call
func Call(pid *Pid, msg interface{}) (bool, interface{}) {
	return CallTimeOut(pid, msg, 5)
}
func CallName(name string, msg interface{}) (bool, interface{}) {
	if pid := WhereIs(name); pid != nil {
		return CallTimeOut(pid, msg, 5)
	}
	fmt.Sprintln("not found pid ", name)
	return false, errors.New(fmt.Sprintf("not found pid %s", name))
}
func CallTimeOut(pid *Pid, msg interface{}, timerOut time.Duration) (bool, interface{}) {
	defer func() { recover() }()
	retChan := make(chan interface{})
	defer close(retChan)
	reqID := makeReqID()
	// 发送消息
	pid.In <- &CallInfo{ID: reqID, RetChan: retChan, Msg: msg}

	// 创建timeout定时器
	t := time.NewTimer(timerOut)

	// 等待消息返回
rec:
	select {
	case result := <-retChan:
		r := result.(*CallBack)
		if r.ID == reqID {
			t.Stop()
			return true, r.Result
		}
		fmt.Sprintln("not match callResult ID", r.ID, reqID)
		goto rec
	case <-t.C:
		fmt.Println("rec callResult timeout")
		return false, errors.New("call time out")
	}
}

func makeReqID() int64 {
	return atomic.AddInt64(&callIndex, 1)
}
