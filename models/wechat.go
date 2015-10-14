package models

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/cache"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	TokenCacheTime int64  = 7000                                      //token的存活存储时间，官方为7200秒所以我们
	AccessTokenUrl string = "https://api.weixin.qq.com/cgi-bin/token" //access_token 获取url
)

type Wechat struct {
	Appid           string
	Secret          string
	Grant_type      string
	Access_token    string
	TokenCreateTime time.Time
}

//根据appid, secret获取微信号的access_token值
func (wx *Wechat) CreateToken() string {
	//增加读取缓存代码
	bm, err := cache.NewCache("file", `{"CachePath":"./cache","FileSuffix":".cache","DirectoryLevel":2,"EmbedExpiry":60}`)
	cacheName := "GTCACHE_access_token"
	bm.Put(cacheName, "1asfahsfhqwhfqhw", 100)
	res := bm.Get(cacheName)
	fmt.Println(res.(string))
	getTokenUrl := AccessTokenUrl + "?grant_type=client_credential&appid=" + wx.Appid + "&secret=" + wx.Secret
	resp, err := http.Get(getTokenUrl)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		res := make(map[string]string)
		err = json.Unmarshal(body, &res)
		//	wx.saveCache("access_token", res["access_token"])
		return res["access_token"]
	} else {
		return ""
	}
	//fmt.Println(string(body), err)
}

//缓存存储token，默认7200秒，所以判断时间接近7200秒的时候进行重新获取工作，试用文件作为缓存存储,获取token名称存储
// func (wx *Wechat) saveCache(tokeyName, tokenValue string) bool {
// 	cacheName := "GTCACHE_" + tokeyName
// 	bm, err := cache.NewCache("file", `{"CachePath":"./cache","FileSuffix":".cache","DirectoryLevel":2,"EmbedExpiry":TokenCacheTime}`)
// 	if err == nil {
// 		err := bm.Put(cacheName, tokenValue, TokenCacheTime)
// 		if err == nil {
// 			return true
// 		} else {
// 			return true
// 		}
// 	} else {
// 		return false
// 	}

// }

//根据key获取缓存
// func (wx *Wechat) getCache(tokeyName string) string {
// 	bm, _ := cache.NewCache("file", `{"CachePath":"./cache","FileSuffix":".cache","DirectoryLevel":2,"EmbedExpiry":TokenCacheTime}`)
// 	cacheName := "GTCACHE_" + tokeyName
// 	bm.Get(cacheName)
// 	fmt.Println(bm)
// 	return ""
// 	// if err == nil {
// 	// 	cacheName := "GTCACHE_" + tokeyName
// 	// 	cacheV := bm.Get(cacheName)
// 	// 	fmt.Println(cacheV)
// 	// 	return ""
// 	// } else {
// 	// 	return ""
// 	// }
// }

/*func (wx *Wechat) GetResponse(url string) []map[string]string {

}

func (wx *Wechat) ErrorCode(errcode string) []map[string]string {

}
*/
