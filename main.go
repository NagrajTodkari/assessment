package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

const (
	serverPort       = ":8000"
	monitorDirectory = "/home/nagraj/SUITS/ForStudy/assessment"
)

func main() {
	InitConnectionsForJobs()
	RunMigration()
	setUpCron()
	startServer()
}

func runBackgroundTask(db *gorm.DB) {
	log.Println("Running background task...")

	files, err := ioutil.ReadDir(monitorDirectory)
	if err != nil {
		log.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		filePath := filepath.Join(monitorDirectory, file.Name())

		fileAdded := checkFileAdded(db, file.Name())
		if fileAdded {
			log.Printf("New file added: %s\n", file.Name())
		}

		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			log.Printf("Error reading file %s: %s\n", file.Name(), err)
			continue
		}

		magicStringCount := strings.Count(string(content), os.Getenv("magicString"))

		err = db.Create(&FileMonitoring{
			RunTimestamp: time.Now(),
			FileName:     file.Name(),
			MagicString:  magicStringCount,
			NewFileAdded: fileAdded,
			FileDeleted:  false,
		}).Error

		if err != nil {
			log.Println("Error inserting data into database:", err)
		}
	}

	checkForDeletedFiles(db, files)
}

func checkFileAdded(db *gorm.DB, fileName string) bool {
	var fileAdded bool

	err := db.Model(&FileMonitoring{}).
		Select("1").
		Where("file_name = ? AND file_deleted = FALSE", fileName).
		Scan(&fileAdded).Error
	if err != nil {
		log.Println("Error checking if file was added:", err)
		return false
	}
	return !fileAdded
}

func checkForDeletedFiles(db *gorm.DB, currentFiles []os.FileInfo) {
	var existingFiles []string

	err := db.Model(&FileMonitoring{}).
		Select("file_name").
		Where("file_deleted = FALSE").
		Scan(&existingFiles).Error

	if err != nil {
		log.Println("Error checking for deleted files:", err)
		return
	}

	for _, existingFile := range existingFiles {
		found := false
		for _, currentFile := range currentFiles {
			if existingFile == currentFile.Name() {
				found = true
				break
			}
		}

		if !found {
			log.Printf("File deleted: %s\n", existingFile)

			err := db.Model(&FileMonitoring{}).
				Update("file_deleted", true).
				Where("file_name = ?", existingFile).Error
			if err != nil {
				log.Println("Error updating database for deleted file:", err)
			}
		}
	}
}

func runBackgroundTaskHandler(c *gin.Context) {
	db := DB
	runBackgroundTask(db)
	c.String(http.StatusOK, "Background task executed successfully")
}
