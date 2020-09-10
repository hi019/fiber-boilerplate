package user

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/hi019/fiber-boilerplate/ent"
)

// User service dependencies
type User struct {
	DB  *ent.Client
	Log *zerolog.Logger
}

// Service contains User service methods
type Service interface {
	Create(ctx context.Context, email string, password string) (user *ent.User, err error)
	Login(ctx context.Context, email string, password string) (e *ent.User, err error)
}

// Initialize initializes a user service
func Initialize(db *ent.Client, log *zerolog.Logger) *User {
	return &User{DB: db, Log: log}
}
