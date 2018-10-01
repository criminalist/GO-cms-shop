package routers

import (
	"github.com/astaxie/beego"
	"github.com/criminalist/GO-cms-shop/controllers"
	"github.com/criminalist/GO-cms-shop/utils/urlpath"
)

func init() {
	memberCtl := &controllers.MemberController{}
	//添加用户收获地址
	beego.Router(urlpath.Path("member_add_address"), memberCtl, "get:GetAddAddress")
	//添加用户收获地址
	beego.Router(urlpath.Path("member_add_address"), memberCtl, "post:PostAddAddress")
	//编辑用户地址
	beego.Router(urlpath.Path("member_edit_address"), memberCtl, "get:GetEditAddress")

}
