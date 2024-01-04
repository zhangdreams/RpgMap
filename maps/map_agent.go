package maps

import (
	"errors"
	"fmt"
	"rpgMap/common"
	"rpgMap/global"
)

func (state *MapState) mapHandleMsg(pid *common.Pid) {
	for {
		select {
		case data, ok := <-pid.In:
			if !ok {
				return
			}
			// 处理接收到的数据
			switch data.(type) {
			case *common.CallInfo:
				state.doMapHandleCall(data.(*common.CallInfo))
			default:
				state.mapHandle(data)
			}
		}
	}
}

func (state *MapState) doMapHandleCall(callInfo *common.CallInfo) {
	ret, err := state.mapHandleCall(callInfo.Msg)
	if err != nil {
		ret = err
	}
	backMsg := common.CallBack{ID: callInfo.ID, Result: ret}
	callInfo.RetChan <- &backMsg
}

func (state *MapState) mapHandleCall(data interface{}) (ret interface{}, err error) {
	switch data.(type) {

	default:
		err = errors.New(fmt.Sprint(state.Name, "unhandled msg", data))
	}
	return
}

func (state *MapState) mapHandle(data interface{}) {
	switch data.(type) {
	case global.Exit:
		close(state.Pid.In)
		close(state.Pid.Out)
	case global.Loop:
		state.MapLoop()
	case global.ModHandle:
		state.Mod.Handle(state, data.(global.ModHandle).Msg)
	default:
		fmt.Println(state.Name, "unhandled msg", data)
	}
}
