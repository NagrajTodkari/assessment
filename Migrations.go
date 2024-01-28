package main

func RunMigration() {
	DB.AutoMigrate(&FileMonitoring{})
}
