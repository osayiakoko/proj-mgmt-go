package data

import (
	"errors"
	"strconv"
)

type TaskPriority string

const (
	High   TaskPriority = "high"
	Medium TaskPriority = "medium"
	Low    TaskPriority = "low"
)

var ErrInvalidTaskPriority = errors.New("invalid task priority")

func isValidPriority(status TaskPriority) bool {
	switch status {
	case High, Medium, Low:
		return true
	}
	return false
}

func (t *TaskPriority) UnmarshalJSON(jsonValue []byte) error {
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidTaskPriority
	}

	status := TaskPriority(unquotedJSONValue)

	if valid := isValidPriority(status); !valid {
		return ErrInvalidTaskPriority
	}

	*t = status

	return nil
}
