package api

import (
	"time"

	"concurrent-task-queue-server/internal/member"
	"concurrent-task-queue-server/internal/task"
)

type TaskResponse struct {
	ID            int       `json:"id"`
	SenderID      int       `json:"sender_id"`
	SenderName    string    `json:"sender_name"`
	RecipientID   int       `json:"recipient_id"`
	RecipientName string    `json:"recipient_name"`
	Type          string    `json:"type"`
	Payload       string    `json:"payload"`
	Status        string    `json:"status"`
	Attempts      int       `json:"attempts"`
	MaxRetries    int       `json:"max_retries"`
	Error         string    `json:"error,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func newTaskResponse(
	t task.Task,
	memberStore *member.Store,
) TaskResponse {
	senderName := memberName(
		memberStore,
		t.SenderID,
	)

	recipientName := memberName(
		memberStore,
		t.RecipientID,
	)

	return TaskResponse{
		ID:            t.ID,
		SenderID:      t.SenderID,
		SenderName:    senderName,
		RecipientID:   t.RecipientID,
		RecipientName: recipientName,
		Type:          t.Type,
		Payload:       t.Payload,
		Status:        t.Status,
		Attempts:      t.Attempts,
		MaxRetries:    t.MaxRetries,
		Error:         t.Error,
		CreatedAt:     t.CreatedAt,
		UpdatedAt:     t.UpdatedAt,
	}
}

func newTaskResponses(
	tasks []task.Task,
	memberStore *member.Store,
) []TaskResponse {
	responses := make(
		[]TaskResponse,
		0,
		len(tasks),
	)

	for _, t := range tasks {
		responses = append(
			responses,
			newTaskResponse(t, memberStore),
		)
	}

	return responses
}

func memberName(
	memberStore *member.Store,
	memberID int,
) string {
	foundMember, exists := memberStore.Get(memberID)
	if !exists {
		return "Unknown Member"
	}

	return foundMember.Name
}
