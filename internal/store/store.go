package store

import (
	"context"
	"errors"

	"github.com/NemanjaKajzer/TaskOrchestrator/internal/model"
	"github.com/google/uuid"
)

var ErrNotFound = errors.New("store: record not found")

type TaskStore interface {
	Create(ctx context.Context, task *model.Task) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.Task, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status model.TaskStatus) error
	List(ctx context.Context, limit, offset int) ([]*model.Task, error)
}

type Closer interface {
	Close() error
}

type Store interface {
	TaskStore
	Closer
}
