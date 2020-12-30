package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type InfoControllers struct {
	beego.Controller
}

func (i *InfoControllers) Get() {
	i.TplName = "count.html"

}

func (i *InfoControllers) Post() {
	u := User{}
	err := i.ParseForm(&u)
	if err != nil {
		fmt.Println(err.Error())

	}


	//获取当前节点区块链总数

	fmt.Println(u.Count)
	fmt.Println(u.Arguments)
	fmt.Println(u.Arguments1)
	Info := BictoinJudge(u.Count, u.Arguments, u.Arguments1) //接收的是万能类型
	fmt.Println("======================", Info)
	i.Data["Info"] = Info
	i.TplName = "count.html"

	//fmt.Println("区块总数：",count)
	//varinfo := "此命令无需填参数"
	//b.Data["Varinfo"] =varinfo

}
//func Str(command string)string{
//	u:=User{}
//	if u.Count=="getblockcount"{
//     s:="当前区块的总数："
//     return s
//	}
//return ""
//}