package routes

import (
	controllers "golang-jwt-project/Controllers"
	middleware "golang-jwt-project/Middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incoming_routes *gin.Engine) {

	incoming_routes.Use(middleware.Authenticate())

	incoming_routes.GET("users/", controllers.GetUsers())
	incoming_routes.GET("Useres/:user_id", controllers.GetUser())
}
