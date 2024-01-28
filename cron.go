package main

import (
	"log"

	"github.com/robfig/cron"
)

func setUpCron() {
	var e error
	c := cron.New()

	// Schedule the runBackgroundTask function to run every 5 minutes
	e = c.AddFunc("* 1 * * * *", backgroundTask)
	if e != nil {
		log.Fatal("Error scheduling cron job:", e)
	}

	// Start the cron scheduler
	c.Start()
}

func backgroundTask() {
	// db, err := sql.Open("postgres", dbConnectionString)
	// if err != nil {
	// 	log.Fatal("Unable to connect to the database:", err)
	// }
	// defer db.Close()
	db := DB
	runBackgroundTask(db)

}
