package unittest

import (
	"message/managerfilesystem"
	"message/modeldatatransfer"
	"message/topicservice"
	"testing"
)

func TestCreateTopic(t *testing.T) {

	var fileSystem managerfilesystem.FileSystem = new(managerfilesystem.FileSystemGobEncoder)

	var manager topicservice.TopicService = new(topicservice.TopicServiceImp).Init(fileSystem)

	response := manager.TopicCreate(modeldatatransfer.TopicProducerRequest{
		"baixarEstoque", "{'id':10,'name':'thiago'}",
	})

	if &response == nil || response.StatusCode == 500 {
		t.Error("Erro create topic: " + response.Message)
	}
}

//func TestCreateFile(t *testing.T) {
//
//	var fs managerfilesystem.FileSystem = new(managerfilesystem.FileSystemGobEncoder)
//
//	fs.CreateFile(fileName, folderName, filesystem.MessageFileStruct{
//		"tes",
//		"sdsdsd",
//	})
//
//	if _, err := os.Stat(folderName + "/" + fileName); err != nil {
//		t.Error("File dont created")
//	}
//}
//
//func TestOpenFile(t *testing.T)  {
//	var fs managerfilesystem.FileSystem = new(managerfilesystem.FileSystemGobEncoder)
//
//	var m filesystem.MessageFileStruct
//	fs.OpenFile(folderName, fileName, &m)
//
//	if &m == nil {
//		t.Error("File dont open")
//	}
//
//}
//
//func TestDeleteFile(t *testing.T)  {
//
//	var fs managerfilesystem.FileSystem = new(managerfilesystem.FileSystemGobEncoder)
//	fs.DeleteFile(folderName, fileName)
//
//}
