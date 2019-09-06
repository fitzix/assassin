package controllers

import (
	"github.com/gin-gonic/gin"
)

// func (r *queryResolver) AllTags(ctx context.Context) ([]*models.Tag, error) {
// 	var down []*models.Tag
// 	if err := r.db.Find(&down).Error; err != nil {
// 		return nil, r.Fail(3000)
// 	}
// 	return down, nil
// }

func TagAll(c *gin.Context) {}
func TagIndex(c *gin.Context) {}
func TagCreate(c *gin.Context) {}
func TagUpdate(c *gin.Context) {}
func TagDelete(c *gin.Context) {}