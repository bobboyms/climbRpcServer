package main

import (
	"github.com/gin-gonic/gin"
	"message/managerfilesystem"
	"message/modeldatatransfer"
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
