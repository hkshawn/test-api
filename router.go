package main

import (
	. "api-test/api"
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", IndexUsers) //http://localhost:8806

	//路由群组
	users := router.Group("api/v1/users")
	{
		//users.GET("", GetAll) //http://localhost:8806/api/v1/users
		//users.POST("/add", AddUsers) //http://localhost:8806/api/v1/users/add
		users.GET("/get/:id", Show) //http://localhost:8806/api/v1/users/get/5
		//users.POST("/update", UpdateUser)
		//users.POST("/del", DelUser)
	}

	return router
}
