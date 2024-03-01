package jet

import (
	"errors"
	derr "sola-test-task/internal/error/data"

	"github.com/lib/pq"
)

func ErrSpec(jetErr error) (derr.DataErrorType, string) {
	var pqErr *pq.Error
	if errors.As(jetErr, &pqErr) {
		switch pqErr.Code {
		case uniqueViolationCode:
			return derr.Conflict, pqErr.Constraint
		case doesNotExistCode:
			return derr.NotFound, ""
		default:
			return derr.Internal, ""
		}
	}
	return derr.Internal, ""
}

const (
	uniqueViolationCode = "23505"
	doesNotExistCode    = "42703"
)
