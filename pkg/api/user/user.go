package user

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"github.com/hi019/fiber-boilerplate/ent"
	"github.com/hi019/fiber-boilerplate/ent/user"
)

// Create creates a user. It returns the created user instance and an error
func (u *User) Create(ctx context.Context, email string, password string) (*ent.User, error) {
	u.Log.Debug().Msg("Creating user")

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return &ent.User{}, nil
	}

	// Create user
	return u.DB.User.Create().SetEmail(email).SetPassword(string(hash)).Save(ctx)
}

// Login authenticates a user with an email and password. It returns the user instance (if found) and an error
func (u *User) Login(ctx context.Context, email string, password string) (e *ent.User, err error) {
	e, err = u.DB.User.Query().Where(user.EmailEQ(email)).Only(ctx)

	return
}
