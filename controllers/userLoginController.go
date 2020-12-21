package controllers

import (
	"BitcoinConnectionWeb/models"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

//直接访问login.html页面的请求
func (l *LoginController) Get(){
	l.TplName = "login.html"
}

//用户登录接口
func (l *LoginController) Post() {
	var user models.User
	err := l.ParseForm(&user)
	if err != nil {
		l.Ctx.WriteString("数据发生错误，请重试！")
		return
	}
	//查询数据库的用户信息
	u, err := user.QueryUser()
	if err != nil {
		l.Ctx.WriteString("用户数据查询失败，请重试！")
		return
	}

	//登录成功，跳转项目核心功能页面(home.html)
	l.Data["Phone"] = u.Phone
	l.TplName = "home.html"
}
