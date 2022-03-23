package controllers

import (
	"course-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
)

type ProgressDetails struct {
	DB *gorm.DB
}

type progressDetailResponse struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	TaskID uint   `json:"taskId"`
}

type progressDetailCreateResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type allProgressDetailResponse struct {
	ID         uint `json:"id"`
	TaskID     uint `json:"taskId"`
	QuestionID uint `json:"questionId"`
	UserID     uint `json:"userId"`
}

type updateProgressDetailForm struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type responseallQuestion struct {
	CourseID uint `json:"courseId"`
	UserID   uint `json:"userId"`
}
type responseallQuestionStatic struct {
	CourseID uint `json:"courseId"`
	UserID   uint `json:"userId"`
	User     []struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	} `json:"user"`
}
type allQuestions struct {
	CourseID uint `json:"id"`
	UserID   uint `json:"userId"`
}
type allQuestions2 struct {
	CourseID uint `json:"courseId"`
	UserID   uint `json:"userId"`
}

type userId struct {
	UserID uint `json:"userId"`
}

func (c *ProgressDetails) FindAll(ctx *gin.Context) {
	var progressDetails []models.ProgressDetail
	c.DB.Order("id desc").Find(&progressDetails)

	var serializedProgressDetail []allProgressDetailResponse
	copier.Copy(&serializedProgressDetail, &progressDetails)
	ctx.JSON(http.StatusOK, gin.H{"courses": serializedProgressDetail})
}

func (c *ProgressDetails) FindOne(ctx *gin.Context) {
	progressDetail, err := c.findProgressDetailByID(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var serializedProgressDetail progressDetailResponse
	copier.Copy(&serializedProgressDetail, &progressDetail)
	ctx.JSON(http.StatusOK, gin.H{"progressdetail": serializedProgressDetail})
}
func (c *ProgressDetails) FindOneuser(ctx *gin.Context) {

	var progressDetail []models.ProgressDetail
	user_id := ctx.Param("id")
	// "id = ?", "1b74413f-f3b8-409f-ac47-e8c062e3472a"
	c.DB.Order("question_id").Find(&progressDetail, "user_id = ?", user_id)

	var serializedProgressesDetail []allProgressDetailResponse
	copier.Copy(&serializedProgressesDetail, &progressDetail)
	ctx.JSON(http.StatusOK, gin.H{"progressdetail": serializedProgressesDetail})
}

func (c *ProgressDetails) CountQuestion(ctx *gin.Context) {
	var progressDetails []models.ProgressDetail

	var requestBodys allQuestions
	if err := ctx.BindJSON(&requestBodys); err != nil {

	}
	if err := c.DB.Find(&progressDetails, "course_id = ? AND user_id = ?", requestBodys.CourseID, requestBodys.UserID).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{"counter": 0})

		return
	}

	var serializedProgressesDetail []responseallQuestion
	copier.Copy(&serializedProgressesDetail, &progressDetails)
	if serializedProgressesDetail == nil {
		ctx.JSON(http.StatusOK, gin.H{"counter": 0})
		return

	} else {
		ctx.JSON(http.StatusOK, gin.H{"counter": serializedProgressesDetail})
		return
	}
}

func (c *ProgressDetails) CountQuestionStatic(ctx *gin.Context) {
	var progressDetails []models.ProgressDetail

	var requestBodys allQuestions2
	if err := ctx.BindJSON(&requestBodys); err != nil {

	}
	if err := c.DB.Preload("User").Find(&progressDetails, "course_id = ? AND user_id = ?", requestBodys.CourseID, requestBodys.UserID).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{"counter": 0})

		return
	}

	var serializedProgressesDetail []responseallQuestionStatic
	copier.Copy(&serializedProgressesDetail, &progressDetails)
	if serializedProgressesDetail == nil {
		ctx.JSON(http.StatusOK, gin.H{"counter": 0})
		return

	} else {
		ctx.JSON(http.StatusOK, gin.H{"counter": serializedProgressesDetail})
		return
	}
}

func (c *ProgressDetails) StaticCourse(ctx *gin.Context) {
	var progressDetail []models.ProgressDetail
	course_id := ctx.Param("id")
	// "id = ?", "1b74413f-f3b8-409f-ac47-e8c062e3472a"
	c.DB.Order("user_id").Find(&progressDetail, "course_id = ?", course_id)

	var serializedProgressesDetail []allProgressDetailResponse
	copier.Copy(&serializedProgressesDetail, &progressDetail)
	ctx.JSON(http.StatusOK, gin.H{"progressdetail": serializedProgressesDetail})

}
func (c *ProgressDetails) Create(ctx *gin.Context) {
	var form progressDetailForm
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var progressDetail models.ProgressDetail
	copier.Copy(&progressDetail, &form)
	if err := c.DB.Create(&progressDetail).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var serializedProgressDetail progressDetailCreateResponse
	copier.Copy(&serializedProgressDetail, &progressDetail)
	ctx.JSON(http.StatusCreated, gin.H{"progressdetail": serializedProgressDetail})
}

func (c *ProgressDetails) Update(ctx *gin.Context) {
	var form updateProgressDetailForm
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	progressDetail, err := c.findProgressDetailByID(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := c.DB.Model(&progressDetail).Update(&form).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var serializedProgressDetail progressDetailResponse
	copier.Copy(&serializedProgressDetail, &progressDetail)
	ctx.JSON(http.StatusOK, gin.H{"ProgressDetails": serializedProgressDetail})
}

func (c *ProgressDetails) Delete(ctx *gin.Context) {
	progressDetail, err := c.findProgressDetailByID(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.DB.Unscoped().Delete(&progressDetail)
	ctx.Status(http.StatusNoContent)
}

func (c *ProgressDetails) findProgressDetailByID(ctx *gin.Context) (*models.ProgressDetail, error) {
	var progressDetail models.ProgressDetail
	id := ctx.Param("id")

	if err := c.DB.Preload("Progress").First(&progressDetail, id).Error; err != nil {
		return nil, err
	}

	return &progressDetail, nil
}
