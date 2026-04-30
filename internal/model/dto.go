package model

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type CreateTaskRequest struct {
	Title   string          `json:"title"`
	Payload json.RawMessage `json:"payload"`
}

type TaskResponse struct {
	ID        uuid.UUID       `json:"id"`
	Title     string          `json:"title"`
	Status    string          `json:"status"`
	Payload   json.RawMessage `json:"payload"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("api error %d: %s", e.Code, e.Message)
}
