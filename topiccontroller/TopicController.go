package topiccontroller

import (
	"message/datatransfer"
	"message/managerfilesystem"
	"message/model"
	"strings"
)

const (
	STATUS_OK = 200
	STATUS_INTERNAL_ERROR = 500
)

type TopicControllerImp struct {

	fileSystem managerfilesystem.FileSystem

}

func (m *TopicControllerImp) Init(filesystem managerfilesystem.FileSystem) *TopicControllerImp {
	m.fileSystem = filesystem
	return m
}

func (m *TopicControllerImp) createDir(topicName string) (string)  {

	return m.fileSystem.CreateDirectory(strings.ToLower(topicName))

}

func (m *TopicControllerImp) createTopicFile(dir, message string) (string, error) {

	idFileName := managerfilesystem.GetUuid()

	return idFileName, m.fileSystem.CreateFile(idFileName + managerfilesystem.FileExtension, dir, model.MessageFileStruct{
		Id: idFileName,
		Message: message,
	})
}

func (m *TopicControllerImp) TopicCreate(request datatransfer.TopicProducerRequest) datatransfer.TopicProducerResponse {

	dir := m.createDir(request.TopicName)

	id, err := m.createTopicFile(dir, request.TopicName)

	if err != nil {
		return m.getError(err.Error())
	}

	return datatransfer.TopicProducerResponse{
		Id: id,
		StatusCode: STATUS_OK,
		Message: "topic created successfully",
	}

}

func (m *TopicControllerImp) TopicGet(request datatransfer.TopicGetRequest) datatransfer.TopicGetResponse {
	panic("implement me")
}

func (m *TopicControllerImp) getError(message string) datatransfer.TopicProducerResponse {
	return datatransfer.TopicProducerResponse {
		Id: "0",
		StatusCode: STATUS_INTERNAL_ERROR,
		Message: message,
	}
}
