package rpcservice

import "message/modeldatatransfer"

type RpcService interface {
	RpcCreate(request modeldatatransfer.RpcProducerRequest) modeldatatransfer.RpcProducerResponse
	RpcGet(request modeldatatransfer.RpcGetRequest) modeldatatransfer.RpcGetResponse
	RpcCreateProceed(request modeldatatransfer.RpcProcessedRequest) modeldatatransfer.RpcProcessedResponse
	RpcGetProceed(request modeldatatransfer.RpcGetProcessedRequest) modeldatatransfer.RpcGetProcessedResponse
}
