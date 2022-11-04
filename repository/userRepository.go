package repository

import "github.com/offlinebrain/go-jwt-example/model"

type UserRepository interface {
	Get(username string) (model.User, error)
	Save(user *model.User) error
}
