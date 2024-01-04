package maps

import (
	"errors"
	"fmt"
	"rpgMap/common"
	"rpgMap/config"
	"rpgMap/global"
)

const MapServer = "map_server"

var ModMaps = make(map[string]*MapMod)

type MapData struct {
	ID      int32
	Pid     *common.Pid
	MapName string
	Line    int32
	RoleNum int32
}

var MapDic = make(map[int32]*[]MapData)
var MapPidDic = make(map[string]*common.Pid)

func StartMapServer() {
	input := make(chan interface{}, 1000)
	output := make(chan interface{})
	pid := common.Pid{In: input, Out: output}
	common.Register(MapServer, &pid) // 注册到线程通道map内
	go InitMgr(&pid)
}

func InitMgr(pid *common.Pid) {
	// 接受消息
	handleMsg(pid)
}

func handleMsg(pid *common.Pid) {
	for {
		select {
		case data, ok := <-pid.In:

			if !ok {
				return
			}
			// 处理接收到的数据
			switch data.(type) {
			case *common.CallInfo:
				doHandleCall(data.(*common.CallInfo))
			default:
				handle(data)
			}
		}
	}
}

func doHandleCall(callInfo *common.CallInfo) {
	ret, err := handleCall(callInfo.Msg)
	if err != nil {
		ret = err
	}
	backMsg := common.CallBack{ID: callInfo.ID, Result: ret}
	callInfo.RetChan <- &backMsg
}

func handleCall(data interface{}) (ret interface{}, err error) {
	switch data.(type) {
	case global.CreateMap:
		m := data.(global.CreateMap)
		ret = createMap(m.ID, m.Name, m.Line, m.ModName)
	default:
		err = errors.New(fmt.Sprint(MapServer, "unhandled msg", data))
	}
	return
}

func handle(data interface{}) {
	switch data.(type) {
	case global.Exit:
		fmt.Println(MapServer, " receive exit")
		if server := common.WhereIs(MapServer); server != nil {
			close(server.In)
			close(server.Out)
			common.UnRegister(MapServer)
		}
	default:
		fmt.Println(MapServer, "unhandled msg", data)
	}
}

// getMap 返回一个可进入的map
func getMap(id int32) *common.Pid {
	maps, ok := MapDic[id]
	if ok {
		conf, err := config.GetMapConfig(id)
		if err != nil {
			return nil
		}
		for _, m := range *maps {
			if m.RoleNum < conf.MaxNum {
				return m.Pid
			}
		}
	}
	conf, err := config.GetMapConfig(id)
	if err != nil {
		return nil
	}
	return createMap(id, conf.Name, 0, "")
}

// getMapName 根据地图名返回map Pid
func getMapName(name string) *common.Pid {
	m, ok := MapPidDic[name]
	if ok {
		return m
	}
	return nil
}

// createMap 创建地图
func createMap(id int32, name string, line int32, modName string) *common.Pid {
	if getMapName(name) != nil {
		fmt.Println("exist Map MapID:", id, " name:", name)
		return nil
	}
	fmt.Println("mapserver create map ", id, name, line)
	if len(modName) == 0 {
		modName = "mod_common"
	}
	mapPid := StartMap(id, name, line, ModMaps[modName])
	MapPidDic[name] = mapPid
	maps, ok := MapDic[id]
	if ok {
		*maps = append(*maps, MapData{id, mapPid, name, line, 0})
	}
	return mapPid
}

func ShowAllMap() {
	fmt.Println("show all map Num:", len(MapPidDic))
	for k, v := range MapPidDic {
		fmt.Println("show map :", k, " ", v)
	}
}
