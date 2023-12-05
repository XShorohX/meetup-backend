package handlers

import (
	"go-backend/models"
	"go-backend/serializers"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SpeakerHandler struct {
	DB *gorm.DB
}

func (h *SpeakerHandler) GetSpeakers(c *gin.Context) {
	var speakers []models.Speaker
	h.DB.Find(&speakers)

	serializedSpeakers := make([]serializers.SpeakerSerializer, len(speakers))
	for i, speaker := range speakers {
		serializedSpeakers[i] = serializers.NewSpeakerSerializer(speaker)
	}

	c.JSON(http.StatusOK, serializedSpeakers)
}

func (h *SpeakerHandler) CreateSpeaker(c *gin.Context) {
	var speaker models.Speaker

	if err := c.ShouldBindJSON(&speaker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.DB.Create(&speaker)

	serializedSpeaker := serializers.NewSpeakerSerializer(speaker)
	c.JSON(http.StatusCreated, serializedSpeaker)
}

func (h *SpeakerHandler) UpdateSpeaker(c *gin.Context) {
	var speaker models.Speaker
	speakerID := c.Param("id")

	if err := h.DB.First(&speaker, speakerID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Speaker not found"})
		return
	}

	if err := c.ShouldBindJSON(&speaker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.DB.Save(&speaker)

	serializedSpeaker := serializers.NewSpeakerSerializer(speaker)
	c.JSON(http.StatusOK, serializedSpeaker)
}

func (h *SpeakerHandler) DeleteSpeaker(c *gin.Context) {
	var speaker models.Speaker
	speakerID := c.Param("id")

	if err := h.DB.First(&speaker, speakerID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Speaker not found"})
		return
	}

	h.DB.Delete(&speaker)

	c.JSON(http.StatusNoContent, gin.H{})
}

func (h *SpeakerHandler) GetSpeaker(c *gin.Context) {
	speakerID := c.Param("id")
	var speaker models.Speaker
	if err := h.DB.First(&speaker, speakerID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Speaker not found"})
		return
	}
	c.JSON(http.StatusOK, speaker)
}
