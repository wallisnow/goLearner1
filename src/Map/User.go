package Map

import (
	"fmt"
	"sort"
)

type User map[string]interface{}

func NewUser() User {
	return make(map[string]interface{})
}

func Modify(u User) {
	u["1"] = "Lili"
}

func (this User) WithValue(key string, value interface{}) User {
	this[key] = value
	return this
}

//这里其实重写了string 方法, 当外部调用Printf则会调用此方法
//有点像java的toString
func (this User) String() string {
	str := ""
	for index, k := range this.Fields() {
		str += fmt.Sprintf("%d, %v -> %v\n", index, k, this[k])
	}
	return str
}

func (this User) Fields() []string {
	var keys []string
	for k, _ := range this {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
