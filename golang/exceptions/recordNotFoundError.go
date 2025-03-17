package exceptions

import (
	"fmt"

	translator "github.com/BehemothLtd/behemoth-pkg/golang/translators"
)

const (
	NotFoundErrorCode = 404
)

// RecordNotFoundError represents a record not found error.
type RecordNotFoundError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error returns the error message.
func (e *RecordNotFoundError) Error() string {
	return fmt.Sprintf("error [%d]: %s", e.Code, e.Message)
}

// Extensions returns additional data associated with the error.
func (e *RecordNotFoundError) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code":    e.Code,
		"message": e.Message,
	}
}

// NewRecordNotFoundError creates a new RecordNotFoundError instance with the default error message and code.
func NewRecordNotFoundError() *RecordNotFoundError {
	return &RecordNotFoundError{
		Code:    NotFoundErrorCode,
		Message: translator.Translate(nil, "errDbMsg_notFound"),
	}
}
