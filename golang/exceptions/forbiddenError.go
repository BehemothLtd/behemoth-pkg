package exceptions

import (
	translator "behemoth-pkg/golang/translators"
	"fmt"
)

const (
	ForbiddenErrorCode = 403
)

// ForbiddenError represents a forbidden access error.
type ForbiddenError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error returns the error message.
func (e *ForbiddenError) Error() string {
	return fmt.Sprintf("error [%d]: %s", e.Code, e.Message)
}

// Extensions returns additional data associated with the error.
func (e *ForbiddenError) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code":    e.Code,
		"message": e.Message,
	}
}

// NewForbiddenError creates a new ForbiddenError instance with the provided message.
// If the message is empty, it uses the default error message.
func NewForbiddenError(message *string) *ForbiddenError {
	var returnMessage string

	if message == nil || *message == "" {
		returnMessage = translator.Translate(nil, "errExceptionMsg_forbidden")
	} else {
		returnMessage = *message
	}

	return &ForbiddenError{
		Code:    ForbiddenErrorCode,
		Message: returnMessage,
	}
}
