package data

import (
	"errors"
	"strconv"
)

type TaskStatus string

const (
	Todo    TaskStatus = "todo"
	Ongoing TaskStatus = "ongoing"
	Done    TaskStatus = "done"
)

var ErrInvalidTaskStatus = errors.New("invalid task status")

func isValidStatus(status TaskStatus) bool {
	switch status {
	case Todo, Ongoing, Done:
		return true
	}
	return false
}

func (t *TaskStatus) UnmarshalJSON(jsonValue []byte) error {
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidTaskStatus
	}

	status := TaskStatus(unquotedJSONValue)

	if valid := isValidStatus(status); !valid {
		return ErrInvalidTaskStatus
	}

	*t = status

	return nil
}
