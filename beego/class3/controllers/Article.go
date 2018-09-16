package controllers

import (
	"github.com/astaxie/beego"
	"path"
	"time"
	"github.com/astaxie/beego/orm"
	"class3/models"
	"strconv"
	"os"
	"math"
)

type ArticleController struct {
	beego.Controller
}

//分页显示文章列表
func (this *ArticleController) ShowArticleList() {
	//1 查询
	orm := orm.NewOrm()
	qs := orm.QueryTable("Article").OrderBy("-time") //按数据库字段名time降序查詢Article表的所有信息, -表示降序,不加表示升序
	var articles [] models.Article

	pageIndex, err := strconv.Atoi(this.GetString("pageIndex"))
	if err != nil {
		pageIndex = 1 //默认是首页
	}

	typeName := this.GetString("select")
	beego.Info("文章类型名称为：", typeName)
	var totalCount int64
	//TODO 只返回文章类型为(typeName)的总数据条数
	if typeName == "" {
		totalCount, err = qs.RelatedSel("ArticleType").Count()
	} else {
		//Filter("ArticleType__TypeName", typeName)函数是指定过滤结构体ArticleType中的属性TypeName对应表中相同的字段为typeName的数据,比如过滤A,那就显示A的数据
		totalCount, err = qs.RelatedSel("ArticleType").Filter("ArticleType__TypeName", typeName).Count()
	}
	if err != nil {
		beego.Error("查询错误")
		return
	}
	beego.Info("文章总数据为：", totalCount)

	//TODO 2 计算总页数
	pageSize := 10                                                     //设置每页显示1条数据
	start := pageSize * (pageIndex - 1)                                //计算起始位置
	qs.Limit(pageSize, start).RelatedSel("ArticleType").All(&articles) //存放(select * from Article)SQL查询语句返回的所有数据
	pageCount := math.Ceil(float64(totalCount) / float64(pageSize))    //总数据/每页显示数据,然后向上取整计算总页数

	//TODO 3 处理当点击上一页到了首页和下一页到了末页的数据
	firstPage := false //标识是否是首页,默认时,上一页按钮可以点击
	if pageIndex == 1 {
		firstPage = true //如果上一页到了首页时按钮不让点击
	}
	endPage := false //标识是否是末页,默认时,下一页按钮可以点击
	if pageIndex == int(pageCount) {
		endPage = true //如果下一页到了总页数-1时按钮不让点击
	}

	//TODO 获取类型
	var articleTypes []models.ArticleType
	orm.QueryTable("ArticleType").All(&articleTypes)
	this.Data["articleTypes"] = articleTypes

	//TODO 根据文章类型获取文章列表
	var articlesWithType []models.Article
	if typeName == "" {
		beego.Info("下拉框传递数据失败")
		//获取所有带文章类型的数据
		datasize, err := qs.Limit(pageSize, start).RelatedSel("ArticleType").All(&articlesWithType)
		if err != nil {
			beego.Error("获取所有的文章列表失败!")
		}
		beego.Info("if 数据大小：", datasize)
	} else {
		//根据文章类型获取文章列表
		datasize, err := qs.Limit(pageSize, start).RelatedSel("ArticleType").Filter("ArticleType__TypeName", typeName).All(&articlesWithType)
		if err != nil {
			beego.Error("根据文章类型查询文章列表数据失败!")
		}
		beego.Info("else 数据大小：", datasize)
	}
	//2 把数据绑定到文章列表页面显示
	this.Data["typeName"] = typeName
	this.Data["firstPage"] = firstPage
	this.Data["endPage"] = endPage
	this.Data["totalCount"] = totalCount
	this.Data["pageCount"] = pageCount
	this.Data["pageIndex"] = pageIndex
	if len(articlesWithType) > 0 { //只有当文章列表数大于0时,才传递数据
		this.Data["articles"] = articlesWithType
	}
	this.TplName = "article_index.html"
}

//TODO 处理选择文章类型查询数据,此函数没有调用
func (this *ArticleController) HandlerSeletctArticleType() {
	typeName := this.GetString("select")
	beego.Info(typeName)
	if typeName == "" {
		beego.Error("下拉框传递数据失败!")
		return
	}
	orm := orm.NewOrm()
	var articles [] models.Article
	//Filter("结构体表名__字段",value)
	//RelatedSel惰性查询
	_, err := orm.QueryTable("Article").RelatedSel("ArticleType").Filter("ArticleType__TypeName", typeName).All(&articles)
	if err != nil {
		beego.Error("根据文章类型查询文章列表数据失败!")
		return
	}
	beego.Error(articles)
	this.Redirect("showArticle", 302)
}

func (this *ArticleController) ShowAddArticle() {
	//TODO 获取类型
	orm := orm.NewOrm()
	var articleTypes []models.ArticleType
	orm.QueryTable("ArticleType").All(&articleTypes)
	this.Data["articleTypes"] = articleTypes
	this.TplName = "article_add.html"
}

//添加文章到数据库
func (this *ArticleController) HandlerAddArticle() {
	articleTitle := this.GetString("articleName")   //标题
	articleTypeName := this.GetString("select")     //文章类型
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
	} else if articleTypeName == "" {
		beego.Error("下拉选择文章类型不能为空!")
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
	var articleType models.ArticleType
	articleType.TypeName = articleTypeName
	err = orm.Read(&articleType, "TypeName")
	if err != nil {
		beego.Error("获取类型错误!")
		return
	}
	//TODO 将articleType存到article表中
	article := models.Article{Title: articleTitle, Content: articleContent, Image: imageFilePath, Time: time.Now(), ArticleType: &articleType}
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

func (this *ArticleController) AddArticleType() {
	//1.获取ORM对象
	orm := orm.NewOrm()
	//2.获取插入对象
	var articleTypes []models.ArticleType
	_, err := orm.QueryTable("ArticleType").All(&articleTypes)
	if err != nil {
		beego.Error("qurey fail")
	}
	this.Data["articles"] = articleTypes
	//3.返回到显示文章页面
	this.TplName = "article_addType.html"
}

func (this *ArticleController) HandlerAddArticleType() {
	typeName := this.GetString("typeName")
	if typeName == "" {
		beego.Error("添加类型数据为空!")
		return
	}
	orm := orm.NewOrm()
	articleType := models.ArticleType{TypeName: typeName}
	_, err := orm.Insert(&articleType)
	if err != nil {
		beego.Error("添加文章类型失败!")
		return
	}
	this.Redirect("/addArticleType", 303)
}
