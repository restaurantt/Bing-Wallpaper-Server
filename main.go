/**
 * @Author: Ne-21
 * @Description:
 * @File:  main.go
 * @Version: 1.0.0
 * @Date: 2022/1/26 13:18
 */

package main

import (
	"github.com/Olixn/Bing-Wallpaper-Server/config"
	"github.com/Olixn/Bing-Wallpaper-Server/controller"
	"github.com/Olixn/Bing-Wallpaper-Server/route"
	"github.com/gin-gonic/gin"
)

func init() {
	config.InitConfig()
	config.InitMySQL()
	controller.InitBingCrawler()
}

func main() {
	r := gin.Default()
	route.RegisterRoute(r)
	r.Run()
}
