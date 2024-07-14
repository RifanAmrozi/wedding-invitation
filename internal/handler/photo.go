package handler

import (
	"net/http"
	"wedding-invitation/internal/usecase"

	"github.com/gin-gonic/gin"
)

type PhotoHandler struct {
	photoUsecase usecase.PhotoUsecase
}

func NewPhotoHandler(photoUsecase usecase.PhotoUsecase) *PhotoHandler {
	return &PhotoHandler{photoUsecase: photoUsecase}
}

func (h *PhotoHandler) PostPhoto(c *gin.Context) {
	userID := c.GetString("userID")
	var input struct {
		PhotoURL string `json:"photo_url"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.photoUsecase.PostPhoto(userID, input.PhotoURL); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo posted successfully"})
}

func (h *PhotoHandler) GetPhotos(c *gin.Context) {
	userID := c.GetString("userID")

	photos, err := h.photoUsecase.GetPhotos(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, photos)
}
