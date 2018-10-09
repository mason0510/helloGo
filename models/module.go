package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id int //`orm`
	Name string `orm:size(100)`
	//多级查询 一对多 一对多
	User_order []*User_order `orm:"reverse(many)"`
}
type User_order struct{
	Id         int
	Order_data string `orm:size(100)`
	User       *User  `orm:"rel(fk)"`
}


func init()  {
	orm.RegisterDataBase("default","mysql","root:123456@tcp(127.0.0.1:3306)/class27?charset=utf8",30)
	//register
	orm.RegisterModel(new(User),new (User_order))
	//create table
	orm.RunSyncdb("default",false,true)
}
