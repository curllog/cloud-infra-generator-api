package routers

import (
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/controllers"
	"github.com/gin-gonic/gin"
)

func SetupNodeGroupRoutes(router *gin.Engine) {
	nodeGroupGroup := router.Group("/billing")
	{
		nodeGroupGroup.POST("/createnodegroupaws",controllers.NodeGroupEksHandlers)
	}
}
