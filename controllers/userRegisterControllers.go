package controllers

import (
	"BeegoBcRPCCode/models"
	"fmt"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

/**
处理用户注册的表单的提交请求
*/

func (r *RegisterController) Get() {
	fmt.Println("hello world")
	r.TplName = "register.html"
}

func (r *RegisterController) Post() {
	//1.解析数据请求
	var user models.User
	err := r.ParseForm(&user)
	//fmt.Println(user)
	if err != nil {
		//2.返回提示信息在浏览器给用户
		r.Ctx.WriteString("抱歉用户解析数据错误，请重试！")
		return
	}
	//保存数据到数据库
	_, err = user.SaveUser()
	if err != nil {

		r.Ctx.WriteString("抱歉用户注册失败，请重试！")
		fmt.Println(err)
		return
	}
	r.TplName = "login.html"

}
