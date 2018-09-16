package routers

import (
	"class3/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.RegisterController{}, "get:ShowRegisterPage;post:HandlerUserRegister")        //显示注册页面,处理注册功能
	beego.Router("/login", &controllers.LoginController{}, "get:ShowLoginPage;post:HandlerLogin")                        //显示登录页面,处理登录功能
	beego.Router("/showArticle", &controllers.ArticleController{}, "get:ShowArticleList;post:HandlerSeletctArticleType") //显示文章列表
	beego.Router("/addArticle", &controllers.ArticleController{}, "get:ShowAddArticle;post:HandlerAddArticle")           //添加文章
	beego.Router("/articleContent", &controllers.ArticleController{}, "get:ShowArticleDetailContent")                    //显示文章详情内容
	beego.Router("/deleteArticle", &controllers.ArticleController{}, "get:HandlerDeleteArticle")                         //删除文章
	beego.Router("/showEditArticle", &controllers.ArticleController{}, "get:ShowEditArticle")                            //显示编辑文章
	beego.Router("/reEditSubmitArticle", &controllers.ArticleController{}, "post:ReEditSubmitArticle")                   //重新编辑文章再提交文章
	beego.Router("/addArticleType", &controllers.ArticleController{}, "get:AddArticleType;post:HandlerAddArticleType")   //重新编辑文章再提交文章
}
