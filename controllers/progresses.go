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
	UserID   uint `json:"userId"`
	Course   struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
		Desc string `json:"desc"`
	} `json:"course"`
}

type progressCreateResponse struct {
	ID       uint `json:"id"`
	UserID   uint `json:"userId"`
	CourseID uint `json:"courseId"`
}

type allProgressesResponse struct {
	ID       uint `json:"id"`
	UserID   uint `json:"userId"`
	CourseID uint `json:"courseId"`
	Course   []struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
		Desc string `json:"desc"`
	} `json:"course"`
}

type enrollCourse struct {
	ID       uint `json:"id"`
	UserID   uint `json:"userId"`
	CourseID uint `json:"courseId"`
	User     []struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	} `json:"user"`
}

type myCourse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type createProgressesForm struct {
	UserID   uint `form:"userId" binding:"required"`
	CourseID uint `form:"courseId" binding:"required"`
}
type oneProgressesForm struct {
	UserID   uint `form:"userId" binding:"required"`
	CourseID uint `form:"courseId" binding:"required"`
}

type updateProgressesForm struct {
	ID       uint `json:"id"`
	UserID   uint `json:"userId"`
	CourseID uint `json:"courseId"`
}

func (c *Progresses) FindAll(ctx *gin.Context) {
	var progresses []models.Progress
	c.DB.Order("course_id desc").Find(&progresses)

	var serializedProgresses []allProgressesResponse
	copier.Copy(&serializedProgresses, &progresses)
	ctx.JSON(http.StatusOK, gin.H{"progress": serializedProgresses})
}

func (c *Progresses) FindOne(ctx *gin.Context) {
	progress, err := c.findProgressByID(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var serializedProgress allProgressesResponse
	copier.Copy(&serializedProgress, &progress)
	ctx.JSON(http.StatusOK, gin.H{"progress": serializedProgress})
}

func (c *Progresses) FindOneuser(ctx *gin.Context) {

	var progresses []models.Progress
	user_id := ctx.Param("id")
	// "id = ?", "1b74413f-f3b8-409f-ac47-e8c062e3472a"
	c.DB.Order("course_id desc").Find(&progresses, "user_id = ?", user_id)

	var serializedProgresses []allProgressesResponse
	copier.Copy(&serializedProgresses, &progresses)
	ctx.JSON(http.StatusOK, gin.H{"progress": serializedProgresses})
}

func (c *Progresses) FindMyCourse(ctx *gin.Context) {

	var progresses []models.Progress
	user_id := ctx.Param("id")
	// "id = ?", "1b74413f-f3b8-409f-ac47-e8c062e3472a"
	if err := c.DB.Preload("Course").Order("course_id").Find(&progresses, "user_id = ?", user_id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	var serializedProgresses []allProgressesResponse
	copier.Copy(&serializedProgresses, &progresses)
	ctx.JSON(http.StatusOK, gin.H{"progresses": serializedProgresses})
}
func (c *Progresses) FinduserEnroll(ctx *gin.Context) {

	var progresses []models.Progress
	course_id := ctx.Param("id")
	// "id = ?", "1b74413f-f3b8-409f-ac47-e8c062e3472a"
	if err := c.DB.Preload("User").Order("user_id").Find(&progresses, "course_id = ?", course_id).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{"users": 0})
	}

	var serializedProgresses []enrollCourse
	copier.Copy(&serializedProgresses, &progresses)
	ctx.JSON(http.StatusOK, gin.H{"users": serializedProgresses})
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

	if err := c.DB.First(&progress, id).Error; err != nil {
		return nil, err
	}

	return &progress, nil
}
