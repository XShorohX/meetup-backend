package handlers

import (
	"go-backend/models"
	"go-backend/serializers"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ReportHandler struct {
	DB *gorm.DB
}

func (h *ReportHandler) GetReports(c *gin.Context) {
	var reports []models.Report
	h.DB.Find(&reports)

	serializedReports := make([]serializers.ReportSerializer, len(reports))
	for i, report := range reports {
		serializedReports[i] = serializers.NewReportSerializer(report)
	}

	c.JSON(http.StatusOK, serializedReports)
}

func (h *ReportHandler) CreateReport(c *gin.Context) {
	var report models.Report

	if err := c.ShouldBindJSON(&report); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.DB.Create(&report)

	serializedReport := serializers.NewReportSerializer(report)
	c.JSON(http.StatusCreated, serializedReport)
}

func (h *ReportHandler) UpdateReport(c *gin.Context) {
	var report models.Report
	reportID := c.Param("id")

	if err := h.DB.First(&report, reportID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Report not found"})
		return
	}

	if err := c.ShouldBindJSON(&report); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.DB.Save(&report)

	serializedReport := serializers.NewReportSerializer(report)
	c.JSON(http.StatusOK, serializedReport)
}

func (h *ReportHandler) DeleteReport(c *gin.Context) {
	var report models.Report
	reportID := c.Param("id")

	if err := h.DB.First(&report, reportID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Report not found"})
		return
	}

	h.DB.Delete(&report)

	c.JSON(http.StatusNoContent, gin.H{})
}


func (h *ReportHandler) GetReport(c *gin.Context) {
	reportID := c.Param("id")
	var report models.Report
	if err := h.DB.First(&report, reportID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Report not found"})
		return
	}
	c.JSON(http.StatusOK, report)
}