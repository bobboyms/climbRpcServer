package datatransfer

type TopicProducerRequest struct {
	TopicName string `json:",topicName"`
	TopicMessage string `json:",topicMessage"`
}

type TopicProducerResponse struct {
	Id         string `json:",id"`
	StatusCode uint   `json:",statusCode"`
	Message    string `json:",statusCode"`
}

type TopicGetRequest struct {
	TopicName string `json:",topicName"`
}

type TopicGetResponse struct {
	Id string `json:",id"`
	Message string `json:",message"`
	StatusCode uint `json:",statusCode"`
}