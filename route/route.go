/**
 * @Author: Ne-21
 * @Description:
 * @File:  route
 * @Version: 1.0.0
 * @Date: 2022/1/26 13:38
 */

package route

import (
	v1 "github.com/Olixn/Bing-Wallpaper-Server/controller/API/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRoute(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 666,
		})
	})

	apiGroup := r.Group("/api")
	// apiGroup.Use(middleware.Cors())
	{
		apiGroup.GET("/v1/getList", v1.GetWallpapersList)
		apiGroup.PUT("/v1/view/:id", v1.View)
		apiGroup.GET("/v1/download", v1.Download)
	}
}
