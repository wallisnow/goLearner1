package main

import (
	"base/src/Map"
	"base/src/Object"
	"base/src/String"
	"base/src/Utils"
	"fmt"
	"reflect"
	"sort"
)

func main() {

	/////////////////lesson 1////////////////////
	fmt.Println("/////////////////lesson 1////////////////////")
	s := String.From("mystring_type")
	fmt.Println(s)

	si := String.FromInt(123)
	fmt.Println(si)

	s.Len()
	fmt.Println(s.Len())

	//遍历字符串
	for i := 0; i < s.Len(); i++ {
		fmt.Println(i, " ", fmt.Sprintf("%c", s[i]))
	}

	s.ForEach(func(chars string) {
		fmt.Println(chars)
	})

	/////////////////lesson 2////////////////////
	fmt.Println("/////////////////lesson 2////////////////////")
	user := Object.NewUser()
	fmt.Println(user)
	user.Id = 1
	//这里使用%+v, 可以打印出来属性名
	fmt.Printf("%+v\n", user)

	//使用错误的方式改变id值
	//这里和java是不同的, 因为go是形式参数, 传入值只是一个值拷贝
	Object.WrongChangeId(user)
	fmt.Printf("%+v\n", user)

	//使用返回值的方式, 则可以, 因为最终你将使用新的user2, 也就是函数内部返回的对象而不是原先的对象
	//original address 0xc00005e020
	//parameter user address 0xc00005e0c0
	//current address 0xc00005e0a0
	//可以看出, 此时user 在传参时, 参数本身是一个新的地址
	fmt.Printf("original address %p\n", &user)
	user2 := Object.OkChangeId(user)
	fmt.Printf("current address %p\n", &user2)
	fmt.Printf("%+v\n", user2)
	fmt.Printf("%+v\n", user)

	//上面这种方式也是不推荐的, 正确的方式是使用指针
	Object.CorrectChangeId(&user)
	//current address after change ID: 0xc00005e020
	fmt.Printf("current address after change ID: %p\n", &user)
	fmt.Printf("%+v\n", user)

	//其实上面的情况, 都是基于构造函数返回的是对象的值, 并不是其地址, 那么我们尝试使用构造函数是指针的方法
	//对比下面代码
	var userx Object.User
	//{Id:0 Sex:0 Name:}
	fmt.Printf("%+v\n", userx)
	var usery *Object.User
	//<nil>
	fmt.Printf("%+v\n", usery)
	userx = Object.NewUser()
	//{Id:0 Sex:0 Name:}
	fmt.Printf("%+v\n", userx)
	usery = Object.NewUUser()
	//&{Id:0 Sex:0 Name:}
	fmt.Printf("%+v\n", usery)
	//上面代码告诉我们, 这里的以指针(*Object.User)的形式声明, 那么必须要进行new或者&, 不然只是一个空指针
	//current usery address 0xc00000e038
	//parameter user address 0xc00000e040
	//运行下面代码, 可以看出, go其实在你传参是, 还是一个值拷贝, 也就是传入的指针的指针还是不一样, 但是其值是指针地址
	fmt.Printf("current usery address %p\n", usery)
	Object.CorrectChangeId(usery)
	fmt.Printf("%p %+v\n", usery, usery)

	/////////////////lesson 3////////////////////
	fmt.Println("/////////////////lesson 3////////////////////")

	customizedUser := Object.NewCustomizedUser(Object.SetId(11), Object.SetName("lucy"))
	fmt.Printf("%p %+v\n", customizedUser, customizedUser)

	useri := Object.NewUUser()
	Object.SetName("Jack").Apply(useri)
	Object.SetId(333).Apply(useri)
	fmt.Printf("%p %+v\n", useri, useri)

	newCustomizedUser123 := Object.NewCustomizedUser(Object.SetId(123), Object.SetName("onetwothree"))
	fmt.Printf("%p %+v\n", newCustomizedUser123, newCustomizedUser123)

	aCustomizedUser555 := Object.NewACustomizedUser(Object.SetId(555), Object.SetName("five3"))
	fmt.Printf("%p %+v\n", aCustomizedUser555, aCustomizedUser555)

	aCustomizedUser555.ApplyFunc(func(u *Object.User) *Object.User {
		u.Id = 666
		u.Name = "sixsixsix"
		u.Sex = 1
		return u
	})

	fmt.Printf("%p %+v\n", aCustomizedUser555, aCustomizedUser555)

	/////////////////lesson 4////////////////////
	fmt.Println("/////////////////lesson 4////////////////////")
	newUser := Map.NewUser()
	newUser["1"] = "Jack"
	newUser["2"] = "Bob"
	fmt.Printf("%+v\n", newUser)

	//go中都是值传递, 但是这里是不需要传指针的, 因为map内部实际已经获取了指针
	Map.Modify(newUser)
	fmt.Printf("%+v\n", newUser)

	newUser.
		WithValue("3", "Eric").
		WithValue("4", "LiuJie")

	fmt.Printf("%+v\n", newUser)

	//_ = newUser.String()

	newUser2 := Map.NewUser()
	newUser2.
		WithValue("1", "Mhamd").
		WithValue("2", "data")

	var users []Map.User
	users = append(users, newUser, newUser2)

	//这里第二个函数类似predictor, 来定义排序的规则
	sort.Slice(users, func(i, j int) bool {
		name := users[i]["1"].(string) //断言
		return name == "Mhamd"
	})

	fmt.Printf("%+v\n", users)

	/////////////////lesson 5////////////////////
	fmt.Println("/////////////////lesson 5////////////////////")

	var service Object.IService
	service = Object.NewUserService()
	userj := Object.NewACustomizedUser(Object.SetId(123), Object.SetName("ZZZ"))
	service.Save(userj)

	service.Save("abc")

	/////////////////lesson 6////////////////////
	fmt.Println("/////////////////lesson 6////////////////////")
	t := reflect.TypeOf(userj)
	fmt.Println(t.String())
	v := reflect.ValueOf(userj)
	//因为传入的是指针, 所以需要先获取指针指向的值
	elem := v.Elem()
	fmt.Println(elem)

	values := []interface{}{12, "test"}

	valmap := make(map[string]interface{})
	valmap["Id"] = 88
	valmap["Sex"] = 1
	valmap["Name"] = "testttt"

	for i := 0; i < elem.NumField(); i++ {
		//使用interface 获取每一个值, 不需要判断
		fmt.Println(elem.Field(i).Interface())

		//如果想通过反射将字段赋值, 可以通过set方法, 但必须确保类型一致
		if elem.Field(i).Kind() == reflect.Int {
			of := reflect.ValueOf(values[0])
			elem.Field(i).Set(of)
		}
	}

	//{10 10 ZZZ} &{10 10 ZZZ}
	fmt.Println(elem, userj)

	Utils.Map2Struct(valmap, userj)
	fmt.Println(elem, userj)

}
