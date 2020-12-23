package controllers

import (
	"BeegoBcRPCCode/models"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (l*LoginController) Post() {
	var user models.User
	//解析用户的信息
	err:=l.ParseForm(&user)
	if err!=nil {
		l.Ctx.WriteString("对不起，用户信息解析失败，请重试！")
		return
	}
	//查询数据库的用户信息
	u,err:=user.QueryUser()
	if err!=nil {
		fmt.Println(err)
		l.Ctx.WriteString("用户登入失败，请重试！")
		return
	}
	//登入成功，跳转到项目的核心功能页面（home.html)
	l.Data["user_name"] = u.Username
	l.TplName = "home.html"


}
