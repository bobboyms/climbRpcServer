package topicservice

import (
	"message/managerfilesystem"
	"message/managerindex"
	"message/model"
	"message/modeldatatransfer"
	"strings"
)

const (
	StatusOk            = 200
	StatusInternalError = 500
	StatusNotFound      = 404
)

const (
	IndexFileName = "INDEX.m"
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

	var index managerindex.ManagerIndex = new(managerindex.ManagerIndexImp).
		Init(dir, m.fileSystem)

	index.CreateIndex(IndexFileName)

	fileName, err := m.createTopicFile(dir, request.TopicName)

	index.AddItem(fileName, IndexFileName)

	if err != nil {
		return m.getError(err.Error())
	}

	return modeldatatransfer.TopicProducerResponse{
		Id:         fileName,
		StatusCode: StatusOk,
		Message:    "topic created successfully",
	}

}

func (m *TopicServiceImp) TopicGet(request modeldatatransfer.TopicGetRequest) modeldatatransfer.TopicGetResponse {

	dir := request.TopicName

	if !m.fileSystem.IsFileExist(dir, IndexFileName) {
		return modeldatatransfer.TopicGetResponse{
			Id:         "0",
			Message:    "ITEM MESSAGE NOT FOUND",
			StatusCode: StatusNotFound,
		}
	}

	var index managerindex.ManagerIndex = new(managerindex.ManagerIndexImp).
		Init(dir, m.fileSystem)

	item, err := index.GetFirstIndex(IndexFileName)

	if err != nil {
		return modeldatatransfer.TopicGetResponse{
			Id:         "0",
			Message:    "ITEM MESSAGE NOT FOUND",
			StatusCode: StatusNotFound,
		}
	}

	fileName, _ := index.GetItem(item, IndexFileName)

	var mf model.MessageFileStruct
	m.fileSystem.OpenFile(dir, fileName+managerfilesystem.FileExtension, &mf)

	defer func() {
		index.RemoveItem(item, IndexFileName)
		m.fileSystem.DeleteFile(dir, fileName+managerfilesystem.FileExtension)
	}()

	return modeldatatransfer.TopicGetResponse{
		Id:         mf.Id,
		Message:    mf.Message,
		StatusCode: StatusOk,
	}

}

func (m *TopicServiceImp) getError(message string) modeldatatransfer.TopicProducerResponse {
	return modeldatatransfer.TopicProducerResponse{
		Id:         "0",
		StatusCode: StatusInternalError,
		Message:    message,
	}
}
