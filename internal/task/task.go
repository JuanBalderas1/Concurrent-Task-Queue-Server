package task

import "time"

type Task struct {
	ID          int       `json:"id"`
	SenderID    int       `json:"sender_id"`
	RecipientID int       `json:"recipient_id"`
	Type        string    `json:"type"`
	Payload     string    `json:"payload"`
	Status      string    `json:"status"`
	Attempts    int       `json:"attempts"`
	MaxRetries  int       `json:"max_retries"`
	Error       string    `json:"error,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
