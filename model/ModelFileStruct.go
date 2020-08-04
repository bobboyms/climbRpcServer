package model

type MessageFileStruct struct {
	Id      string
	Message string
}

type IndexFileStruct struct {
	Count uint16
	Itens map[uint16]string
}

type RpcFileStruct struct {
	Id         string
	ClassName  string
	MethodName string
	Args       []RpcArgs
}

type RpcArgs struct {
	Name  string
	Type  string
	Value string
}

type RpcProcessed struct {
	Id         string
	ChanelName string
	Result     string
}
