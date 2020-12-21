package controllers

import (
	"BitcoinConnectionWeb/utils"
	"fmt"
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (l *HomeController) Get() {
	l.TplName = "home.html"
}

func (l *HomeController) Post() {
	tx := utils.GetTxOutSetInfo()
	fmt.Println(tx.Bestblock)
}
