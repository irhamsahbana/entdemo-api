package service

import (
	"context"
	"entdemo-api/ent"
	"entdemo-api/model"
	"entdemo-api/repository"
)

type UserService interface {
	FindAll(ctx context.Context) ([]*ent.User, error)
	FindByID(ctx context.Context, ID int) (*ent.User, error)
	UserCreate(ctx context.Context, userRequest model.UserRequest) (*ent.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func UserNewService(userRepository repository.UserRepository) *userService {
	return &userService{userRepository}
}

func (s *userService) FindAll(ctx context.Context) ([]*ent.User, error) {
	users, err := s.userRepository.FindAll(ctx)
	return users, err
}

func (s *userService) FindByID(ctx context.Context, ID int) (*ent.User, error) {
	user, err := s.userRepository.FindByID(ctx, ID)
	return user, err
}

func (s *userService) UserCreate(ctx context.Context, userRequest model.UserRequest) (*ent.User, error) {
	user := ent.User{
		Name: userRequest.Name,
		Age:  userRequest.Age,
	}
	newUser, err := s.userRepository.UserCreate(ctx, user)
	return newUser, err
}
