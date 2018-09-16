package routers

import (
	"class4/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	//过滤器匹配/article/后面的所有路径,执行FilterFunc方法
	beego.InsertFilter("/article/*", beego.BeforeRouter, FilterFunc)
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.RegisterController{}, "get:ShowRegisterPage;post:HandlerUserRegister")                //显示注册页面,处理注册功能
	beego.Router("/login", &controllers.LoginController{}, "get:ShowLoginPage;post:HandlerLogin")                                //显示登录页面,处理登录功能
	beego.Router("/article/showArticle", &controllers.ArticleController{}, "get:ShowArticleList;post:HandlerSeletctArticleType") //显示文章列表
	beego.Router("/article/addArticle", &controllers.ArticleController{}, "get:ShowAddArticle;post:HandlerAddArticle")           //添加文章
	beego.Router("/article/articleContent", &controllers.ArticleController{}, "get:ShowArticleDetailContent")                    //显示文章详情内容
	beego.Router("/article/deleteArticle", &controllers.ArticleController{}, "get:HandlerDeleteArticle")                         //删除文章
	beego.Router("/article/showEditArticle", &controllers.ArticleController{}, "get:ShowEditArticle")                            //显示编辑文章
	beego.Router("/article/reEditSubmitArticle", &controllers.ArticleController{}, "post:ReEditSubmitArticle")                   //重新编辑文章再提交文章
	beego.Router("/article/addArticleType", &controllers.ArticleController{}, "get:AddArticleType;post:HandlerAddArticleType")   //重新编辑文章再提交文章
	beego.Router("/article/deleteArticleType", &controllers.ArticleController{}, "get:DeleteArticleType")                        //重新编辑文章再提交文章
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")                                                        //退出
}

var FilterFunc = func(cxt *context.Context) {
	//从session中获取用户名
	userName := cxt.Input.Session("userName")
	if userName == nil {
		//如果session中的用户名为nil,跳转到登录界面
		cxt.Redirect(302, "/login")
	}
}
