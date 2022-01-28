/**
 * @Author: Ne-21
 * @Description:
 * @File:  utils
 * @Version: 1.0.0
 * @Date: 2022/1/28 11:02
 */

package utils

import (
	"io/ioutil"
	"net/http"
)

func ReadImgData(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	pix, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return pix
}
