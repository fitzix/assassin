package controllers

import (
	"github.com/gin-gonic/gin"
)

// func (r *queryResolver) AllCategories(ctx context.Context) ([]*models.Category, error) {
// 	var down []*models.Category
// 	if err := r.db.Find(&down).Error; err != nil {
// 		return nil, r.Fail(3000)
// 	}
// 	return down, nil
// }

func CategoryAll(c *gin.Context) {}
func CategoryIndex(c *gin.Context) {}
func CategoryUpdate(c *gin.Context) {}
func CategoryCreate(c *gin.Context) {}
func CategoryDelete(c *gin.Context) {}
