package service

import (
	"context"
	"database/sql"

	"github.com/bernardn38/ebuy/authentication-service/models"
	"github.com/bernardn38/ebuy/authentication-service/sql/users"
	"github.com/bernardn38/ebuy/authentication-service/token"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userQuries   *users.Queries
	usersDb      *sql.DB
	tokenManager *token.Manager
}

func NewAuthService(q *users.Queries, db *sql.DB, tm *token.Manager) *AuthService {
	return &AuthService{
		userQuries:   q,
		usersDb:      db,
		tokenManager: tm,
	}
}

func (a *AuthService) RegisterUser(ctx context.Context, user models.RegisterPayload) error {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return err
	}
	_, err = a.userQuries.CreateUser(ctx, users.CreateUserParams{
		Username: user.Username,
		Email:    user.Email,
		Password: string(encryptedPassword),
	})
	if err != nil {
		return err
	}

	//to do send queue for proccessing registration in other services
	return nil
}

func (a *AuthService) LoginUser(ctx context.Context, p models.LoginPayload) (*users.User, error) {
	user, err := a.userQuries.GetUserByUsername(ctx, p.Username)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(p.Password))
	if err != nil {
		return nil, err
	}
	return &user, nil
}
