package member

import "time"

type Member struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Contact   string    `json:"contact,omitempty"`
	Role      string    `json:"role"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
