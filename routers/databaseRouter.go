package routers

import (
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/controllers"
	"github.com/gin-gonic/gin"
)

func SetupDatabaseRoutes(router *gin.Engine) {
	databaseGroup := router.Group("/database")
	{
		databaseGroup.GET("/getdatabaseaws/:userid",controllers.GetDatabaseAwsHandler)
		databaseGroup.POST("/createdatabaseaws/:userid",controllers.CreateDatabaseAwsHandler)
		databaseGroup.POST("/deletedatabaseaws/:userid",controllers.DeleteDatabaseAwsHandler)
		databaseGroup.GET("/getdatabaseazure",controllers.GetDatabaseAzureHandler)
		databaseGroup.POST("/createdatabaseazure",controllers.CreateDatabaseAzureHandler)
		databaseGroup.POST("/deletedatabaseazure",controllers.DeleteDatabaseAzureHandler)
	}
}
