package main

import (
	_ "class1/routers"
	"github.com/astaxie/beego"
	"class1/models"
)

func main() {
	models.Init()
	beego.Run()
}