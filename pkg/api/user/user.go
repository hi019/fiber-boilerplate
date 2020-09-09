package user

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"github.com/hi019/fiber-boilerplate/ent"
	"github.com/hi019/fiber-boilerplate/ent/user"
)

func (u *User) Create(email string, password string, ctx context.Context) (*ent.User, error) {
	u.Log.Debug().Msg("Creating user")

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return &ent.User{}, nil
	}

	// Create user
	return u.DB.User.Create().SetEmail(email).SetPassword(string(hash)).Save(ctx)
}

func (u *User) Login(email string, password string, ctx context.Context) (e *ent.User, err error) {
	e, err = u.DB.User.Query().Where(user.EmailEQ(email)).Only(ctx)

	return
}