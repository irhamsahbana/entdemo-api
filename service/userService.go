package service

import (
	"entdemo-api/ent"
	"entdemo-api/model"
	"entdemo-api/repository"
)

type UserService interface {
	FindAll() ([]*ent.User, error)
	FindByID(ID int) (*ent.User, error)
	UserCreate(userRequest model.UserRequest) (*ent.User, error)
}

type userService struct {
	userRepository repository.UserRepository
	
}

func UserNewService(userRepository repository.UserRepository) *userService {
	return &userService{userRepository} 
}

func (s *userService ) FindAll() ([]*ent.User, error)  {
		users, err := s.userRepository.FindAll()
		return users, err
}

func (s *userService ) FindByID(ID int) (*ent.User, error)  {
	user, err := s.userRepository.FindByID(ID)
	return user, err
}

func (s *userService ) UserCreate(userRequest model.UserRequest) (*ent.User, error){
	user := ent.User{
		Name: userRequest.Name,
		Age: userRequest.Age,
	}
	newUser, err := s.userRepository.UserCreate(user)
	return newUser, err
}