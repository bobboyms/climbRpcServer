package main

import (
	"message/datatransfer"
	"message/managerfilesystem"
	"message/topiccontroller"
)

func main() {

	var fileSystem managerfilesystem.FileSystem =
		new(managerfilesystem.FileSystemBinaryEncoder)

	var manager topiccontroller.TopicController =
		new(topiccontroller.TopicControllerImp).Init(fileSystem)

	manager.TopicCreate(datatransfer.TopicProducerRequest{
		"baixarEstoque","",
	})

	println(manager)

}
