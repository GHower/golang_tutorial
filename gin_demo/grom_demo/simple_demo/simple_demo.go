package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID   int
	Name string
	Age  int
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) //打开数据库连接
	if err != nil {
		panic(err)
	}
	DB, _ := db.DB()
	defer DB.Close() //在函数执行完成后关闭数据库句柄

	//新增
	rsOne := db.Create(&User{ID: 1, Name: "aaaaaaaaaa", Age: 32})
	if rsOne.RowsAffected > 0 {
		fmt.Println("用户1创建成功")
	}

	//rsTwo := db.Create(&User{ID: 2, Name: "bbbbb", Age: 32})
	//if rsTwo.RowsAffected > 0 {
	//	fmt.Println("用户2创建成功")
	//}

	//查询
	var u1 User
	db.Find(&u1, "id = ?", 1)
	fmt.Println(u1)

	//批量查询
	var users []User
	db.Find(&users)
	fmt.Println(users)

	//修改
	u1.Name = "11111111"
	rsSave := db.Save(u1)
	if rsSave.RowsAffected > 0 {
		fmt.Println("用户创建成功")
	}
	//修改后查询数据是否已经变化
	db.Debug().Find(&u1)
	fmt.Println(u1)

	//删除
	rsDelete := db.Delete(u1)
	if rsDelete.RowsAffected > 0 {
		fmt.Println("用户删除成功")
	}
}
