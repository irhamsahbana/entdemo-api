package repository

import (
	"context"
	"entdemo-api/ent"
)

type UserRepository interface {
	FindAll()([]*ent.User, error)
	FindByID(ID int)(*ent.User, error)
	UserCreate(newUser ent.User) (*ent.User, error)
}

type userRepository struct{
	ctx context.Context
	client *ent.Client
}

func UserNewRepository(client *ent.Client, ctx context.Context) *userRepository {
	return &userRepository{ctx: ctx, client: client} 
}

func (r *userRepository)FindAll()([]*ent.User, error)  {

	users, err := r.client.User.Query().All(r.ctx)
	
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) FindByID(ID int)(*ent.User, error) {
	
	user, err := r.client.User.Get(r.ctx, ID)
	
	if err != nil {
		return nil, err
	}
	
	return user, nil
}


func (r *userRepository) UserCreate(newUser ent.User) (*ent.User, error) {
	newCreatedUser, err := r.client.User.Create().
			SetAge(newUser.Age).
			SetName(newUser.Name).
			Save(r.ctx)

	if err != nil {
		return nil, err
	}

	return newCreatedUser, nil
} 