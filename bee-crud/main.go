// package main

// import (
// 	"fmt"

// 	"github.com/beego/beego/v2/client/orm"
// 	_ "github.com/go-sql-driver/mysql"
// )

// func init() {
// 	orm.RegisterDriver("mysql", orm.DRMySQL)

// 	orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8")
// }

// func main() {

// 	// Using default, you can use other database
// 	o := orm.NewOrm()

// 	profile := new(Profile)
// 	profile.Age = 30

// 	user := new(User)
// 	user.Profile = profile
// 	user.Name = "slene"

// 	fmt.Println(o.Insert(profile))
// 	fmt.Println(o.Insert(user))
// }

package main

import (
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql" // import your required driver
)

// Model Struct
type User struct {
	Id   int
	Name string `orm:"size(100)"`
}

func init() {
	// register model
	orm.RegisterModel(new(User))

	// set default database
	orm.RegisterDataBase("default", "mysql", "root:my-secret-pw@/my_db?charset=utf8")
}

func main() {
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
