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

type progressDetailCreateResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type allProgressDetailResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}


type updateProgressDetailForm struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
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
