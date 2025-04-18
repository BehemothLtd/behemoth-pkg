package exceptions

import (
	"fmt"

	translator "github.com/BehemothLtd/behemoth-pkg/golang/translators"
)

const (
	BadRequestErrorCode = 400
)

// BadRequestError represents a bad request error.
type BadRequestError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error returns the error message.
func (e BadRequestError) Error() string {
	return fmt.Sprintf("error [%d]: %s", e.Code, e.Message)
}

// Extensions returns additional data associated with the error.
func (e BadRequestError) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code":    e.Code,
		"message": e.Message,
	}
}

// NewBadRequestError creates a new BadRequestError instance with the provided message.
// If the message is empty, it uses the default error message.
func NewBadRequestError(message *string) BadRequestError {
	var returnMessage string
	if message == nil || *message == "" {
		returnMessage = translator.Translate(nil, "errExceptionMsg_badRequest")
	} else {
		returnMessage = *message
	}

	return BadRequestError{
		Code:    BadRequestErrorCode,
		Message: returnMessage,
	}
}
