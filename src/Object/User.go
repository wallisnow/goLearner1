package Object

import (
	"fmt"
)

type User struct {
	Id   int
	Sex  int
	Name string
}

//不适用指针的方式, 也是go中不推荐的方式
func NewUser() User {
	return User{}
}

//go中正确的"构造函数", 返回指针, 方法1
func NewUUser() *User {
	return new(User)
}

//go中正确的"构造函数", 返回指针, 方法2
func NewUUUser() *User {
	return &User{}
}

//写一个通用的构造函数
//
func NewCustomizedUser(f []func(u *User)) *User {
	u := new(User)
	for _, fun := range f {
		fun(u)
	}
	return u
}

func SetId(id int) func(u *User) {
	return func(u *User) {
		u.Id = id
	}
}

//这个方法, 即便传入了对象u, 其id还是不会变的
func WrongChangeId(u User) {
	u.Id = 10
}

func OkChangeId(u User) User {
	fmt.Printf("parameter user address %p\n", &u)
	u.Id = 20
	return u
}

func CorrectChangeId(u *User) {
	fmt.Printf("parameter user address %p %p %T %v\n", &u, u, u, *u)
	u.Id = 30
}
