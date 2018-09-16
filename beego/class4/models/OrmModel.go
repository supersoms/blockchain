package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//用户表字段信息
type User struct {
	Id       int
	UserName string
	PassWord string
	Articles [] *Article `orm:"rel(m2m)"` //TODO 创建一个多对多的关系表,表示一个用户可以有多个文章
}

//文章表字段信息
type Article struct {
	Id          int
	Title       string       `orm:"size(100)"`                   //文章标题长度为20
	Content     string       `orm:"size(10000)"`                 //文章内容
	Image       string       `orm:"size(50);null"`               //图片路径
	Time        time.Time    `orm:"type(datetime);auto_now_add"` //发布时间
	ReadCount   int          `orm:"default(0)"`                  //阅读量
	ArticleType *ArticleType `orm:"rel(fk)"`                     //TODO rel(fk)外键,文章类型,表示一个文章有多个类型
	Users       [] *User     `orm:"reverse(many)"`               //TODO 多对多,表示一个文章有多个读者
}

/******
	文章类型表字段信息
	文章表和文章类型表是1对多
 */
type ArticleType struct {
	Id       int
	TypeName string      `orm:"size(20)"`
	Articles [] *Article `orm:"reverse(many)"`
}

func Init() {
	//往数据库表插入时间时,用本地时间,&loc=Local这句作用
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/newsWeb?charset=utf8&loc=Local")
	//创建表User,Article,ArticleType
	orm.RegisterModel(new(User), new(Article), new(ArticleType))
	//参数2表示是否强制更新,true：会删表重新建表
	orm.RunSyncdb("default", false, true)
}
