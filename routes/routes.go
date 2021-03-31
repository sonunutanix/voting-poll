package routes

import (
	"Project/controllers"

	"github.com/gin-gonic/gin"
)

func Setup(app *gin.Engine) {
	app.POST("/api/register", controllers.RegisterUser)
	app.POST("/api/login", controllers.Login)
	app.GET("/api/user", controllers.User)
	app.POST("/api/logout", controllers.Logout)
	app.POST("/api/create-poll", controllers.CreatePoll)
	app.GET("/api/getallpolls", controllers.GetAllPolls)
	app.GET("/api/getpoll/:id", controllers.GetPollById)
}
