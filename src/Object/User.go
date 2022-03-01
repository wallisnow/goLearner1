package Object

import (
	"fmt"
)

type User struct {
	Id   int
	Sex  int
	Name string
}

type UserAttrFunc func(u *User)

type UserAttrFuncs []UserAttrFunc

func (this *User) ApplyFunc(f func(u *User) *User) *User {
	return f(this)
}

func (this UserAttrFunc) Apply(u *User) {
	//这里的this, 就是UserAttrFunc, 也就是
	//func(u *User) {}(u)
	this(u)
}

//将上面的多个apply封装成多层函数
func (this UserAttrFuncs) Applys(u *User) {
	for _, attrFunc := range this {
		attrFunc.Apply(u)
	}
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
//给其自由定义属性
//缺点是不好理解
func NewCustomizedUser(f ...UserAttrFunc) *User {
	u := new(User)
	for _, fun := range f {
		//fun(u)
		fun.Apply(u)
	}
	return u
}

func NewACustomizedUser(fs ...UserAttrFunc) *User {
	u := new(User)
	UserAttrFuncs(fs).Applys(u)
	return u
}

func SetId(id int) UserAttrFunc {
	return func(u *User) {
		u.Id = id
	}
}

func SetName(name string) UserAttrFunc {
	return func(u *User) {
		u.Name = name
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
