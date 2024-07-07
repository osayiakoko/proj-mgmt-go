package data

import (
	"database/sql"
	"errors"
	"time"
)

// Define a custom ErrRecordNotFound error. We'll return this from our Get() method when
// looking up a movie that doesn't exist in our database.
var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

// Create a Stores struct which wraps the TaskModel. We'll add other models to this,
// like a UserStore and PermissionStore, as our build progresses.
type Stores struct {
	Tasks interface {
		Create(*Task) error
		GetAll(string, string, string, Filters) (PaginatedData[Tasks], error)
		Get(int64) (*Task, error)
		Update(*Task) error
		Delete(int64) error
	}
	Users interface {
		GetByEmail(string) (*User, error)
		Insert(*User) error
		Update(*User) error
		GetForToken(string, string) (*User, error)
	}
	Tokens interface {
		New(int64, time.Duration, string) (*Token, error)
		Insert(*Token) error
		DeleteAllForUser(string, int64) error
	}
	Permissions interface {
		GetAllForUser(int64) (Permissions, error)
		AddForUser(int64, ...string) error
	}
}

// For ease of use, we also add a New() method which returns a Models struct containing
// the initialized TaskStore.
func NewStore(db *sql.DB) Stores {
	return Stores{
		Tasks:       TaskStore{DB: db},
		Users:       UserStore{DB: db},
		Tokens:      TokenStore{DB: db},
		Permissions: PermissionStore{DB: db},
	}
}
