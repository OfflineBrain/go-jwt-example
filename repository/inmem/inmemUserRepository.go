package inmem

import (
	"errors"

	"github.com/offlinebrain/go-jwt-example/model"
	"github.com/offlinebrain/go-jwt-example/repository"
)

var Instance *UserRepository

type UserRepository struct {
	users map[string]model.User
}

func (r *UserRepository) Get(username string) (model.User, error) {
	user, ok := r.users[username]
	if !ok {
		return model.User{}, errors.New("user not found")
	}

	return user, nil
}

func (r *UserRepository) Save(user *model.User) error {
	r.users[user.Username] = *user

	return nil
}

func NewUserRepository() repository.UserRepository {
	Instance = &UserRepository{
		users: make(map[string]model.User),
	}

	return Instance
}
