package main

import (
	_ "class2/routers"
	"github.com/astaxie/beego"
	"class2/models"
)

func main() {
	models.Init()
	beego.Run()
}
