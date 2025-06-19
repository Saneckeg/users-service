// UserService (бизнес-логика)
package user

type UserService struct {
	repo UserRepository
}

func NewService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user User) (User, error) {
	return s.repo.CreateUser(user)
}

func (s *UserService) GetAllUsers() ([]User, error) {
	return s.repo.GetAllUsers()
}

func (s *UserService) UpdateUserByID(id uint, user interface{}) (User, error) {
	return s.repo.UpdateUserByID(id, user)
}

func (s *UserService) DeleteUserByID(id uint) (User, error) {
	return s.repo.DeleteUserByID(id)
}
