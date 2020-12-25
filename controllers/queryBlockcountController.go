package controllers

import (
	"BeegoBcRPCCode/utils"
	"fmt"
	"github.com/astaxie/beego"
)

type BlockcountController struct {
	beego.Controller
}

func (b *BlockcountController) Get() {
	count:=utils.GetBlockcount()
	fmt.Println("区块总数：",count)
	b.Data["Count"]=count
	b.TplName = "count.html"
}

func (b *BlockcountController)Post(){
	//获取当前节点区块链总数



}

