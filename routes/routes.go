package routes

import (
	"gi-api-rest/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	GinServer := gin.Default()
	GinServer.GET("/students", controllers.ShowAllStudents)
	GinServer.POST("/students", controllers.NewStudent)
	GinServer.PATCH("/students/:id", controllers.EditStudent)
	GinServer.GET("/students/:id", controllers.FindById)
	GinServer.DELETE("/students/:id", controllers.DeleteStudent)

	GinServer.Run(":3333")
}
