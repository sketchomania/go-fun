package main

import (
	"github.com/beego/beego/v2/client/orm"
)

type User struct {
	Id      int
	Name    string
	Profile *Profile `orm:"rel(one)"` // OneToOne relation
}

type Profile struct {
	Id   int
	Age  int16
	User *User `orm:"reverse(one)"` // Reverse relationship (optional)
}

type Product struct {
	Id    int
	Name  string
	Price int
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(User), new(Profile), new(Product))
}
