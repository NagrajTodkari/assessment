package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type FileDetails struct {
	RunTimeStamp time.Time `json:"run_timestamp"`
	FileName     string    `json:"file_name"`
	MagicString  int       `json:"magic_string"`
}

func fetchFileMonitoringDetails(c *gin.Context) {
	var (
		fileMonitoringDetails []FileDetails
		param                 Payload

		err error
	)

	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := DB.Model(&FileMonitoring{}).
		Select("run_timestamp, file_name, magic_string").
		Where("file_deleted = FALSE")

	if len(param.FileName) > 0 {
		db = db.Where("file_name = ?", param.FileName)
	}

	err = db.Scan(&fileMonitoringDetails).Error

	if err != nil {
		log.Println("Error fetching file monitoring details:", err)
		return
	}

	log.Println("File monitoring details fetched successfully")
	c.JSON(http.StatusOK, fileMonitoringDetails)

}

func fetchAllFileNames(c *gin.Context) {
	var (
		fileNames []string
		param     Payload
	)

	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := DB.Model(&FileMonitoring{}).
		Select("file_name").
		Where("file_deleted = FALSE").
		Scan(&fileNames).Error

	if err != nil {
		log.Println("Error fetching file names:", err)
		return
	}

	log.Println("File names fetched successfully")
	c.JSON(http.StatusOK, fileNames)

}
