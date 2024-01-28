package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Payload struct {
	FileName string `json:"file_name"`
	FetchAll bool   `json:"fetch_all"`
}

func startServer() {
	var err error
	// Set up Gin router
	r := gin.Default()
	r.GET("/api/run-background-task", runBackgroundTaskHandler)         //api to run background task manually
	r.POST("/api/file-monitoring/details", fetchFileMonitoringDetails)  //api to get file details
	r.POST("/api/file-monitoring/details/fetch-all", fetchAllFileNames) //api to fetch all file names

	// Start the server
	//port := os.Getenv("serverPort")
	log.Println("Server started on: ", serverPort)

	err = r.Run(serverPort)
	if err != nil {
		log.Fatal("Error starting the server:", err)
	}
}
