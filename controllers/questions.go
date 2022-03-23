package controllers

import (
	"course-go/models"
	"net/http"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type Questions struct {
	DB *gorm.DB
}
type ProgressDetailses struct {
	DB *gorm.DB
}
type ProgressDetail struct {
	CourseID uint

	TaskID     uint
	QuestionID uint
	UserID     uint
}

type QuestionAnser struct {
	ID       uint   `json:"id"`
	Answer   string `json:"answer"`
	CourseID uint   `json:"courseId"`
	TaskID   uint   `json:"taskId"`
	UserID   uint   `json:"userId"`
}
type ProgressDetaialCreate struct {
	TaskID     uint `json:"takId"`
	QuestionID uint `json:"questionId"`
	UserID     uint `json:"userId"`
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
type allQuestionName struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Hint string `json:"hint"`
}
type createQuestionForm struct {
	Name     string `form:"name" binding:"required"`
	Answer   string `form:"answer" binding:"required"`
	Hint     string `form:"hint" binding:"required"`
	CourseID uint   `form:"courseId" binding:"required"`
	TaskID   uint   `form:"taskId" binding:"required"`
}

type updateQuestionForm struct {
	Name   string `form:"name"`
	Desc   string `form:"desc"`
	Answer string `form:"answer"`
	Hint   string `form:"hint"`
	TaskID uint   `form:"taskId"`
}

type progressDetailForm struct {
	UserID     uint `form:"userId" binding:"required"`
	TaskID     uint `form:"taskId" binding:"required"`
	QuestionID uint `form:"questionId" binding:"required"`
}
type User struct {
	ID   int64
	Name string
	Age  byte
}
type allQuestionx struct {
	CourseID uint `json:"courseId"`
}
type allQuestion2 struct {
	CourseID uint `json:"courseId"`
}
type responseallQuestions struct {
	CourseID uint `json:"courseId"`
}

func (c *Questions) FindAll(ctx *gin.Context) {
	var questions []models.Question
	c.DB.Order("id desc").Find(&questions)

	var serializedQuestion []allQuestionResponse
	copier.Copy(&serializedQuestion, &questions)
	ctx.JSON(http.StatusOK, gin.H{"questions": serializedQuestion})
}
func (c *Questions) FindAllName(ctx *gin.Context) {
	var questions []models.Question
	c.DB.Order("id desc").Find(&questions)

	var serializedQuestion []allQuestionName
	copier.Copy(&serializedQuestion, &questions)
	ctx.JSON(http.StatusOK, gin.H{"question": serializedQuestion})
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
	// var progressdetail models.ProgressDetail

	var requestBody QuestionAnser
	if err := ctx.BindJSON(&requestBody); err != nil {
		// DO SOMETHING WITH THE ERROR
	}
	// question_id := ctx.PostForm("id")

	err := c.DB.Find(&question, "id = ? AND answer = ?", requestBody.ID, requestBody.Answer).Error

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})

		// ctx.JSON(http.StatusNotFound, gin.H{"error": "no"})
		return
	}

	progessDetailTables := ProgressDetail{CourseID: requestBody.CourseID, TaskID: requestBody.TaskID, QuestionID: requestBody.ID, UserID: requestBody.UserID}
	if err := c.DB.Create(&progessDetailTables).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"question": "yes"})
}

func (cc *ProgressDetailses) createProgressDetail(ctx *gin.Context) {
	var requestBody QuestionAnser
	if err := ctx.BindJSON(&requestBody); err != nil {
		// DO SOMETHING WITH THE ERROR
	}
	progressdetail := ProgressDetail{CourseID: requestBody.CourseID, TaskID: requestBody.TaskID, QuestionID: requestBody.ID, UserID: requestBody.UserID}
	if err := cc.DB.Create(&progressdetail).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	return
}

func (c *Questions) FindQuestionByCourse(ctx *gin.Context) {
	var question []models.Question
	var requestBodys allQuestion2
	if err := ctx.BindJSON(&requestBodys); err != nil {

	}
	if err := c.DB.Find(&question, "course_id = ?", requestBodys.CourseID).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{"counter": 0})

		return
	}

	var serializedQuestion []responseallQuestions
	copier.Copy(&serializedQuestion, &question)

	ctx.JSON(http.StatusOK, gin.H{"counter": serializedQuestion})
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
