package util

import (
	setting "github.com/MrChenCode/blog/pkg"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func Getpage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}
	return result
}
