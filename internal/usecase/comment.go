package usecase

import (
	"wedding-invitation/internal/entity"
	"wedding-invitation/internal/repository"

	"github.com/google/uuid"
)

type CommentUsecase interface {
	PostComment(photoID, guestName, comment string) error
	GetComments(photoID string) ([]entity.Comment, error)
}

type commentUsecase struct {
	commentRepo repository.CommentRepository
}

func NewCommentUsecase(commentRepo repository.CommentRepository) CommentUsecase {
	return &commentUsecase{commentRepo: commentRepo}
}

func (c *commentUsecase) PostComment(photoID, guestName, comment string) error {
	commentEntity := &entity.Comment{
		PhotoID:   uuid.MustParse(photoID),
		GuestName: guestName,
		Comment:   comment,
	}

	return c.commentRepo.CreateComment(commentEntity)
}

func (c *commentUsecase) GetComments(photoID string) ([]entity.Comment, error) {
	return c.commentRepo.GetCommentsByPhotoID(photoID)
}
