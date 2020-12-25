package main

import (
	"BeegoBcRPCCode/db_mysql"
	_ "BeegoBcRPCCode/routers"
	"BeegoBcRPCCode/utils"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {

	//连接数据库
	db_mysql.ConnectDB()

	//获取最新区块 hash
	BlockHash:=utils.GetBestBlockHash()
	fmt.Println("最新区块的哈希：",BlockHash)

    //设置静态资源路径
    beego.SetStaticPath("js","./static/js")
    beego.SetStaticPath("css","./static/css")
    beego.SetStaticPath("img","./static/img")
	beego.Run()
}

