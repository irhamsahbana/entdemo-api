package repository

import (
	"context"
	"entdemo-api/ent"
)

type UserRepository interface {
	FindAll(ctx context.Context) ([]*ent.User, error)
	FindByID(ctx context.Context, ID int) (*ent.User, error)
	UserCreate(ctx context.Context, newUser ent.User) (*ent.User, error)
}

type userRepository struct {
	client *ent.Client
}

func UserNewRepository(client *ent.Client) *userRepository {
	return &userRepository{client: client}
}

func (r *userRepository) FindAll(ctx context.Context) ([]*ent.User, error) {

	users, err := r.client.User.Query().All(ctx)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) FindByID(ctx context.Context, ID int) (*ent.User, error) {

	user, err := r.client.User.Get(ctx, ID)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) UserCreate(ctx context.Context, newUser ent.User) (*ent.User, error) {
	newCreatedUser, err := r.client.User.Create().
		SetAge(newUser.Age).
		SetName(newUser.Name).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return newCreatedUser, nil
}
