package model

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type TaskStatus string

const (
	StatusPending   TaskStatus = "pending"
	StatusRunning   TaskStatus = "running"
	StatusCompleted TaskStatus = "completed"
	StatusFailed    TaskStatus = "failed"
)

type Task struct {
	ID        uuid.UUID       `json:"id" 			db:"id"`
	Title     string          `json:"title" 		db:"title"`
	Status    TaskStatus      `json: "status" 		db:"status"`
	Payload   json.RawMessage `json:"payload" 		db:"payload"`
	CreatedAt time.Time       `json:"created_at" 	db:"created_at"`
	UpdatedAt time.Time       `json:"updated_at" 	db:"updated_at"`
}

func (t Task) IsTerminal() bool {
	return t.Status == StatusCompleted || t.Status == StatusFailed
}
