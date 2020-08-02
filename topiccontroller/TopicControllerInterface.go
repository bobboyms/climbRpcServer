package topiccontroller

import "message/datatransfer"

type TopicController interface {
	TopicCreate(request datatransfer.TopicProducerRequest) datatransfer.TopicProducerResponse
	TopicGet(request datatransfer.TopicGetRequest) datatransfer.TopicGetResponse
}
