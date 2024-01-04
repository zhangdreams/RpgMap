package global

// / 消息匹配

type Exit struct{}

type Ok struct{}

type CreateMap struct {
	ID      int32
	Name    string
	Line    int32
	ModName string
}

type ModHandle struct {
	Msg interface{}
}

type Loop struct{}
