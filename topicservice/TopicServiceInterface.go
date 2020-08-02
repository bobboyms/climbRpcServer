package topicservice

import "message/modeldatatransfer"

type TopicService interface {
	TopicCreate(request modeldatatransfer.TopicProducerRequest) modeldatatransfer.TopicProducerResponse
	TopicGet(request modeldatatransfer.TopicGetRequest) modeldatatransfer.TopicGetResponse
}
