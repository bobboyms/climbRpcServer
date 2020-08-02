package topicservice

import (
	"message/managerfilesystem"
	"message/model"
	"message/modeldatatransfer"
	"strings"
)

const (
	STATUS_OK             = 200
	STATUS_INTERNAL_ERROR = 500
)

type TopicServiceImp struct {
	fileSystem managerfilesystem.FileSystem
}

func (m *TopicServiceImp) Init(filesystem managerfilesystem.FileSystem) *TopicServiceImp {
	m.fileSystem = filesystem
	return m
}

func (m *TopicServiceImp) createDir(topicName string) string {

	return m.fileSystem.CreateDirectory(strings.ToLower(topicName))

}

func (m *TopicServiceImp) createTopicFile(dir, message string) (string, error) {

	idFileName := managerfilesystem.GetUuid()

	return idFileName, m.fileSystem.CreateFile(idFileName+managerfilesystem.FileExtension, dir, model.MessageFileStruct{
		Id:      idFileName,
		Message: message,
	})
}

func (m *TopicServiceImp) TopicCreate(request modeldatatransfer.TopicProducerRequest) modeldatatransfer.TopicProducerResponse {

	dir := m.createDir(request.TopicName)

	id, err := m.createTopicFile(dir, request.TopicName)

	if err != nil {
		return m.getError(err.Error())
	}

	return modeldatatransfer.TopicProducerResponse{
		Id:         id,
		StatusCode: STATUS_OK,
		Message:    "topic created successfully",
	}

}

func (m *TopicServiceImp) TopicGet(request modeldatatransfer.TopicGetRequest) modeldatatransfer.TopicGetResponse {
	panic("implement me")
}

func (m *TopicServiceImp) getError(message string) modeldatatransfer.TopicProducerResponse {
	return modeldatatransfer.TopicProducerResponse{
		Id:         "0",
		StatusCode: STATUS_INTERNAL_ERROR,
		Message:    message,
	}
}
