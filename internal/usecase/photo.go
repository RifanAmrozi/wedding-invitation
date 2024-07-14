package usecase

import (
	"wedding-invitation/internal/entity"
	"wedding-invitation/internal/repository"

	"github.com/google/uuid"
)

type PhotoUsecase interface {
	PostPhoto(userID, photoURL string) error
	GetPhotos(userID string) ([]entity.Photo, error)
}

type photoUsecase struct {
	photoRepo repository.PhotoRepository
}

func NewPhotoUsecase(photoRepo repository.PhotoRepository) PhotoUsecase {
	return &photoUsecase{photoRepo: photoRepo}
}

func (p *photoUsecase) PostPhoto(userID, photoURL string) error {
	photo := &entity.Photo{
		UserID:   uuid.MustParse(userID),
		PhotoURL: photoURL,
	}

	return p.photoRepo.CreatePhoto(photo)
}

func (p *photoUsecase) GetPhotos(userID string) ([]entity.Photo, error) {
	return p.photoRepo.GetPhotosByUserID(userID)
}
