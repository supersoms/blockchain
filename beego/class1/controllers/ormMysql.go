package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"class1/models"
	"strconv"
)

type OrmController struct {
	beego.Controller
}

/******
	one action insert more data
 */
func (c *OrmController) OrmInsert() {
	o := orm.NewOrm()
	for i := 0; i < 5; i++ {
		user := models.User{Name: "supersom" + strconv.Itoa(i), Age: (33 + i), Address: "Shenzhen" + strconv.Itoa(i)}
		_, err := o.Insert(&user)
		if err != nil {
			beego.Info("insert data error:", err)
			return
		}
	}
	c.Ctx.WriteString("insert data successful!\n")
}

/******
	TODO how query all data ? mingtian
 */
func (c *OrmController) OrmQuery() {
	o := orm.NewOrm()
	user := models.User{Id: 10}
	err := o.Read(&user)
	if err != nil {
		beego.Info("query table error:", err)
		return
	}
	beego.Info("age=", user.Age, "name=", user.Name, "address=", user.Address)
	c.Ctx.WriteString("query data successful!\n")
}

/******
	TODO mingtian
	update one line data is ok
	update more line data is error
 */
func (c *OrmController) OrmUpdate() {
	o := orm.NewOrm()
	user := models.User{Name: "supersom1"}
	err := o.Read(&user)
	if err != nil {
		beego.Info("query table error:", err)
		return
	}
	user.Age = 19
	user.Name = "zhangsan"
	user.Address = "bj"
	updateNum, err := o.Update(&user)
	//_, err := o.QueryTable("User_order").Filter("user__id", user.Id).All(&orders)
	if err != nil {
		beego.Info("update table data error:", err)
		return
	}
	beego.Info("update table data num=", updateNum)
	c.Ctx.WriteString("update table data successful! \n")
}

func (c *OrmController) OrmDelete() {
	o := orm.NewOrm()
	user := models.User{Id: 1}
	_, err := o.Delete(&user)
	if err != nil {
		beego.Info("delete data error:", err)
		return
	}
	c.Ctx.WriteString("delete data successful!\n")
}
