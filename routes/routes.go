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
	app.GET("/api/getallpolls", controllers.GetAllPolls) // get all questions
	app.GET("/api/getpoll/:id", controllers.GetPollById)
	app.PATCH("/api/option/:id", controllers.UpdateOptionVote) //Update Option vote
	app.POST("/api/option-user", controllers.SaveOptionIdUserId) // save userId who vote for the option
	app.GET("api/option-user/:id", controllers.GetUserListOfOption) //Get all user who vote for a option
	app.GET("api/options/:id", controllers.GetOptionsByQuesId) // Get Options By question Id
}
