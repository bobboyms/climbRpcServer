package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message/managerfilesystem"
	"message/modeldatatransfer"
	"message/rpcservice"
	"message/topicservice"
	"net/http"
)

func main() {

	router := gin.Default()

	v1 := router.Group("v1")
	v1.Use(gin.Recovery())
	{
		v1.GET("/user/:name", func(c *gin.Context) {

			name := c.Param("name")
			c.String(http.StatusOK, "Hello %s", name)

		})

		v1.POST("/topic/producer", func(c *gin.Context) {

			var p modeldatatransfer.TopicProducerRequest
			c.BindJSON(&p)

			var fileSystem managerfilesystem.FileSystem = new(managerfilesystem.FileSystemGobEncoder)
			var service topicservice.TopicService = new(topicservice.TopicServiceImp).Init(fileSystem)

			response := service.TopicCreate(p)
			c.JSON(int(response.StatusCode), response)

		})

		v1.GET("/topic/consumer/:name", func(c *gin.Context) {

			name := c.Param("name")

			p := modeldatatransfer.TopicGetRequest{
				TopicName: name,
			}

			var fileSystem managerfilesystem.FileSystem = new(managerfilesystem.FileSystemGobEncoder)
			var service topicservice.TopicService = new(topicservice.TopicServiceImp).Init(fileSystem)

			response := service.TopicGet(p)
			c.JSON(int(response.StatusCode), response)

		})

		v1.POST("/rpc/create", func(c *gin.Context) {

			var p modeldatatransfer.RpcProducerRequest
			c.BindJSON(&p)

			fmt.Printf("%+v\n", p)

			var fileSystem managerfilesystem.FileSystem = new(managerfilesystem.FileSystemGobEncoder)
			var service rpcservice.RpcService = new(rpcservice.RpcServiceImp).Init(fileSystem)

			response := service.RpcCreate(p)
			c.JSON(int(response.StatusCode), response)

		})

		v1.GET("/rpc/consumer/:channel", func(c *gin.Context) {

			channel := c.Param("channel")

			p := modeldatatransfer.RpcGetRequest{
				ChanelName: channel,
			}

			var fileSystem managerfilesystem.FileSystem = new(managerfilesystem.FileSystemGobEncoder)
			var service rpcservice.RpcService = new(rpcservice.RpcServiceImp).Init(fileSystem)

			response := service.RpcGet(p)
			c.JSON(int(response.StatusCode), response)

		})

		v1.POST("/rpc/create/processed", func(c *gin.Context) {

			var p modeldatatransfer.RpcProcessedRequest
			c.BindJSON(&p)

			fmt.Printf("%+v\n", p)

			var fileSystem managerfilesystem.FileSystem = new(managerfilesystem.FileSystemGobEncoder)
			var service rpcservice.RpcService = new(rpcservice.RpcServiceImp).Init(fileSystem)

			response := service.RpcCreateProceed(p)
			c.JSON(int(response.StatusCode), response)

		})

		v1.GET("/rpc/get/processed/:channel/:id", func(c *gin.Context) {

			id := c.Param("id")
			channel := c.Param("channel")

			p := modeldatatransfer.RpcGetProcessedRequest{
				Id:         id,
				ChanelName: channel,
			}

			var fileSystem managerfilesystem.FileSystem = new(managerfilesystem.FileSystemGobEncoder)
			var service rpcservice.RpcService = new(rpcservice.RpcServiceImp).Init(fileSystem)

			response := service.RpcGetProceed(p)
			c.JSON(int(response.StatusCode), response)

		})
	}

	router.Run("0.0.0.0:8090")

}

func xxmain() {

	//var index managerindex.ManagerIndex = new(managerindex.ManagerIndexImp)

	var fileSystem managerfilesystem.FileSystem = new(managerfilesystem.FileSystemBinaryEncoder)

	var manager topicservice.TopicService = new(topicservice.TopicServiceImp).Init(fileSystem)

	manager.TopicCreate(modeldatatransfer.TopicProducerRequest{
		"baixarEstoque", "",
	})

	println(manager)

}
