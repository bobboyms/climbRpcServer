package modeldatatransfer

type RpcProducerRequest struct {
	ChanelName string    `json:"chanelName"`
	ClassName  string    `json:"className"`
	MethodName string    `json:"methodName"`
	Args       []RpcArgs `json:"args"`
}

type RpcArgs struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type RpcProducerResponse struct {
	Id         string `json:",id"`
	StatusCode uint   `json:",statusCode"`
	Message    string `json:",statusCode"`
}

type RpcGetRequest struct {
	ChanelName string `json:",chanelName"`
}

type RpcGetResponse struct {
	Id         string    `json:",chanelName"`
	StatusCode uint16    `json:",chanelName"`
	ClassName  string    `json:",className"`
	MethodName string    `json:",methodName"`
	Args       []RpcArgs `json:",rpcArgs"`
}

type RpcProcessedRequest struct {
	Id         string `json:"id"`
	ChanelName string `json:"chanelName"`
	Result     string `json:"result"`
}

type RpcProcessedResponse struct {
	Id         string `json:",chanelName"`
	StatusCode uint16 `json:",chanelName"`
}

type RpcGetProcessedRequest struct {
	Id         string `json:"id"`
	ChanelName string `json:"chanelName"`
}

type RpcGetProcessedResponse struct {
	Id         string `json:"id"`
	ChanelName string `json:"chanelName"`
	Result     string `json:"result"`
	StatusCode uint16 `json:",chanelName"`
}
