package userservice

type UserService struct {
	storage storage
}

type storage interface {
	CreateUser()
}
