package main

import (
	_ "class3/routers"
	"github.com/astaxie/beego"
	"class3/models"
	"strconv"
)

func main() {
	models.Init()
	beego.AddFuncMap("ShowPrvPage", HandlerPrvPage) //index页面上的ShowPrvPage函数名与HandlerPrvPage函数进行绑定映射
	beego.AddFuncMap("ShowNextPage", HandlerNextPage)
	beego.Run()
}

/******
	点击时处理上一页逻辑
	data：就是页面{{.pageIndex | ShowPrvPage}}上传过来的pageIndex
	return：返回值string|int都可以,页面上不区分类型
 */
func HandlerPrvPage(data int) string {
	pageIndex := data - 1
	return strconv.Itoa(pageIndex)
}

/******
	点击时处理下一页逻辑
	data：就是页面{{.pageIndex | ShowNextPage}}上传过来的pageIndex
	return：返回值string|int都可以,页面上不区分类型
 */
func HandlerNextPage(data int) string {
	pageIndex := data + 1
	return strconv.Itoa(pageIndex)
}
