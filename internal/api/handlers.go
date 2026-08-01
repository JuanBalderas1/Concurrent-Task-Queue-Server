package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"concurrent-task-queue-server/internal/member"
	"concurrent-task-queue-server/internal/task"
)

type TaskHandler struct {
	TaskQueue   chan task.Task
	TaskStore   *task.Store
	MemberStore *member.Store
}

func (h *TaskHandler) HandleTasks(
	w http.ResponseWriter,
	r *http.Request,
) {
	switch r.Method {
	case http.MethodPost:
		h.CreateTask(w, r)

	case http.MethodGet:
		h.ListTasks(w)

	default:
		http.Error(
			w,
			"Method not allowed",
			http.StatusMethodNotAllowed,
		)
	}
}

func (h *TaskHandler) CreateTask(
	w http.ResponseWriter,
	r *http.Request,
) {
	var newTask task.Task

	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		http.Error(
			w,
			"Invalid JSON request body",
			http.StatusBadRequest,
		)
		return
	}

	newTask.Type = strings.TrimSpace(newTask.Type)
	newTask.Payload = strings.TrimSpace(newTask.Payload)

	if newTask.SenderID <= 0 {
		http.Error(
			w,
			"Valid sender ID is required",
			http.StatusBadRequest,
		)
		return
	}

	if newTask.RecipientID <= 0 {
		http.Error(
			w,
			"Valid recipient ID is required",
			http.StatusBadRequest,
		)
		return
	}

	if newTask.SenderID == newTask.RecipientID {
		http.Error(
			w,
			"Sender and recipient must be different members",
			http.StatusBadRequest,
		)
		return
	}

	if newTask.Type == "" {
		newTask.Type = "message"
	}

	if newTask.Payload == "" {
		http.Error(
			w,
			"Payload is required",
			http.StatusBadRequest,
		)
		return
	}

	_, senderExists := h.MemberStore.Get(
		newTask.SenderID,
	)
	if !senderExists {
		http.Error(
			w,
			"Sender member not found",
			http.StatusBadRequest,
		)
		return
	}

	_, recipientExists := h.MemberStore.Get(
		newTask.RecipientID,
	)
	if !recipientExists {
		http.Error(
			w,
			"Recipient member not found",
			http.StatusBadRequest,
		)
		return
	}

	createdTask := h.TaskStore.Create(newTask)

	h.TaskQueue <- createdTask

	writeJSON(
		w,
		http.StatusAccepted,
		createdTask,
	)
}

func (h *TaskHandler) ListTasks(
	w http.ResponseWriter,
) {
	tasks := h.TaskStore.GetAll()

	responses := newTaskResponses(
		tasks,
		h.MemberStore,
	)

	writeJSON(
		w,
		http.StatusOK,
		responses,
	)
}

func (h *TaskHandler) GetTask(
	w http.ResponseWriter,
	r *http.Request,
) {
	if r.Method != http.MethodGet {
		http.Error(
			w,
			"Method not allowed",
			http.StatusMethodNotAllowed,
		)
		return
	}

	idString := strings.TrimPrefix(
		r.URL.Path,
		"/tasks/",
	)

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(
			w,
			"Invalid task ID",
			http.StatusBadRequest,
		)
		return
	}

	createdTask, exists := h.TaskStore.Get(id)
	if !exists {
		http.Error(
			w,
			"Task not found",
			http.StatusNotFound,
		)
		return
	}

	response := newTaskResponse(
		createdTask,
		h.MemberStore,
	)

	writeJSON(
		w,
		http.StatusAccepted,
		response,
	)
}

func writeJSON(
	w http.ResponseWriter,
	statusCode int,
	data any,
) {
	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(
			w,
			"Unable to encode response",
			http.StatusInternalServerError,
		)
	}
}
