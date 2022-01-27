/**
 * @Author: Ne-21
 * @Description:
 * @File:  bing
 * @Version: 1.0.0
 * @Date: 2022/1/27 11:23
 */

package v1

import (
	"github.com/Olixn/Bing-Wallpaper-Server/config"
	"github.com/Olixn/Bing-Wallpaper-Server/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetWallpapersList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "9"))

	var MySQL = config.MySQL
	var total int64
	MySQL.Model(&model.BingWallpaper{}).Count(&total)

	var bingWallpaperList *[]model.BingWallpapers
	MySQL.Model(&model.BingWallpaper{}).Order("created_at desc").Limit(size).Offset((page - 1) * size).Find(&bingWallpaperList)

	c.JSON(http.StatusOK, &model.BingAPI{
		Status: 1,
		Msg:    "",
		Total:  total,
		Data:   *bingWallpaperList,
	})
	return
}

func View(c *gin.Context) {
	id := c.Param("id")

	var MySQL = config.MySQL
	MySQL.Model(&model.BingWallpaper{}).Where("id = ?", id).Update("view", gorm.Expr("view+ ?", 1))

	c.JSON(http.StatusOK, nil)
}
