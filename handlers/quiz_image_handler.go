package handlers

import (
	"go-backend/models"
	"go-backend/serializers"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type QuizImageHandler struct {
	DB *gorm.DB
}

func (h *QuizImageHandler) GetQuizImages(c *gin.Context) {
	var quizImages []models.QuizImage
	h.DB.Find(&quizImages)

	serializedQuizImages := make([]serializers.QuizImageSerializer, len(quizImages))
	for i, quizImage := range quizImages {
		serializedQuizImages[i] = serializers.NewQuizImageSerializer(quizImage)
	}

	c.JSON(http.StatusOK, serializedQuizImages)
}

func (h *QuizImageHandler) CreateQuizImage(c *gin.Context) {
	var quizImage models.QuizImage

	if err := c.ShouldBindJSON(&quizImage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.DB.Create(&quizImage)

	serializedQuizImage := serializers.NewQuizImageSerializer(quizImage)
	c.JSON(http.StatusCreated, serializedQuizImage)
}

func (h *QuizImageHandler) UpdateQuizImage(c *gin.Context) {
	var quizImage models.QuizImage
	quizImageID := c.Param("id")

	if err := h.DB.First(&quizImage, quizImageID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "QuizImage not found"})
		return
	}

	if err := c.ShouldBindJSON(&quizImage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.DB.Save(&quizImage)

	serializedQuizImage := serializers.NewQuizImageSerializer(quizImage)
	c.JSON(http.StatusOK, serializedQuizImage)
}

func (h *QuizImageHandler) DeleteQuizImage(c *gin.Context) {
	var quizImage models.QuizImage
	quizImageID := c.Param("id")

	if err := h.DB.First(&quizImage, quizImageID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "QuizImage not found"})
		return
	}

	h.DB.Delete(&quizImage)

	c.JSON(http.StatusNoContent, gin.H{})
}
