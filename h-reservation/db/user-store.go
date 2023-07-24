package db

import "github.com/playground/h-reservation/types"

type UserStore interface {
	GetUserById(string) (*types.User, error)
}

type MongoDBStore struct{}
