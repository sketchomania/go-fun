package main

import (
	"fmt"
	"log"
	"os"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

// Model Struct
type User struct {
	Id   int
	Name string `orm:"size(100)"`
}

func init() {
	// register model
	orm.RegisterModel(new(User))
}

func main() {
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		log.Fatal("MYSQL_DSN must be set")
	}
	if err := orm.RegisterDataBase("default", "mysql", dsn); err != nil {
		log.Fatal(err)
	}

	o := orm.NewOrm()

	user := User{Name: "Selena", Id: 0}

	// insert
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// // update
	// user.Name = "astaxie"
	// num, err := o.Update(&user)
	// fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// o := orm.NewOrm()
	// user1 := User{Id: 6}
	// if o.Read(&user1) == nil {
	// 	user.Name = "MyName"
	// 	if num, err := o.Update(&user1); err == nil {
	// 		fmt.Println(num)
	// 	}
	// }

	// // read one
	// u := User{Id: user.Id}
	// err = o.Read(&u)
	// fmt.Printf("ERR: %v\n", err)

	// delete
	// u := User{Id: user.Id}
	// num, err = o.Delete(&u)
	// fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}
