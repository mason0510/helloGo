package main

import (
	_ "hello/routers"
	"github.com/astaxie/beego"
"hello/models"
	"github.com/astaxie/beego/orm"
	"fmt"
)

func main() {

	//set static path
	//beego.SetStaticPath("down1","down")
	//insertOrder()
	//insertUser()
	queryOrder(2)

	beego.Run()

	//queryUser()
	//userUpdate()
	//deleteUser()


}
func insertUser()  {
	o:=orm.NewOrm()
	user:=models.User{}
	user.Name="zxc"
	id,err:=o.Insert(&user)
	if err!=nil {
		beego.Info("inser erro")
	}
	beego.Info("insert success id=",id)
}

func queryUser()  {
	o:=orm.NewOrm()
	user:=models.User{Id:1}
	err:=o.Read(&user)
	if err!=nil {
		beego.Info("query is err")
		return
	}
	beego.Info("query success user=",user)
}
func userUpdate()  {
	o:=orm.NewOrm()
	user:=models.User{Id:1,Name:"zhangxiaocong"}
	_,err:=o.Update(&user)
	if err!=nil {
		beego.Info("update err")
		return
	}else {
		beego.Info("delete err")
	}
	beego.Info("success update " )
	}
func deleteUser()  {
	o:=orm.NewOrm()
	user:=models.User{Id:1}
	_,err:=o.Delete(&user)
	if err!=nil {
		beego.Error("delede err")
		return
	}
	beego.Info("delete success")
}
//func deleteUser()  {
//	o:=orm.NewOrm()
//	user:=models.User{Id:2}
//	_,err:=o.Delete(&user)
//	if err!=nil {
//		beego.Error("delete err")
//		return
//	}
//	beego.Error("delete success")
//}
func insertOrder()  {
	o:=orm.NewOrm()
	order:=models.User_order{}
	order.Order_data="this is order"
	user:=models.User{Id:3}
	order.User=&user
	id,err:=o.Insert(&order)
	if err!=nil {
		beego.Info("inser erro",id)
	}
	beego.Info("insert success id=",id)
}

func queryOrder(id int) {
	var orders []*models.User_order
	orm.Debug = true
	o := orm.NewOrm()
	qs := o.QueryTable("User_order")
	//qs_user := o.QueryTable("User")
	//fmt.Printf(qs_user.All())
	order_num, err := qs.Filter("user",id).All(&orders)
//	order_num1, err := qs.All(&orders)
//	beego.Info("order_number1=",order_num1)
	if err!=nil {
		beego.Error("query error")
		return
	}
	beego.Info("query order num=",order_num)
	for _,order:=range orders{
		fmt.Printf("query order success order= %#v",order)
		//beego.Info("query order success order= %#v",order)
	}
	}
//func TransparentStatic(ctx *context.Context)  {
//	orpath:=ctx.Request.URL.Path
//	beego.Debug("request url:",orpath)
//	if string.Index(orpath,"api")>=0 {
//		return
//	}
//	http.ServeFile(ctx.ResponseWriter,ctx.Request,"static/html/"+ctx.Request.URL.Path)
//
//	}
//func ignoreStaticPath()  {
//	beego.InsertFilter("/",beego.BeforeRouter,TransparentStatic)
//	beego.InsertFilter("/*",beego.BeforeRouter,TransparentStatic)
//}