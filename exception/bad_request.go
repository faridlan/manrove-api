package exception

import (
	"strings"
)

type BadRequestError struct {
	Message []string
}

func NewBadRequestError(msg []string) *BadRequestError {
	return &BadRequestError{
		Message: msg,
	}
}

func (b *BadRequestError) Error() string {
	return strings.Join(b.Message, ", ")
	// return fmt.Sprintf("Error: %v", b.Message)
}
