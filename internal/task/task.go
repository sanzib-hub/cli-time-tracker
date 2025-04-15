package task

import "time"



type Task struct {
	ID	   int
	Title  string
	IsRunning  bool
	StartTime time.Time
	TotalTime time.Duration
	CreatedAt time.Time
}



