package repository

import (
	"wedding-invitation/internal/entity"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type CommentRepository interface {
	CreateComment(comment *entity.Comment) error
	GetCommentsByPhotoID() ([]entity.Comment, error)
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{db: db}
}

func (r *commentRepository) CreateComment(comment *entity.Comment) error {
	comment.ID = uuid.New()
	return r.db.Create(comment).Error
}

func (r *commentRepository) GetCommentsByPhotoID() ([]entity.Comment, error) {
	var comments []entity.Comment
	err := r.db.Find(&comments).Error
	return comments, err
}
