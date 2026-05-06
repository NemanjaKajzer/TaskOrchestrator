package sql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/NemanjaKajzer/TaskOrchestrator/internal/model"
	"github.com/NemanjaKajzer/TaskOrchestrator/internal/store"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type PostgresTaskStore struct {
	db *sqlx.DB
}

func NewPostgresTaskStore(db *sqlx.DB) *PostgresTaskStore {
	return &PostgresTaskStore{db: db}
}

func (s *PostgresTaskStore) Create(ctx context.Context, task *model.Task) error {
	query := `
				INSERT INTO tasks (title, status, payload)
				VALUES (:title, :status, :payload)
			  	RETURNING id, created_at, updated_at
			  `
	rows, err := s.db.NamedQueryContext(ctx, query, task)
	if err != nil {
		return err
	}

	defer rows.Close()

	if rows.Next() {
		return rows.Scan(&task.ID, &task.CreatedAt, &task.UpdatedAt)
	}

	return nil
}

func (s *PostgresTaskStore) GetByID(ctx context.Context, id uuid.UUID) (*model.Task, error) {
	query := `
		SELECT id, title, status, payload, created_at, updated_at
		FROM tasks
		WHERE id = $1
	`

	task := &model.Task{}
	err := s.db.GetContext(ctx, task, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrNotFound
		}
		return nil, err
	}

	return task, nil
}

func (s *PostgresTaskStore) UpdateStatus(ctx context.Context, id uuid.UUID, status model.TaskStatus) error {
	query := `
		UPDATE tasks
		SET    status     = :status,
		       updated_at = now()
		WHERE  id         = :id
	`

	args := map[string]any{
		"id":     id,
		"status": status,
	}

	result, err := s.db.NamedExecContext(ctx, query, args)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return store.ErrNotFound
	}

	return nil
}

func (s *PostgresTaskStore) List(ctx context.Context, limit, offset int) ([]*model.Task, error) {
	query := `
		SELECT id, title, status, payload, created_at, updated_at
		FROM   tasks
		ORDER  BY created_at DESC
		LIMIT  $1 OFFSET $2
	`

	tasks := []*model.Task{}
	err := s.db.SelectContext(ctx, &tasks, query, limit, offset)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *PostgresTaskStore) Close() error {
	return s.db.Close()
}
