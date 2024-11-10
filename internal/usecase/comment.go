package usecase

import (
	"wedding-invitation/internal/entity"
	"wedding-invitation/internal/repository"
)

type CommentUsecase interface {
	PostComment(guestName, comment string, presence bool) error
	GetComments() ([]entity.Comment, error)
}

type commentUsecase struct {
	commentRepo repository.CommentRepository
}

func NewCommentUsecase(commentRepo repository.CommentRepository) CommentUsecase {
	return &commentUsecase{commentRepo: commentRepo}
}

func (c *commentUsecase) PostComment(guestName, comment string, presence bool) error {
	commentEntity := &entity.Comment{
		GuestName: guestName,
		Comment:   comment,
		Presence:  presence,
	}

	return c.commentRepo.CreateComment(commentEntity)
}

func (c *commentUsecase) GetComments() ([]entity.Comment, error) {
	return c.commentRepo.GetCommentsByPhotoID()
}
