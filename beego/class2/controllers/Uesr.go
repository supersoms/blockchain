package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"class2/models"
)

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) ShowRegisterPage() {
	this.TplName = "user_register.html"
}

func (this *RegisterController) HandlerUserRegister() {
	//1.获取浏览器控件上的用户名和密码
	userName := this.GetString("userName")
	password := this.GetString("password")
	if userName == "" || password == "" {
		beego.Error("用户名或密码不能为空")
		this.TplName = "user_register.html"
		return
	}
	//1.获取ORM对象
	o := orm.NewOrm()
	//2.获取插入对象
	user := models.User{UserName: userName, PassWord: password}
	//3.插入操作
	_, err := o.Insert(&user)
	if err != nil {
		beego.Error("用户注册失败!", err)
		return
	}
	//4.跳转到登陆页面
	this.Redirect("/", 302)
}

type LoginController struct {
	beego.Controller
}

func (this *LoginController) ShowLoginPage() {
	this.TplName = "user_login.html"
}

//处理用户登陆
func (this *LoginController) HandlerLogin() {
	//1 拿到浏览器的用户名和密码
	userName := this.GetString("userName")
	password := this.GetString("password")
	beego.Info(userName)
	beego.Info(password)
	//2 数据判断
	if userName == "" || password == "" {
		beego.Error("用户名或者密码不能为空")
		this.TplName = "user_login.html"
		return
	}
	//3 查找用户数据
	o := orm.NewOrm()
	user := models.User{UserName: userName}
	err := o.Read(&user, "user_name")
	if err != nil {
		beego.Error("用户不存在!", err)
		this.TplName = "user_login.html"
		return
	}
	if user.PassWord != password {
		beego.Error("密码错误!")
		this.TplName = "user_login.html"
		return
	}
	beego.Info("UserName=", user.UserName, "PassWord=", user.PassWord)
	//4 登录成功跳转到文章列表页面
	this.Redirect("/showArticle", 302)
}
