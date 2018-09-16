package models

import "github.com/astaxie/beego/orm"

type User struct {
	Id      int
	Name    string
	Age     int
	Address string
}

func Init() {
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/class1?charset=utf8")
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default", false, true)
}
