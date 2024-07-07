package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/osayiakoko/project-mgmt-sys/internal/validator"
)

const (
	TaskReadPermission  = "tasks:read"
	TaskWritePermission = "tasks:write"
)

// Task Struct
type Task struct {
	ID          int64        `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Priority    TaskPriority `json:"priority"`
	Status      TaskStatus   `json:"status"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
}

type Tasks []*Task

func ValidateTask(v *validator.Validator, task *Task) {
	v.Check(task.Title != "", "title", "must be provided")
	v.Check(task.Description != "", "description", "must be provided")
	v.Check(task.Priority != "", "priority", "must be provided")
	v.Check(task.Status != "", "status", "must be provided")
}

// Task Store
type TaskStore struct {
	DB *sql.DB
}

func (t TaskStore) Create(task *Task) error {
	query := `
		INSERT INTO tasks (title, description, priority, status)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
	`

	args := []any{task.Title, task.Description, task.Priority, task.Status}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return t.DB.QueryRowContext(ctx, query, args...).Scan(&task.ID, &task.CreatedAt, &task.UpdatedAt)
}

func (t TaskStore) GetAll(title string, priority string, status string, filters Filters) (PaginatedData[Tasks], error) {
	query := fmt.Sprintf(`
		SELECT count(*) OVER(), id, title, description, priority, status, created_at, updated_at
		FROM tasks
		WHERE (to_tsvector('simple', title) @@ plainto_tsquery('simple', $1) or $1 = '')
		AND (LOWER(priority) = LOWER($2) or $2 = '')
		AND (LOWER(status) = LOWER($3) or $3 = '')
		ORDER BY %s %s, id ASC
		LIMIT $4 OFFSET $5`, filters.sortColumn(), filters.sortDirection())

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := t.DB.QueryContext(ctx, query, title, priority, status,
		filters.limit(), filters.offset())
	if err != nil {
		return PaginatedData[Tasks]{}, err
	}

	// Importantly, defer a call to rows.Close() to ensure that the resultset is closed
	// before GetAll() returns.
	defer rows.Close()

	totalRecords := 0
	tasks := Tasks{}

	for rows.Next() {
		var task Task

		err := rows.Scan(
			&totalRecords,
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Priority,
			&task.Status,
			&task.CreatedAt,
			&task.UpdatedAt,
		)

		if err != nil {
			return PaginatedData[Tasks]{}, err
		}

		tasks = append(tasks, &task)
	}

	// When the rows.Next() loop has finished, call rows.Err() to retrieve any error
	// that was encountered during the iteration.
	if err = rows.Err(); err != nil {
		return PaginatedData[Tasks]{}, err
	}

	// Generate a Metadata struct, passing in the total record count and pagination
	// parameters from the client.
	paginatedTasks := paginateData(tasks, totalRecords, filters.Page, filters.PageSize)

	return paginatedTasks, nil
}

func (t TaskStore) Get(id int64) (*Task, error) {
	query := `
		SELECT id, title, description, priority, status, created_at, updated_at
		FROM tasks
		WHERE id = $1 
	`
	var task Task

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := t.DB.QueryRowContext(ctx, query, id).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Priority,
		&task.Status,
		&task.CreatedAt,
		&task.UpdatedAt,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &task, nil
}

func (t TaskStore) Update(task *Task) error {
	query := `
		UPDATE tasks
		SET title = $1, description = $2, priority = $3, status = $4, updated_at = CURRENT_TIMESTAMP
		WHERE id = $5
		RETURNING created_at
	`

	args := []any{task.Title, task.Description, task.Priority, task.Status, task.ID}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := t.DB.QueryRowContext(ctx, query, args...).Scan(&task.CreatedAt)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil
}

func (t TaskStore) Delete(id int64) error {
	query := `
		DELETE FROM tasks
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := t.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

// MOCK STORE
type MockTaskStore struct{}

func (m MockTaskStore) Create(movie *Task) error {
	return nil // Mock the action...
}

func (t MockTaskStore) GetAll() (*[]Task, error) {
	return nil, nil
}

func (m MockTaskStore) Get(id int64) (*Task, error) {
	return nil, nil // Mock the action...
}

func (m MockTaskStore) Update(task *Task) error {
	return nil // Mock the action...
}

func (m MockTaskStore) Delete(id int64) error {
	return nil // Mock the action...
}
