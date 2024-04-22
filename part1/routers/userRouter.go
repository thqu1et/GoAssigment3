package routers

import (
	"assigment3/controller"
	"assigment3/middleWare"
	"github.com/gin-gonic/gin"
)

func UserRouter(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleWare.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
