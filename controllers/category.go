package controllers

import (
	"github.com/fitzix/assassin/models"
	"github.com/fitzix/assassin/service"
	"github.com/gin-gonic/gin"
)

// func (r *queryResolver) AllCategories(ctx context.Context) ([]*models.Category, error) {
// 	var down []*models.Category
// 	if err := r.db.Find(&down).Error; err != nil {
// 		return nil, r.Fail(3000)
// 	}
// 	return down, nil
// }

func CategoryList(c *gin.Context) {
	a := service.NewAsnGin(c)
	var down []models.Category
	if err := a.D.Find(&down).Error; err != nil {
		a.Fail(service.StatusWebBadRequest, err)
		return
	}
	a.Success(down)
}
func CategoryIndex(c *gin.Context)  {}
func CategoryUpdate(c *gin.Context) {}
func CategoryCreate(c *gin.Context) {}
func CategoryDelete(c *gin.Context) {}
