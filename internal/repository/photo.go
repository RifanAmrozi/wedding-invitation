package repository

import (
	"wedding-invitation/internal/entity"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type PhotoRepository interface {
	CreatePhoto(photo *entity.Photo) error
	GetPhotosByUserID(userID string) ([]entity.Photo, error)
}

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) PhotoRepository {
	return &photoRepository{db: db}
}

func (r *photoRepository) CreatePhoto(photo *entity.Photo) error {
	photo.ID = uuid.New()
	return r.db.Create(photo).Error
}

func (r *photoRepository) GetPhotosByUserID(userID string) ([]entity.Photo, error) {
	var photos []entity.Photo
	err := r.db.Where("user_id = ?", userID).Find(&photos).Error
	return photos, err
}
