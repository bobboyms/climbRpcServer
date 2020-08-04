package rpcservice

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

const IndexReceiver = "RECEIVER.m"
const IndexProcced = "PROCCED.m"

//const INDEX_RESPONSE = "RESPONSE.m"

type RpcServiceImp struct {
	fileSystem managerfilesystem.FileSystem
}

func (r *RpcServiceImp) Init(filesystem managerfilesystem.FileSystem) *RpcServiceImp {
	r.fileSystem = filesystem
	return r
}

func (r *RpcServiceImp) createDir(chanelName string) string {

	return r.fileSystem.CreateDirectory(strings.ToLower(chanelName))

}

func (r *RpcServiceImp) createRpcProceedFile(dir string, request modeldatatransfer.RpcProcessedRequest) (string, error) {

	idFileName := request.Id

	return idFileName, r.fileSystem.CreateFile(idFileName+managerfilesystem.FileExtension, dir, model.RpcProcessed{
		Id:         idFileName,
		ChanelName: request.ChanelName,
		Result:     request.Result,
	})
}

func (r *RpcServiceImp) createRpcFile(dir string, request modeldatatransfer.RpcProducerRequest) (string, error) {

	idFileName := managerfilesystem.GetUuid()

	args := make([]model.RpcArgs, len(request.Args))

	for i, arg := range request.Args {
		args[i] = model.RpcArgs{
			Name:  arg.Name,
			Type:  arg.Type,
			Value: arg.Value,
		}
	}

	return idFileName, r.fileSystem.CreateFile(idFileName+managerfilesystem.FileExtension, dir, model.RpcFileStruct{
		Id:         idFileName,
		ClassName:  request.ClassName,
		MethodName: request.MethodName,
		Args:       args,
	})
}

func (r *RpcServiceImp) RpcCreate(request modeldatatransfer.RpcProducerRequest) modeldatatransfer.RpcProducerResponse {

	dir := r.createDir(request.ChanelName)

	var index managerindex.ManagerIndex = new(managerindex.ManagerIndexImp).
		Init(dir, r.fileSystem)

	index.CreateIndex(IndexReceiver)

	fileName, err := r.createRpcFile(dir, request)

	index.AddItem(fileName, IndexReceiver)

	if err != nil {
		return r.getError(err.Error())
	}

	return modeldatatransfer.RpcProducerResponse{
		Id:         fileName,
		StatusCode: StatusOk,
		Message:    "RPC created successfully",
	}

}

func (r *RpcServiceImp) RpcGet(request modeldatatransfer.RpcGetRequest) modeldatatransfer.RpcGetResponse {

	dir := request.ChanelName

	if !r.fileSystem.IsFileExist(dir, IndexReceiver) {
		return modeldatatransfer.RpcGetResponse{
			Id:         "",
			StatusCode: StatusNotFound,
			ClassName:  "",
			MethodName: "",
			Args:       make([]modeldatatransfer.RpcArgs, 0),
		}
	}

	var index managerindex.ManagerIndex = new(managerindex.ManagerIndexImp).
		Init(dir, r.fileSystem)

	item, err := index.GetFirstIndex(IndexReceiver)

	if err != nil {
		return modeldatatransfer.RpcGetResponse{
			Id:         "",
			StatusCode: StatusNotFound,
			ClassName:  "",
			MethodName: "",
			Args:       make([]modeldatatransfer.RpcArgs, 0),
		}
	}

	fileName, _ := index.GetItem(item, IndexReceiver)

	var rpc model.RpcFileStruct
	r.fileSystem.OpenFile(dir, fileName+managerfilesystem.FileExtension, &rpc)

	defer func() {
		index.RemoveItem(item, IndexReceiver)
		r.fileSystem.DeleteFile(dir, fileName+managerfilesystem.FileExtension)
	}()

	println(len(rpc.Args))

	args := make([]modeldatatransfer.RpcArgs, len(rpc.Args))

	for i, arg := range rpc.Args {
		args[i] = modeldatatransfer.RpcArgs{
			Name:  arg.Name,
			Type:  arg.Type,
			Value: arg.Value,
		}
	}

	return modeldatatransfer.RpcGetResponse{
		Id:         rpc.Id,
		StatusCode: StatusOk,
		ClassName:  rpc.ClassName,
		MethodName: rpc.MethodName,
		Args:       args,
	}

}

func (r *RpcServiceImp) getError(message string) modeldatatransfer.RpcProducerResponse {
	return modeldatatransfer.RpcProducerResponse{
		Id:         "0",
		StatusCode: StatusInternalError,
		Message:    message,
	}
}

func (r *RpcServiceImp) RpcCreateProceed(request modeldatatransfer.RpcProcessedRequest) modeldatatransfer.RpcProcessedResponse {

	dir := r.createDir(request.ChanelName)

	var index managerindex.ManagerIndex = new(managerindex.ManagerIndexImp).
		Init(dir, r.fileSystem)

	index.CreateIndex(IndexProcced)

	fileName, _ := r.createRpcProceedFile(dir, request)

	index.AddItem(fileName, IndexProcced)

	//if err != nil {
	//	return r.getError(err.Error())
	//}

	return modeldatatransfer.RpcProcessedResponse{
		Id:         fileName,
		StatusCode: StatusOk,
	}

}

func (r *RpcServiceImp) RpcGetProceed(request modeldatatransfer.RpcGetProcessedRequest) modeldatatransfer.RpcGetProcessedResponse {

	dir := request.ChanelName

	if !r.fileSystem.IsFileExist(dir, IndexProcced) {
		return modeldatatransfer.RpcGetProcessedResponse{
			Id:         "",
			ChanelName: "",
			Result:     "",
			StatusCode: StatusNotFound,
		}
	}

	var index managerindex.ManagerIndex = new(managerindex.ManagerIndexImp).
		Init(dir, r.fileSystem)

	item, err := index.GetIndex(request.Id, IndexProcced)

	if err != nil {
		return modeldatatransfer.RpcGetProcessedResponse{
			Id:         "",
			ChanelName: "",
			Result:     "",
			StatusCode: StatusInternalError,
		}
	}

	fileName, _ := index.GetItem(item, IndexProcced)
	var rpc model.RpcProcessed

	println("****** file name *******")
	println(fileName)

	if !r.fileSystem.IsFileExist(dir, fileName+managerfilesystem.FileExtension) {
		return modeldatatransfer.RpcGetProcessedResponse{
			Id:         "",
			ChanelName: "",
			Result:     "",
			StatusCode: StatusNotFound,
		}
	}

	r.fileSystem.OpenFile(dir, fileName+managerfilesystem.FileExtension, &rpc)

	defer func() {
		index.RemoveItem(item, IndexProcced)
		r.fileSystem.DeleteFile(dir, fileName+managerfilesystem.FileExtension)
	}()

	return modeldatatransfer.RpcGetProcessedResponse{
		Id:         rpc.Id,
		ChanelName: rpc.ChanelName,
		Result:     rpc.Result,
		StatusCode: StatusOk,
	}

}
