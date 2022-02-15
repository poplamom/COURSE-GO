package controllers

import (
	"course-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
)

type Progresses struct {
	DB *gorm.DB
}

type progressResponse struct {
	ID       uint `json:"id"`
	CourseID uint `json:"courseId"`
	Course   struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
		Desc string `json:"desc"`
	} `json:"course"`
	Progressdetaile []struct {
		ID     uint `json:"id"`
		TaskID uint `json:"taskId"`
	} `json:"progressdetail"`
}

type progressCreateResponse struct {
	ID       uint `json:"id"`
	CourseID uint `json:"courseId"`
}

type allProgressesResponse struct {
	ID       uint `json:"id"`
	CourseID uint `json:"courseId"`
}

type createProgressesForm struct {
	CourseID uint `form:"courseId" binding:"required"`
}

type updateProgressesForm struct {
	ID       uint `json:"id"`
	CourseID uint `json:"CourseID"`
}

func (c *Progresses) FindAll(ctx *gin.Context) {
	var progresses []models.Progress
	c.DB.Order("id desc").Find(&progresses)

	var serializedProgresses []allProgressesResponse
	copier.Copy(&serializedProgresses, &progresses)
	ctx.JSON(http.StatusOK, gin.H{"progresses": serializedProgresses})
}

func (c *Progresses) FindOne(ctx *gin.Context) {
	progress, err := c.findProgressByID(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var serializedProgress progressResponse
	copier.Copy(&serializedProgress, &progress)
	ctx.JSON(http.StatusOK, gin.H{"progress": serializedProgress})
}

func (c *Progresses) Create(ctx *gin.Context) {
	var form createProgressesForm
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var progress models.Progress
	copier.Copy(&progress, &form)
	if err := c.DB.Create(&progress).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var serializedProgress progressCreateResponse
	copier.Copy(&serializedProgress, &progress)
	ctx.JSON(http.StatusCreated, gin.H{"progress": serializedProgress})
}

func (c *Progresses) Update(ctx *gin.Context) {
	var form createProgressesForm
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	progress, err := c.findProgressByID(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := c.DB.Model(&progress).Update(&form).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var serializedProgress progressResponse
	copier.Copy(&serializedProgress, &progress)
	ctx.JSON(http.StatusOK, gin.H{"progress": serializedProgress})
}

func (c *Progresses) Delete(ctx *gin.Context) {
	progress, err := c.findProgressByID(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.DB.Unscoped().Delete(&progress)
	ctx.Status(http.StatusNoContent)
}

func (c *Progresses) findProgressByID(ctx *gin.Context) (*models.Progress, error) {
	var progress models.Progress
	id := ctx.Param("id")

	if err := c.DB.Preload("progress").First(&progress, id).Error; err != nil {
		return nil, err
	}

	return &progress, nil
}
