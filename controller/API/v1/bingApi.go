/**
 * @Author: Ne-21
 * @Description:
 * @File:  bing
 * @Version: 1.0.0
 * @Date: 2022/1/27 11:23
 */

package v1

import (
	"fmt"
	"github.com/Olixn/Bing-Wallpaper-Server/config"
	"github.com/Olixn/Bing-Wallpaper-Server/model"
	"github.com/Olixn/Bing-Wallpaper-Server/utils"
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

func Download(c *gin.Context) {
	picName := c.DefaultQuery("copyright", "")
	picUrlBase := c.DefaultQuery("urlBase", "")
	h := c.DefaultQuery("h", "1080")
	w := c.DefaultQuery("w", "1920")
	fmt.Println(picName, picUrlBase)
	if picUrlBase == "" {
		c.JSON(http.StatusNotFound, nil)
		return
	}
	url := config.AppConfig.API.Baseurl + picUrlBase + "_" + w + "x" + h + ".jpg"
	imgData := utils.ReadImgData(url)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+picName+".jpg")
	c.Header("Content-Transfer-Encoding", "binary")
	_, err := c.Writer.Write(imgData)
	if err != nil {
		return
	}
	return
}

func Today(c *gin.Context) {
	w := c.DefaultQuery("w", "1920")
	h := c.DefaultQuery("h", "1080")
	t := c.DefaultQuery("t", "jpg")

	var bingWallpaper *model.BingWallpaper
	var MySQL = config.MySQL
	MySQL.Model(&model.BingWallpaper{}).Last(&bingWallpaper)

	if t == "jpg" {
		url := config.AppConfig.API.Baseurl + bingWallpaper.UrlBase + "_" + w + "x" + h + ".jpg"
		imgData := utils.ReadImgData(url)
		c.Header("Content-Type", "image/jpg")
		_, err := c.Writer.Write(imgData)
		if err != nil {
			return
		}
	} else if t == "json" {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":    1,
			"msg":       "success",
			"url":       config.AppConfig.API.Baseurl + bingWallpaper.UrlBase + "_" + w + "x" + h + ".jpg",
			"copyright": bingWallpaper.Copyright,
			"date":      bingWallpaper.EndDate,
		})
	} else {
		c.JSON(http.StatusBadGateway, nil)
	}

	return
}
