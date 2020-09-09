package user

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/hi019/fiber-boilerplate/ent"
)

type User struct {
	DB *ent.Client
	Log *zerolog.Logger
}

type Service interface {
	Create(email string, password string, ctx context.Context) (user *ent.User, err error)
	Login(email string, password string, ctx context.Context) (e *ent.User, err error)
}

func Initialize(db *ent.Client, log *zerolog.Logger) *User {
	return &User{DB: db, Log: log}
}