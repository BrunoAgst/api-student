package routes

import (
	"api-rest/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.ViewAllStudents)
	r.GET("/:name", controllers.Welcome)
	r.GET("/students/:id", controllers.SearchStudent)
	r.POST("/students", controllers.CreateStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.PATCH("/students/:id", controllers.EditStudent)
	r.GET("/students/cpf/:cpf", controllers.SearchStudentWithCPF)
	r.Run(":3000")
}
