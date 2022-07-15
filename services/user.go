package services

import (
	"backbootcamp/models"
)

type IUserRepository interface {
	Login(email, password string) (*models.User, error)
	AddUser(user *models.User) (*models.User, error)
	GetUserById(idUser int) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	UpdateUser(idUser int, user *models.User) (*models.User, error)
	DeleteUser(idUser int) error
}

type UserService struct {
	repository IUserRepository
}

func NewUserService(repository IUserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) Login(email, password string) (*models.User, error) {
	user, err := s.repository.Login(email, password)
	return user, err
}

func (s *UserService) InsertUser(user *models.User) (*models.User, error) {
	if err := user.ValidUser(); err != nil {
		return nil, err
	}
	user, err := s.repository.AddUser(user)
	return user, err
}

func (s *UserService) GetUser(idUser int) (*models.User, error) {
	user, err := s.repository.GetUserById(idUser)
	return user, err
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	users, err := s.repository.GetAllUsers()
	return users, err
}

func (s *UserService) UpdateUser(idUser int, user *models.User) (*models.User, error) {
	if err := user.ValidUser(); err != nil {
		return nil, err
	}
	found, err := s.GetUser(idUser)
	if err != nil {
		return found, err
	}
	userUpdate, err := s.repository.UpdateUser(idUser, user)
	return userUpdate, err
}

func (s *UserService) DeleteUser(idUser int) error {
	_, err := s.GetUser(idUser)
	if err != nil {
		return err
	}
	err = s.repository.DeleteUser(idUser)
	return err
}
