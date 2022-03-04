package controllers

import (
	"course-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
)

type Questions struct {
	DB *gorm.DB
}

type questionResponse struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Answer string `json:"answer"`
	Hint   string `json:"hint"`
	TaskID uint   `json:"taskId"`
	Task   struct {
		ID        uint   `json:"id"`
		Name      string `json:"name"`
		Desc      string `json:"desc"`
		Objective string `json:"objective"`
	} `json:"task"`
}

type questionCreateResponse struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Answer string `json:"answer"`
	Hint   string `json:"hint"`
	TaskID uint   `json:"taskId"`
}
type questionCheckResponse struct {
	ID     uint   `json:"id"`
	Status string `json:"status"`
}

type allQuestionResponse struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Answer string `json:"answer"`
	Hint   string `json:"hint"`
}

type createQuestionForm struct {
	Name   string `form:"name" binding:"required"`
	Answer string `form:"answer" binding:"required"`
	Hint   string `form:"hint" binding:"required"`
	TaskID uint   `form:"taskId" binding:"required"`
}

type updateQuestionForm struct {
	Name   string `form:"name"`
	Desc   string `form:"desc"`
	Answer string `form:"answer"`
	Hint   string `form:"hint"`
	TaskID uint   `form:"taskId"`
}

func (c *Questions) FindAll(ctx *gin.Context) {
	var questions []models.Question
	c.DB.Order("id desc").Find(&questions)

	var serializedQuestion []allQuestionResponse
	copier.Copy(&serializedQuestion, &questions)
	ctx.JSON(http.StatusOK, gin.H{"questions": serializedQuestion})
}

func (c *Questions) FindOne(ctx *gin.Context) {
	question, err := c.findQuestionByID(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var serializedQuestion questionResponse
	copier.Copy(&serializedQuestion, &question)
	ctx.JSON(http.StatusOK, gin.H{"question": serializedQuestion})
}

func (c *Questions) CheckAns(ctx *gin.Context) {
	var question models.Question
	question_id := ctx.PostForm("id")
	answer := ctx.PostForm("answer")

	err := c.DB.Find(&question, "id = ? AND answer = ?", question_id, answer).Error

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"question": "no"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"question": "yes"})

}

func (c *Questions) Create(ctx *gin.Context) {
	var form createQuestionForm
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"errorx": err.Error()})
		return
	}

	var question models.Question
	copier.Copy(&question, &form)
	if err := c.DB.Create(&question).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"errorz": err.Error()})
		return
	}

	var serializedQuestion questionCreateResponse
	copier.Copy(&serializedQuestion, &question)
	ctx.JSON(http.StatusCreated, gin.H{"question": serializedQuestion})
}

func (c *Questions) Update(ctx *gin.Context) {
	var form updateQuestionForm
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	question, err := c.findQuestionByID(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := c.DB.Model(&question).Update(&form).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var serializedQuestion questionResponse
	copier.Copy(&serializedQuestion, &question)
	ctx.JSON(http.StatusOK, gin.H{"question": serializedQuestion})
}

func (c *Questions) Delete(ctx *gin.Context) {
	question, err := c.findQuestionByID(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.DB.Unscoped().Delete(&question)
	ctx.Status(http.StatusNoContent)
}

func (c *Questions) findQuestionByID(ctx *gin.Context) (*models.Question, error) {
	var question models.Question
	id := ctx.Param("id")

	if err := c.DB.Preload("Task").First(&question, id).Error; err != nil {
		return nil, err
	}

	return &question, nil
}
