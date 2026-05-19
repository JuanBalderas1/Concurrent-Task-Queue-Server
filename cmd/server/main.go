package main

import (
	"fmt"
	"sync"

	"concurrent-task-queue-server/internal/queue"
	"concurrent-task-queue-server/internal/task"
	"concurrent-task-queue-server/internal/worker"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Server starting...")

	taskQueue := queue.NewTaskQueue(10)

	worker.StartWorker(1, taskQueue.Tasks, &wg)
	worker.StartWorker(2, taskQueue.Tasks, &wg)
	worker.StartWorker(3, taskQueue.Tasks, &wg)

	tasks := []task.Task{
		{
			ID:         1,
			Type:       "email",
			Payload:    "Send welcome email",
			Status:     "queued",
			MaxRetries: 3,
		},
		{
			ID:         2,
			Type:       "sms",
			Payload:    "Send login code",
			Status:     "queued",
			MaxRetries: 3,
		},
		{
			ID:         3,
			Type:       "report",
			Payload:    "Generate daily report",
			Status:     "queued",
			MaxRetries: 3,
		},
		{
			ID:         4,
			Type:       "backup",
			Payload:    "Run database backup",
			Status:     "queued",
			MaxRetries: 3,
		},

		{
			ID:         5,
			Type:       "email",
			Payload:    "Send password reset email",
			Status:     "queued",
			MaxRetries: 3,
		},
	}

	for _, t := range tasks {
		wg.Add(1)
		taskQueue.Tasks <- t
		fmt.Println("Task added to queue:")
		fmt.Println(t)
	}

	wg.Wait()
	close(taskQueue.Tasks)
	fmt.Println("Server shutting down...")

}
