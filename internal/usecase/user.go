package usecase

import (
	"wedding-invitation/internal/entity"
	"wedding-invitation/internal/repository"
	"wedding-invitation/pkg/utils"
)

type UserUsecase interface {
	Register(username, password string) error
	Login(username, password string) (string, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

func (u *userUsecase) Register(username, password string) error {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := &entity.User{
		Username:     username,
		PasswordHash: hashedPassword,
	}

	return u.userRepo.CreateUser(user)
}

func (u *userUsecase) Login(username, password string) (string, error) {
	user, err := u.userRepo.GetUserByUsername(username)
	if err != nil {
		return "", err
	}

	if !utils.CheckPasswordHash(password, user.PasswordHash) {
		return "", err
	}

	token, err := utils.GenerateToken(user.ID.String())
	if err != nil {
		return "", err
	}

	return token, nil
}
