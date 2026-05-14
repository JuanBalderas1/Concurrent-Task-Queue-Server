package main

import (
	"fmt"

	"concurrent-task-queue-server/internal/queue"
	"concurrent-task-queue-server/internal/task"
)

func main() {
	fmt.Println("Server starting...")

	taskQueue := queue.NewTaskQueue(10)

	newTask := task.Task{
		ID:      1,
		Type:    "email",
		Payload: "Send welcome email",
		Status:  "queued",
	}

	taskQueue.Tasks <- newTask

	fmt.Println("Task added to queue:")
	fmt.Println(newTask)
}
