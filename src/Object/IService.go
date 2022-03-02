package Object

type IService interface {
	Save(s interface{}) IService
	List() IService
}
