package routes

import (
	controller "github.com/ddebixx/go-jwt-auth/controllers"
	"github.com/ddebixx/go-jwt-auth/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Aunthenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
