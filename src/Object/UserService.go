package Object

import "fmt"

type UserService struct {
}

func NewUserService() *UserService {
	return new(UserService)
}

func (this *UserService) Save(data interface{}) IService {
	//由于接口本身是不知道类型, 那么外部传入参数时, 那么我们就需要判断他的类型
	if user, ok := data.(*User); ok {
		fmt.Println("save user", user.Name)
	} else {
		fmt.Println("ERROR: save an not User type object")
	}
	return this
}

func (this *UserService) List() IService {
	fmt.Println("list users")
	return this
}
