package controllers

import (
	"github.com/astaxie/beego"
	"github.com/php0532/gotocms/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	wx := &models.Wechat{
		Appid:  "wx6b888d9c51a16d7f",
		Secret: "2437653a08aea9f7945dddc75a81f645",
	}
	str := wx.CreateToken()
	c.Ctx.WriteString(str)
	//	c.TplNames = "login.tpl"
}

func (c *MainController) Login() {
	c.TplNames = "login.tpl"
}
