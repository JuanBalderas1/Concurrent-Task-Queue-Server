package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"concurrent-task-queue-server/internal/task"
)

type TaskHandler struct {
	TaskQueue chan task.Task
	Store     *task.Store
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newTask task.Task

	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	newTask.Status = "queued"

	h.Store.Save(newTask)
	h.TaskQueue <- newTask

	fmt.Printf("Received task %d from API\n", newTask.ID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTask)
}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idString := strings.TrimPrefix(r.URL.Path, "/tasks/")

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	foundTask, exists := h.Store.Get(id)
	if !exists {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(foundTask)
}
