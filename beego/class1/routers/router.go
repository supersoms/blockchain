package routers

import (
	"class1/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/index", &controllers.IndexController{}, "Get:ShowGet;post:Post")
	beego.Router("/mysql", &controllers.MysqlController{}, "Get:ShowMysql")
	//beego.Router("/orm", &controllers.OrmController{},"Get:ShowOrm")
	beego.Router("/orminsert", &controllers.OrmController{}, "Get:OrmInsert")
	beego.Router("/ormquery", &controllers.OrmController{}, "Get:OrmQuery")
	beego.Router("/ormupdate", &controllers.OrmController{}, "Get:OrmUpdate")
	beego.Router("/ormdelete", &controllers.OrmController{}, "Get:OrmDelete")
}
