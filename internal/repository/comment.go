package repository

import (
	"wedding-invitation/internal/entity"

	"github.com/jinzhu/gorm"
)

type CommentRepository interface {
	CreateComment(comment *entity.Comment) error
	GetCommentsByPhotoID(photoID string) ([]entity.Comment, error)
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{db: db}
}

func (r *commentRepository) CreateComment(comment *entity.Comment) error {
	return r.db.Create(comment).Error
}

func (r *commentRepository) GetCommentsByPhotoID(photoID string) ([]entity.Comment, error) {
	var comments []entity.Comment
	err := r.db.Where("photo_id = ?", photoID).Find(&comments).Error
	return comments, err
}
