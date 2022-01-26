/**
 * @Author: Ne-21
 * @Description:
 * @File:  route
 * @Version: 1.0.0
 * @Date: 2022/1/26 13:38
 */

package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRoute(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 666,
		})
	})
}
