package service

type BaseService interface {
	List()
	Page()
	oneById()
	Create()
	Update()
	Remove()
}
