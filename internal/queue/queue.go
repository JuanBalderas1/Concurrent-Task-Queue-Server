package queue

import "concurrent-task-queue-server/internal/task"

type TaskQueue struct {
	Tasks chan task.Task
}

func NewTaskQueue(bufferSize int) *TaskQueue {
	return &TaskQueue{
		Tasks: make(chan task.Task, bufferSize),
	}
}
