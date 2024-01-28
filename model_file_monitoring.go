package main

import "time"

type FileMonitoring struct {
	Id           int
	RunTimestamp time.Time
	FileName     string
	MagicString  int
	NewFileAdded bool
	FileDeleted  bool
}
