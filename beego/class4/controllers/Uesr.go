package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"class4/models"
	"time"
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
	userName := this.Ctx.GetCookie("userName")
	if userName != "" {
		this.Data["userName"] = userName
		this.Data["checked"] = "checked"
	}
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

	//设置cookie
	check := this.GetString("remember")
	if check != "" && check == "on" {
		//cookie的有效期是30分钟
		this.Ctx.SetCookie("userName", user.UserName, time.Second*1800)
	} else {
		//删除cookie
		this.Ctx.SetCookie("userName", "", -1)
	}
	//设置session用于后面的页面进行判断是否进入相关页面
	this.SetSession("userName", user.UserName)

	//4 登录成功跳转到文章列表页面
	this.Redirect("/article/showArticle", 302)
}

//退出功能
func (this *LoginController) Logout() {
	this.DelSession("userName")  //删除Session
	this.Redirect("/login", 302) //重定向到登录页面
}
