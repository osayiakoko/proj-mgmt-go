package data

import (
	"database/sql"
	"errors"
)

// Define a custom ErrRecordNotFound error. We'll return this from our Get() method when
// looking up a movie that doesn't exist in our database.
var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict = errors.New("edit conflict")
)

// create a store interface
type Store[T any] interface {
	Create(*T) error 
	GetAll(string, string, string, Filters) ([]*T, error) 
	Get(int64) (*T, error) 
	Update(*T) error 
	Delete(int64) error
}

// Create a Stores struct which wraps the TaskModel. We'll add other models to this, 
// like a UserStore and PermissionStore, as our build progresses.
type Stores struct {
	Tasks Store[Task] 
}

// For ease of use, we also add a New() method which returns a Models struct containing 
// the initialized TaskStore.
func NewStore(db *sql.DB) Stores {
	return Stores{
		Tasks: TaskStore{DB: db},
	} 
}