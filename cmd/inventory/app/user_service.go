package app

import (
	"github.com/miun173/inventory-dispusibda/cmd/inventory/models"
	"github.com/miun173/inventory-dispusibda/cmd/inventory/repository"
)

// UserService user domain service
type UserService interface {
	Save(user *models.User) error
	Get(id int) (models.User, error)
	GetAll() ([]models.User, error)
	UserOfName(username string) (models.User, error)
}

type userService struct {
	repo repository.UserRepo
}

// NewUserService get user service object
func NewUserService(repo repository.UserRepo) UserService {
	return &userService{repo}
}

func (s *userService) Save(user *models.User) error {
	err := s.repo.Save(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) Get(id int) (models.User, error) {
	var user models.User
	user, err := s.repo.Get(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *userService) GetAll() ([]models.User, error) {
	users, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *userService) UserOfName(username string) (models.User, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return user, err
	}

	return user, nil
}
