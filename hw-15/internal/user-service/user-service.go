package userservice

type UserService struct {
	storage storage
}

func NewUserService(s storage) *UserService {
	return &UserService{storage: s}
}

type storage interface {
	CreateUser(u User)
	GetUserById(id int) (User, bool)
}

func (us UserService) CreateUser(u User) {

	us.storage.CreateUser(u)
}

func (us UserService) GetUserById(id int) (User, bool) {
	user, ok := us.storage.GetUserById(id)
	return user, ok
}
