package routers

import (
	"BeegoBcRPCCode/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//默认页面登入注册界面
	beego.Router("/", &controllers.MainController{})
	//登入页面
	beego.Router("/login", &controllers.RegisterController{})
	beego.Router("/register.html", &controllers.RegisterController{})
	//进入主页面
	beego.Router("/home",&controllers.LoginController{})
	//进入Blockcount页面
	beego.Router("/count",&controllers.BlockcountController{})
	beego.Router("/count.html",&controllers.BlockcountController{})
	//进入信息页
	beego.Router("/info",&controllers.InfoControllers{})
}
