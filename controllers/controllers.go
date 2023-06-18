package controllers

import (
	"gi-api-rest/database"
	"gi-api-rest/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowAllStudents(context *gin.Context) {
	var students []models.Student

	database.DB.Find(&students)
	context.JSON(200, students)
}

func NewStudent(context *gin.Context) {
	var student models.Student

	if error := context.ShouldBindJSON(&student); error != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": error.Error(),
		})
		return
	}
	database.DB.Create(&student)
	context.JSON(http.StatusOK, student)
}

func FindById(context *gin.Context) {
	var student models.Student
	id := context.Params.ByName("id")
	database.DB.First(&student, id)

	if student.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Student not found",
		})
		return
	}

	context.JSON(http.StatusOK, student)
}

func DeleteStudent(context *gin.Context) {
	var student models.Student
	id := context.Params.ByName("id")
	database.DB.Delete(&student, id)

	context.JSON(http.StatusOK, gin.H{
		"message": "Student was deleted!",
	})
}

func EditStudent(context *gin.Context) {
	var student models.Student
	id := context.Params.ByName("id")
	database.DB.First(&student, id)

	if error := context.ShouldBindJSON(&student); error != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": error.Error(),
		})
		return
	}

	database.DB.Model(&student).UpdateColumns(student)
	context.JSON(http.StatusOK, student)
}
