package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["test"] = "區塊連" //
	c.TplName = "test.html"
}

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Post() {
	c.Data["data"] = "This is Post method！" //
	c.TplName = "test.html"
}

func (c *IndexController) ShowGet() {
	path := c.GetString(":path")
	ext := c.GetString(":ext")
	beego.Info(path, ext)

	c.Data["test"] = "ShowGet method" //
	c.TplName = "test.html"
}
