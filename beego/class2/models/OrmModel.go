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
}

//文章表字段信息
type Article struct {
	Id      int
	Title   string `orm:"size(100)"`     //文章标题长度为20
	Content string `orm:"size(10000)"`   //文章内容
	Image   string `orm:"size(50);null"` //图片路径
	//Type      string    									//类型
	Time      time.Time `orm:"type(datetime);auto_now_add"` //发布时间
	ReadCount int       `orm:"default(0)"`                  //阅读量
}

func Init() {
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/newsWeb?charset=utf8")
	//创建表User和Article
	orm.RegisterModel(new(User), new(Article))
	//参数2表示是否强制更新,true：会删表重新建表
	orm.RunSyncdb("default", false, true)
}
