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
	TokenSaveTime  int    = 7000                                      //token的存活存储时间，官方为7200秒所以我们
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
		return res["access_token"]
	} else {
		return ""
	}
	//fmt.Println(string(body), err)
}

//缓存存储token，默认7200秒，所以判断时间接近7200秒的时候进行重新获取工作
func (wx *Wechat) SaveToken() {

}

/*func (wx *Wechat) GetResponse(url string) []map[string]string {

}

func (wx *Wechat) ErrorCode(errcode string) []map[string]string {

}
*/
