package controllers

import (
	"github.com/fitzix/assassin/models"
	"github.com/fitzix/assassin/service"
	"github.com/gin-gonic/gin"
)

// func (r *queryResolver) AllTags(ctx context.Context) ([]*models.Tag, error) {
// 	var down []*models.Tag
// 	if err := r.db.Find(&down).Error; err != nil {
// 		return nil, r.Fail(3000)
// 	}
// 	return down, nil
// }

func TagList(c *gin.Context) {
	a := service.NewAsnGin(c)
	var down []models.Tag
	if err := a.D.Find(&down).Error; err != nil {
		a.Fail(service.StatusWebBadRequest, err)
		return
	}
	a.Success(down)
}
func TagIndex(c *gin.Context) {}
func TagCreate(c *gin.Context) {
	a := service.NewAsnGin(c)
	var up models.Tag
	if err :=c.ShouldBindJSON(&up); err != nil {
		a.Fail(service.StatusWebParamErr, err)
		return
	}

	if err :=a.D.Create(&up).Error; err != nil {
		a.Fail(service.StatusWebBadRequest, err)
		return
	}

	a.Success(up)
}
func TagUpdate(c *gin.Context) {}
func TagDelete(c *gin.Context) {}