package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type BlockcountController struct {
	beego.Controller
}

func (b *BlockcountController)Get()  {
	b.TplName="count.html"

}

func (b *BlockcountController) Post() {
	u := User{}
	err := b.ParseForm(&u)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	b.Data["Count"]= u.Count
	b.TplName = "count.html"
}



