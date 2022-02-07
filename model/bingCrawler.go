/**
 * @Author: Ne-21
 * @Description:
 * @File:  bing
 * @Version: 1.0.0
 * @Date: 2022/1/26 13:43
 */

package model

import "gorm.io/gorm"

type BingImages struct {
	EndDate       string `json:"enddate" gorm:"index"`
	Url           string `json:"url"`
	UrlBase       string `json:"urlbase"`
	Copyright     string `json:"copyright"`
	CopyrightLink string `json:"copyrightlink"`
	Hsh           string `json:"hsh" gorm:"index"`
}

type BingWallpaper struct {
	gorm.Model
	BingImages
	View int `json:"view"`
}
