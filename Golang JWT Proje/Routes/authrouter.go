package routes

import (
	controllers "golang-jwt-project/Controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incoming_routes *gin.Engine) {

	incoming_routes.POST("users/signup", controllers.Signup())
	incoming_routes.POST("users/Login", controllers.Login())

}
