package exceptions

import (
	"fmt"
)

const (
	UnauthorizedErrorCode = 401
)

// UnauthorizedError represents an unauthorized access error.
type UnauthorizedError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error returns the error message.
func (e *UnauthorizedError) Error() string {
	return fmt.Sprintf("error [%d]: %s", e.Code, e.Message)
}

// Extensions returns additional data associated with the error.
func (e *UnauthorizedError) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code":    e.Code,
		"message": e.Message,
	}
}

// NewUnauthorizedError creates a new UnauthorizedError instance with the provided message.
// If the message is empty, it uses the default error message.
func NewUnauthorizedError(message *string) *UnauthorizedError {
	var returnMessage string

	if message == nil || *message == "" {
		returnMessage = "errExceptionMsg_unauthorized"
	} else {
		returnMessage = *message
	}

	return &UnauthorizedError{
		Code:    UnauthorizedErrorCode,
		Message: returnMessage,
	}
}
