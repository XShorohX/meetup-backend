package handlers

import (
	"go-backend/models"
	"go-backend/serializers"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MeetHandler struct {
	DB *gorm.DB
}

func (h *MeetHandler) GetMeets(c *gin.Context) {
	var meets []models.Meet
	h.DB.Find(&meets)

	serializedMeets := make([]serializers.MeetSerializer, len(meets))
	for i, meet := range meets {
		serializedMeets[i] = serializers.NewMeetSerializer(meet)
	}

	c.JSON(http.StatusOK, serializedMeets)
}

func (h *MeetHandler) CreateMeet(c *gin.Context) {
	var meet models.Meet

	if err := c.ShouldBindJSON(&meet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.DB.Create(&meet).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create meetup"})
		return
	}
	
	serializedMeet := serializers.NewMeetSerializer(meet)
	c.JSON(http.StatusCreated, serializedMeet)
}

func (h *MeetHandler) UpdateMeet(c *gin.Context) {
	var meet models.Meet
	meetID := c.Param("id")

	if err := h.DB.First(&meet, meetID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Meet not found"})
		return
	}

	if err := c.ShouldBindJSON(&meet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.DB.Save(&meet)

	serializedMeet := serializers.NewMeetSerializer(meet)
	c.JSON(http.StatusOK, serializedMeet)
}

func (h *MeetHandler) DeleteMeet(c *gin.Context) {
	var meet models.Meet
	meetID := c.Param("id")

	if err := h.DB.First(&meet, meetID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Meet not found"})
		return
	}

	h.DB.Delete(&meet)

	c.JSON(http.StatusNoContent, gin.H{})
}

func (h *MeetHandler) LastStartedMeetIfExist(c *gin.Context) {
    now := time.Now()
    
    var lastStartedMeet, lastEndedMeet models.Meet
    
    // Находим последнее завершенное мероприятие
    h.DB.Where("date < ?", now).Order("date desc").First(&lastEndedMeet)

    // Находим последнее начавшееся мероприятие
    h.DB.Where("date > ?", now).Order("date asc").First(&lastStartedMeet)
    
    // Проверяем наличие митапов и возвращаем null, если митапов вообще нет
    if lastStartedMeet.ID == 0 && lastEndedMeet.ID == 0 {
        c.JSON(http.StatusOK, nil)
        return
    }
    
    // Выбираем последний митап в зависимости от условий
    var resultMeet models.Meet
    if lastStartedMeet.ID == 0 {
        resultMeet = lastEndedMeet
    } else {
        resultMeet = lastStartedMeet
    }

    serializer := serializers.NewMeetSerializer(resultMeet)

    c.JSON(http.StatusOK, serializer)
}

func (h *MeetHandler) GetMeet(c *gin.Context) {
	meetID := c.Param("id")
	var meet models.Meet
	if err := h.DB.First(&meet, meetID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Meet not found"})
		return
	}
	c.JSON(http.StatusOK, meet)
}
