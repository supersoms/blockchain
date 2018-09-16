package controllers

import (
	"github.com/astaxie/beego"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

type MysqlController struct {
	beego.Controller
}

func (c *MysqlController) ShowMysql() {
	conn, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/class1?charset=utf8")
	if err != nil {
		beego.Info("conn error:", err)
		return
	}
	defer conn.Close()

	//createTable(c, conn)

	//insertData(c, conn)

	queryData(c, conn)
}

func createTable(c *MysqlController, conn *sql.DB) {
	//创建表时如果不加if not exists判断，当第2次执行创建语句时立马报错。
	_, err := conn.Exec("create table if not exists userinfo(id int,name varchar(15))")
	if err != nil {
		beego.Info("create table error:", err)
		return
	}
	c.Ctx.WriteString("create table successful!")
}

func insertData(c *MysqlController, conn *sql.DB) {
	_, err := conn.Exec("insert userinfo(id,name) values(?,?)", 3, "zhongxuemei")
	if err != nil {
		beego.Info("insert data error:", err)
		return
	}
	c.Ctx.WriteString("insert data successful!")
}

func queryData(c *MysqlController, conn *sql.DB) {
	rows, err := conn.Query("select id from userinfo")
	beego.Info(err)
	var id int
	for rows.Next() {
		rows.Scan(&id)
		c.Ctx.WriteString(strconv.Itoa(id))
	}
}
