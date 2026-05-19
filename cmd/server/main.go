package main

import (
	"fmt"
	"net/http"

	"concurrent-task-queue-server/internal/api"
	"concurrent-task-queue-server/internal/queue"
	"concurrent-task-queue-server/internal/task"
	"concurrent-task-queue-server/internal/worker"
)

func main() {
	fmt.Println("Server starting...")

	taskQueue := queue.NewTaskQueue(10)
	taskStore := task.NewStore()

	worker.StartWorker(1, taskQueue.Tasks, taskStore)
	worker.StartWorker(2, taskQueue.Tasks, taskStore)
	worker.StartWorker(3, taskQueue.Tasks, taskStore)

	taskHandler := &api.TaskHandler{
		TaskQueue: taskQueue.Tasks,
		Store:     taskStore,
	}

	http.HandleFunc("/tasks", taskHandler.CreateTask)
	http.HandleFunc("/tasks/", taskHandler.GetTask)

	fmt.Println("Listening on http://localhost:8080...")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}

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
		taskQueue.Tasks <- t
		fmt.Println("Task added to queue:")
		fmt.Println(t)
	}

	close(taskQueue.Tasks)
	fmt.Println("Server shutting down...")

}
