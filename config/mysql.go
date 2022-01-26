/**
 * @Author: Ne-21
 * @Description:
 * @File:  database
 * @Version: 1.0.0
 * @Date: 2022/1/26 13:49
 */

package config

import (
	"fmt"
	"github.com/Olixn/Bing-Wallpaper-Server/model"
	"log"
)
import "gorm.io/gorm"
import "gorm.io/driver/mysql"

var MySQL *gorm.DB

func InitMySQL() {
	var err error
	var AppConfig = AppConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		AppConfig.MySql.UserName,
		AppConfig.MySql.PassWord,
		AppConfig.MySql.Address,
		AppConfig.MySql.Port,
		AppConfig.MySql.Dbname)
	MySQL, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 256, // string 类型字段的默认长度
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
		return
	}

	// 自动迁移
	err = MySQL.AutoMigrate(&model.BingWallpaper{})
	if err != nil {
		log.Fatalf("数据表初始化失败: %v", err)
		return
	}
}
