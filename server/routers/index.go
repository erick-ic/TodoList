package routers

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func Routers(r *gin.Engine) {
	routers := r.Group("/")
	{
		routers.GET("/hello", controllers.TodosController{}.Hello)
		routers.POST("/createEvent", controllers.TodosController{}.CreateEvent)
		routers.DELETE("/deleteEvent/:id", controllers.TodosController{}.DeleteEvent)
		routers.PUT("/editEvent/:id", controllers.TodosController{}.EditEvent)
		routers.GET("/getEvent", controllers.TodosController{}.GetEvent)
	}
}
