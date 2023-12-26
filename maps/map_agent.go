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
			fmt.Println(MapServer, "handleMsg", data)
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
		err = errors.New(fmt.Sprint(MapServer, "unhanle msg", data))
	}
	return
}

func (state *MapState) mapHandle(data interface{}) {
	switch data.(type) {
	case global.Exit:
		if server := common.WhereIs(MapServer); server != nil {
			common.UnRegister(MapServer)
			close(server.In)
			close(server.Out)
		}
	default:
		fmt.Println(MapServer, "unhanle msg", data)
	}
}
