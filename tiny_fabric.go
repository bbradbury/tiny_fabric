package main

import (
	"fmt"
	"github.com/fsouza/go-dockerclient"
	"github.com/gin-gonic/gin"
)

//ServiceManager holds the magic
type ServiceManager struct {
	DockerClient *docker.Client
}

type Service struct {
	Name string
}

func NewServiceManager() (manager *ServiceManager, err error) {
	manager = new(ServiceManager)
	manager.DockerClient, err = docker.NewClient("tcp://192.168.59.103:2375")
	if err != nil {
		err = fmt.Errorf("service_manager couldn't start docker.Client: %s", err)
	}
	return
}

func (s *ServiceManager) NewServiceInstance(c *gin.Context) {
	c.String(200, "Started a service!")
}

func hello(c *gin.Context) {
	c.String(200, "Welcome too the home page, %s!", c.Params.ByName("name"))
}

func main() {
	manager, _ := NewServiceManager()
	fmt.Print(manager)
	//todo: probably use https://github.com/emicklei/go-restful
	router := gin.Default()
	router.GET("/hello/:name", hello)
	router.POST("/services/:service_name/instances", manager.NewServiceInstance)
	/*router.GET("/services", func(_ http.ResponseWriter, _ *http.Request) {})
	router.POST("/services/:service_name", createService)
	router.POST("/services/:service_name/instances", func(_ http.ResponseWriter, _ *http.Request) {})
	router.GET("/services/:service_name/instances/:instance_id", func(_ http.ResponseWriter, _ *http.Request) {})
	router.POST("/services/:service_name/instances/:instance_id/logs", func(_ http.ResponseWriter, _ *http.Request) {})
	router.POST("/services/:service_name/instances/:instance_id/logs/:log_id", func(_ http.ResponseWriter, _ *http.Request) {}) */
	router.Run(":8080")
}
