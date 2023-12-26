package maps

import (
	"fmt"
	"rpgMap/common"
)

func StartMap(args ...interface{}) *common.Pid {
	input := make(chan interface{}, 1000)
	output := make(chan interface{})
	pid := common.Pid{In: input, Out: output}
	go InitMap(&pid, args)
	return &pid
}

func InitMap(pid *common.Pid, args []interface{}) {
	name := args[1].(string)
	common.Register(name, pid) // 注册到线程通道map内
	common.RegisterTimer(name, pid, 100)
	fmt.Println("init map done", name)
}
