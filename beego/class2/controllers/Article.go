package controllers

import (
	"github.com/astaxie/beego"
	"path"
	"time"
	"github.com/astaxie/beego/orm"
	"class2/models"
	"strconv"
	"os"
)

type ArticleController struct {
	beego.Controller
}

//显示文章列表
func (this *ArticleController) ShowArticleList() {
	//1 查询
	orm := orm.NewOrm()
	qs := orm.QueryTable("Article") //查詢Article表
	var articles [] models.Article
	qs.All(&articles) //存放select * from Article查询所有的数据
	if len(articles) > 0 { //只有当文章列表数大于0时,才传递数据
		beego.Info(articles[0].Title)
		this.Data["articles"] = articles
	}
	//2 把数据绑定到文章列表页面显示
	this.TplName = "article_index.html"
}

func (this *ArticleController) ShowAddArticle() {
	this.TplName = "article_add.html"
}

//添加文章
func (this *ArticleController) HandlerAddArticle() {
	articleTitle := this.GetString("articleName") //标题
	//articleType := this.GetString("select")         //文章类型
	articleContent := this.GetString("content")     //内容
	file, header, err := this.GetFile("uploadname") //获取上传的静态文件
	if err != nil {
		beego.Error("error info=", err)
		return
	}
	defer file.Close()
	if articleTitle == "" {
		beego.Error("文章标题不能为空!")
		this.TplName = "article_add.html"
		return
	} else if articleContent == "" {
		beego.Error("文章内容不能为空!")
		this.TplName = "article_add.html"
		return
	} else if header.Filename == "" {
		beego.Error("请选择待上传的图片!")
		this.TplName = "article_add.html"
		return
	}
	//1 处理文件格式
	ext := path.Ext(header.Filename) //获取文件的后缀
	if ext != ".jpg" && ext != ".png" && ext != "jpeg" {
		beego.Error("上传文件的格式不正确!")
		return
	}
	//2 判断文件大小
	if header.Size > 5000000 {
		beego.Error("上传的文件太大,不允许上传!")
		return
	}
	//3 处理文件不能重名
	fileName := time.Now().Format("2006-01-02 15:04:05") //此格式字符串是固定的
	imageFilePath := getImageDir() + fileName + ext
	if err != nil {
		beego.Error("上传文件失败!", err)
		return
	}
	beego.Info(articleTitle, articleContent, fileName+ext)

	//1.获取ORM对象
	orm := orm.NewOrm()
	//2.获取插入对象
	article := models.Article{Title: articleTitle, Content: articleContent, Image: imageFilePath, Time: time.Now()}
	//3.插入操作
	_, err = orm.Insert(&article)
	if err != nil {
		beego.Error("添加文章失败!", err)
		return
	}
	//添加文章成功之后再上传图片
	this.SaveToFile("uploadname", imageFilePath) //保存文件到/static/updateImg目录下
	//4.返回到显示文章页面
	this.Redirect("showArticle", 302)
}

//显示文章详情内容
func (this *ArticleController) ShowArticleDetailContent() {
	//1 获取Id
	id := this.GetString("id")
	//2 查询数据
	//2.1 获取orm对象
	orm := orm.NewOrm()
	//2.2 获取查询对象
	id1, _ := strconv.Atoi(id)
	article := models.Article{Id: id1}
	err := orm.Read(&article)
	if err != nil {
		beego.Error("查询数据为空", err)
		return
	}
	article.ReadCount += 1
	orm.Update(&article) //更新阅读数
	//3 传递数据给视图
	this.Data["article"] = article
	this.TplName = "article_content.html"
}

//根据id删除文章
func (this *ArticleController) HandlerDeleteArticle() {
	//1 从页面上获取Id
	id := this.GetString("id")
	//2.1 获取orm对象
	orm := orm.NewOrm()
	//2.2 获取查询对象
	id1, _ := strconv.Atoi(id)
	article := models.Article{Id: id1}
	orm.Delete(&article)
	//3 传递数据给视图
	this.Redirect("/showArticle", 302)
}

//显示编辑文章
func (this *ArticleController) ShowEditArticle() {
	//1 获取Id
	id := this.GetString("id")
	//2 查询数据
	//2.1 获取orm对象
	orm := orm.NewOrm()
	//2.2 获取查询对象
	id1, _ := strconv.Atoi(id)
	article := models.Article{Id: id1}
	err := orm.Read(&article)
	if err != nil {
		beego.Error("查询数据为空!", err)
		return
	}
	//3 传递数据给视图
	this.Data["article"] = article
	//4 跳转到编辑文章UI界面
	this.TplName = "article_update.html"
}

//重新编辑提交文章
func (this *ArticleController) ReEditSubmitArticle() {
	idStr := this.GetString("id")                   //标题
	articleTitle := this.GetString("articleName")   //标题
	articleContent := this.GetString("content")     //内容
	file, header, err := this.GetFile("uploadname") //获取上传的静态文件
	if err != nil {
		beego.Error("error info=", err)
		return
	}
	if articleTitle == "" {
		beego.Error("文章标题不能为空!")
		this.TplName = "article_update.html"
		return
	} else if articleContent == "" {
		beego.Error("文章内容不能为空!")
		this.TplName = "article_update.html"
		return
	}

	imageFilePath := "" //有可能这个图片用户没有替换,所以要进行判断
	if header != nil && header.Filename != "" {
		defer file.Close()
		//1 处理文件格式
		ext := path.Ext(header.Filename) //获取文件的后缀
		if ext != ".jpg" && ext != ".png" && ext != "jpeg" {
			beego.Error("上传文件的格式不正确!")
			return
		}
		//2 判断文件大小
		if header.Size > 5000000 {
			beego.Error("上传的文件太大,不允许上传!")
			return
		}
		//3 处理文件不能重名
		fileName := time.Now().Format("2006-01-02 15:04:05") //此格式字符串是固定的
		imageFilePath = getImageDir() + fileName + ext
		if err != nil {
			beego.Error("上传文件失败!", err)
			return
		}
	}

	//1.获取ORM对象
	orm := orm.NewOrm()
	//2.获取更新对象
	id, err := strconv.Atoi(idStr)
	article := models.Article{Id: id}
	//3.先用read查询看有没有此文章,如果不用read先查询,也可以直接用update更新,但是一定要根据主键id去update,不然更新失败!
	err = orm.Read(&article)
	if err != nil {
		beego.Error("要更新的文章不存在!")
		return
	}
	article.Id = id
	article.Title = articleTitle
	article.Content = articleContent
	if imageFilePath != "" {
		article.Image = imageFilePath
	}
	article.Time = time.Now()
	beego.Info("显示待更新article信息:", article)

	//4.更新操作
	_, err = orm.Update(&article)
	if err != nil {
		beego.Error("重新编辑文章失败!", err)
		return
	}
	if imageFilePath != "" {
		//重新编辑文章成功之后再上传图片
		this.SaveToFile("uploadname", imageFilePath) //保存文件到/static/updateImg目录下
	}
	//5.重定向到显示文章列表页面
	this.Redirect("showArticle", 302)
}

func getImageDir() string {
	imageDir := "./static/updateImg/"
	_, error := os.Stat(imageDir)
	if error != nil && os.IsNotExist(error) { //TODO 如果图片目录不存在，创建一个新的目录
		os.Mkdir(imageDir, 0777)
	}
	return imageDir
}
