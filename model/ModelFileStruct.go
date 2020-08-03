package model

type MessageFileStruct struct {
	Id      string
	Message string
}

type IndexFileStruct struct {
	Count uint16
	Itens map[uint16]string
}
