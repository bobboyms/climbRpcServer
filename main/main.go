package main

import (
	"message/datatransfer"
	"message/managerfilesystem"
	"message/topicservice"
)

func main() {

	var fileSystem managerfilesystem.FileSystem = new(managerfilesystem.FileSystemBinaryEncoder)

	var manager topicservice.TopicService = new(topicservice.TopicServiceImp).Init(fileSystem)

	manager.TopicCreate(datatransfer.TopicProducerRequest{
		"baixarEstoque", "",
	})

	println(manager)

}
