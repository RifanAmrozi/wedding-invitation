package repository

import (
	"wedding-invitation/internal/entity"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	CreateUser(user *entity.User) error
	GetUserByUsername(username string) (*entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *entity.User) error {
	user.ID = uuid.New()
	return r.db.Table("wedding_users").Create(user).Error
}

func (r *userRepository) GetUserByUsername(username string) (*entity.User, error) {
	var user entity.User
	err := r.db.Table("wedding_users").Where("username = ?", username).First(&user).Error
	return &user, err
}
