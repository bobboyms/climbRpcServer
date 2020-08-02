package main

import (
	"message/managerfilesystem"
	"message/modeldatatransfer"
	"message/topicservice"
)

func main() {

	var fileSystem managerfilesystem.FileSystem = new(managerfilesystem.FileSystemBinaryEncoder)

	var manager topicservice.TopicService = new(topicservice.TopicServiceImp).Init(fileSystem)

	manager.TopicCreate(modeldatatransfer.TopicProducerRequest{
		"baixarEstoque", "",
	})

	println(manager)

}
