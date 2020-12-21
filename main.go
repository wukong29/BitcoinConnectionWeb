package main

import (
	"BitcoinConnectionWeb/db_mysql"
	_ "BitcoinConnectionWeb/routers"
	"github.com/astaxie/beego"
)

func main() {
	//1、链接数据库
	db_mysql.ConnerDB()

	//静态资源路径设置
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")

	beego.Run()
}

