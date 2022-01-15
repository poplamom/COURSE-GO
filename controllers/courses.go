package controllers

import (
	"course-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
)

type Courses struct {
	DB *gorm.DB
}

type courseResponse struct {
	ID         uint   `json:"id"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	Task []struct {
		ID       	uint   `json:"id"`
		Name     	string `json:"name"`
		Desc     	string `json:"desc"`
		Objective 	string `json:"objective"`
		Status 		string `json:"status"`
		// Question []struct{
		// 	ID			uint	`json:"id"`
		// 	Name		string	`json:"name"`
		// 	Answer		string	`json:"answer"`
		// 	Hint		string	`json:"hint"`
		// 	Status 		string
		// }
	} `json:"tasks"`
}

type courseCreateResponse struct {
	ID         uint   `json:"id"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
}


type allCourseResponse struct {
	ID         uint   `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type createCourseForm struct {
	Name string `json:"name" binding:"required"`
	Desc string `json:"desc" binding:"required"`
}

type updateCourseForm struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

func (c *Courses) FindAll(ctx *gin.Context) {
	var courses []models.Course
	c.DB.Order("id desc").Find(&courses)

	var serializedCourse []allCourseResponse
	copier.Copy(&serializedCourse, &courses)
	ctx.JSON(http.StatusOK, gin.H{"courses": serializedCourse})
}

func (c *Courses) FindOne(ctx *gin.Context) {
	course, err := c.findCourseByID(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var serializedCourse courseResponse
	copier.Copy(&serializedCourse, &course)
	ctx.JSON(http.StatusOK, gin.H{"course": serializedCourse})
}

func (c *Courses) Create(ctx *gin.Context) {
	var form createCourseForm
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var course models.Course
	copier.Copy(&course, &form)
	if err := c.DB.Create(&course).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var serializedCourse courseCreateResponse
	copier.Copy(&serializedCourse, &course)
	ctx.JSON(http.StatusCreated, gin.H{"course": serializedCourse})
}

func (c *Courses) Update(ctx *gin.Context) {
	var form updateCourseForm
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	course, err := c.findCourseByID(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := c.DB.Model(&course).Update(&form).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var serializedCourse courseResponse
	copier.Copy(&serializedCourse, &course)
	ctx.JSON(http.StatusOK, gin.H{"course": serializedCourse})
}

func (c *Courses) Delete(ctx *gin.Context) {
	course, err := c.findCourseByID(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.DB.Unscoped().Delete(&course)
	ctx.Status(http.StatusNoContent)
}

func (c *Courses) findCourseByID(ctx *gin.Context) (*models.Course, error)  {
	var course models.Course
	id := ctx.Param("id")

	if err := c.DB.Preload("Task").First(&course, id).Error; err != nil {
		return nil, err
	}

	return &course, nil
}

