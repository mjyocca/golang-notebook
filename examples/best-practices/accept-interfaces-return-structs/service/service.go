package service

import "context"

type UserStorer interface {
	Insert(item interface{}) error
	Get(id string, val interface{}) error
}

type User struct {
	ID    string
	Email string
}

type UserService struct {
	store UserStorer
}

func NewUserService(s UserStorer) *UserService {
	return &UserService{
		store: s,
	}
}

func (u *UserService) CreateUser(ctx context.Context, user *User) error {
	err := u.store.Insert(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) RetrieveUser(ctx context.Context, id string) (interface{}, error) {
	var user User
	err := u.store.Get(id, &user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
