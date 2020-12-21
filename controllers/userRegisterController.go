package controllers

import (
	"BitcoinConnectionWeb/models"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

//直接请求用户注册页面
func (r *RegisterController) Get() {
	r.TplName = "register.html"
}

//用户注册的表单提交请求
func (r *RegisterController) Post(){
	//1、解析请求数据
	var user models.User
	err := r.ParseForm(&user)
	if err != nil {
		r.Ctx.WriteString("解析错误")
		return
	}
	//2、保存用户信息到数据库
	_, err = user.SaveUser()
	if err != nil {
		r.Ctx.WriteString("数据错误")
		return
	}
	//注册成功
	r.TplName = "login.html"
}
