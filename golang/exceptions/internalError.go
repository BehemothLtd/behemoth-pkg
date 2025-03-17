package exceptions

import (
	"fmt"

	translator "github.com/BehemothLtd/behemoth-pkg/golang/translators"
)

const (
	InternalErrorCode = 500
)

// InternalError represents a bad request error.
type InternalError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error returns the error message.
func (e InternalError) Error() string {
	return fmt.Sprintf("error [%d]: %s", e.Code, e.Message)
}

// Extensions returns additional data associated with the error.
func (e InternalError) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code":    e.Code,
		"message": e.Message,
	}
}

// NewInternalError creates a new InternalError instance with the provided message.
// If the message is empty, it uses the default error message.
func NewInternalError(message *string) InternalError {
	var returnMessage string

	if message == nil || *message == "" {
		returnMessage = translator.Translate(nil, "errExceptionMsg_internal")
	} else {
		returnMessage = *message
	}

	return InternalError{
		Code:    InternalErrorCode,
		Message: returnMessage,
	}
}
