/**
 * @Author: Ne-21
 * @Description: 自动获取每日壁纸爬虫
 * @File:  bingCrawler
 * @Version: 1.0.0
 * @Date: 2022/1/26 13:47
 */

package controller

import (
	"encoding/json"
	"log"

	"github.com/imroc/req"
	"github.com/restaurantt/Bing-Wallpaper-Server/config"
	"github.com/restaurantt/Bing-Wallpaper-Server/model"
	"github.com/robfig/cron"
	"gorm.io/gorm"
)

type BingCrawler struct {
	MySql *gorm.DB
}

func InitBingCrawler() {
	bingCrawler := &BingCrawler{
		MySql: config.MySQL,
	}

	// 凌晨1点爬取
	c := cron.New()
	spec := "0 0 1 * * ? "
	// spec := "0 */1 * * * ?"
	c.AddFunc(spec, func() {
		bingCrawler.Start()
	})
	c.Start()
	log.Printf("定时任务已启动！")
}
func (b *BingCrawler) Start() {
	get, err := req.Get(config.AppConfig.API.Baseurl + "/HPImageArchive.aspx?format=js&idx=0&n=1&mkt=zh-CN")
	if err != nil {
		log.Printf("请求官方API出错: %v", err)
		return
	}

	var res = make(map[string]interface{})
	err = json.Unmarshal(get.Bytes(), &res)
	if err != nil {
		log.Printf("反序列化出错: %v", err)
		return
	}

	obj, err := json.Marshal(res["images"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return
	}
	var bingImages *model.BingImages
	err = json.Unmarshal(obj, &bingImages)
	if err != nil {
		return
	}

	if b.getByDate(bingImages.EndDate) {
		b.MySql.Model(&model.BingWallpaper{}).Create(&model.BingWallpaper{
			BingImages: *bingImages,
		})
		log.Printf("%v存入数据库", bingImages.EndDate)
	} else {
		log.Printf("数据已存在或出错了")
	}

}

func (b *BingCrawler) getByDate(d string) bool {
	var bingWallpaper *model.BingWallpaper
	res := b.MySql.Model(&model.BingWallpaper{}).Where("end_date = ?", d).Find(&bingWallpaper)
	if res.RowsAffected > 0 {
		return false
	}
	return true
}
