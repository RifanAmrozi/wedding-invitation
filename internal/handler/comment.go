package handler

import (
	"net/http"
	"wedding-invitation/internal/usecase"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	commentUsecase usecase.CommentUsecase
}

func NewCommentHandler(commentUsecase usecase.CommentUsecase) *CommentHandler {
	return &CommentHandler{commentUsecase: commentUsecase}
}

func (h *CommentHandler) PostComment(c *gin.Context) {
	var input struct {
		PhotoID   string `json:"photo_id"`
		GuestName string `json:"guest_name"`
		Comment   string `json:"comment"`
		Presence  bool   `json:"presence"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.commentUsecase.PostComment(input.PhotoID, input.GuestName, input.Comment, input.Presence); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment posted successfully"})
}

func (h *CommentHandler) GetComments(c *gin.Context) {
	photoID := c.Param("photo_id")

	comments, err := h.commentUsecase.GetComments(photoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}
