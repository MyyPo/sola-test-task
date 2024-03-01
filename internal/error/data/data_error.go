package data

import (
	"fmt"
	"slices"
	errVerb "sola-test-task/internal/error/verbose"
	"sola-test-task/pkg/context"
)

type DataError interface {
	errVerb.VerboseError
	NotFound() bool
	Conflict() bool
	BadRequest() bool
}

type dataError struct {
	userErr error
	techErr error
	errType DataErrorType
}

func (e *dataError) NotFound() bool {
	return e.errType == NotFound
}

func (e *dataError) Conflict() bool {
	return e.errType == Conflict
}

func (e *dataError) BadRequest() bool {
	return e.errType == BadRequest
}

func (e *dataError) Error() string {
	return e.userErr.Error()
}

func (e *dataError) Verbose() error {
	return e.techErr
}

type DataErrorType int

const (
	NotFound DataErrorType = iota
	Conflict
	BadRequest
	Internal
)

func newErrNotFound(loc context.Locale, modName string, techErr error) DataError {
	return &dataError{
		userErr: errNotFound[loc](modName),
		techErr: techErr,
		errType: NotFound,
	}
}

func newErrConflict(loc context.Locale, modName, fieldName string, techErr error) DataError {
	return &dataError{
		userErr: errConflict[loc](modName, fieldName),
		techErr: techErr,
		errType: Conflict,
	}
}

func newErrBadRequest(loc context.Locale, modName string, techErr error) DataError {
	return &dataError{
		userErr: errBadRequest[loc](modName),
		techErr: techErr,
		errType: BadRequest,
	}
}

func newErrInternal(loc context.Locale, modName string, techErr error) DataError {
	return &dataError{
		userErr: errInternal[loc](modName),
		techErr: techErr,
		errType: Internal,
	}
}

var (
	errBadRequest = map[context.Locale]func(string) error{
		context.English: func(n string) error { return fmt.Errorf("the %s related request is invalid", n) },
		context.Spanish: func(n string) error { return fmt.Errorf("la solicitud relacionada con %s no es válida", n) },
	}

	errNotFound = map[context.Locale]func(string) error{
		context.English: func(n string) error { return fmt.Errorf("the requested %s does not exist", n) },
		context.Spanish: func(n string) error { return fmt.Errorf("el %s solicitado no existe", n) },
	}

	errConflict = map[context.Locale]func(string, string) error{
		context.English: func(n, f string) error {
			return fmt.Errorf("the provided %s data is using duplicate value for %s", n, f)
		},
		context.Spanish: func(n, f string) error {
			return fmt.Errorf(
				"los datos proporcionados para %s están utilizando algunos valores duplicados %s",
				n, f,
			)
		},
	}

	errInternal = map[context.Locale]func(string) error{
		context.English: func(n string) error { return fmt.Errorf("unexpected error occured when performing operation on %s", n) },
		context.Spanish: func(n string) error {
			return fmt.Errorf("se produjo un error inesperado al realizar la operación en %s", n)
		},
	}
)

func NewErr(
	errTyp DataErrorType,
	techErr error,
	modName string,
	fieldName string,
	loc context.Locale,
	expTypes ...DataErrorType,
) DataError {
	if !slices.Contains(expTypes, errTyp) {
		return newErrInternal(loc, modName, techErr)
	}

	switch errTyp {
	case Conflict:
		return newErrConflict(loc, modName, fieldName, techErr)
	case NotFound:
		return newErrNotFound(loc, modName, techErr)
	case BadRequest:
		return newErrBadRequest(loc, modName, techErr)
	default:
		return newErrInternal(loc, modName, techErr)
	}
}
