package model

import (
	"testing"
)

/**
构造体 类似于实体  项目中通常作为 数据库返回数据被装载使用  (Service层使用，将实体给与Controller)

项目中通常使用beego的Rom工具，将SQL查询出来的信息直接指向 &User 的指针地址 （返回参数是[]User 创建：user := []User  QueryRows(&user)）
*/
type User struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Age     int64  `json:"age"`
	Address string `json:"address"`
}

func TestModel(t *testing.T) {
	var user User
	user.Id = 1
	user.Name = "LuTong"
	user.Age = 23
	user.Address = "City"
}
