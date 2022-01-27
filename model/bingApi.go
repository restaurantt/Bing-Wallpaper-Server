/**
 * @Author: Ne-21
 * @Description:
 * @File:  bingApi
 * @Version: 1.0.0
 * @Date: 2022/1/27 11:18
 */

package model

type BingAPI struct {
	Status int              `json:"status"`
	Msg    string           `json:"msg"`
	Total  int64            `json:"total"`
	Data   []BingWallpapers `json:"data"`
}

type BingWallpapers struct {
	Id int `json:"id"`
	BingImages
	View int `json:"view"`
}
