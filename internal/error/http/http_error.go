package http

import (
	"errors"
	"fmt"
	"net/http"
	"sola-test-task/pkg/context"
)

var (
	errBadRequest = map[context.Locale]error{
		context.English: errors.New("Bad request input. Please, provide valid data"),
		context.Spanish: errors.New(
			"Entrada de solicitud incorrecta. Por favor, proporcione datos válidos",
		),
	}

	errUnauthorized = map[context.Locale]error{
		context.English: errors.New("Access denied. Please, authorize"),
		context.Spanish: errors.New("Acceso denegado. Por favor, autorice"),
	}

	errConflict = map[context.Locale]error{
		context.English: errors.New("Please, try providing alternative input"),
		context.Spanish: errors.New("Por favor, intente proporcionar una entrada alternativa"),
	}

	errNotFound = map[context.Locale]error{
		context.English: errors.New(
			"Sorry, but it seems like what you have requested does not exist. Please, try providing different input",
		),
		context.Spanish: errors.New(
			"Lo siento, pero parece que lo que ha solicitado no existe. Por favor, intente proporcionar una entrada diferente",
		),
	}

	errInternal = map[context.Locale]error{
		context.English: errors.New(
			"We are sorry, but something went wrong. Please, try again later",
		),
		context.Spanish: errors.New(
			"Lo siento, pero algo salió mal. Por favor, inténtelo de nuevo más tarde",
		),
	}

	errForbidden = map[context.Locale]error{
		context.English: errors.New("Sorry, but the action you have tried to take is forbidden"),
		context.Spanish: errors.New(
			"Lo siento, pero la acción que intentó realizar está prohibida",
		),
	}
)

type ErrorHttp interface {
	error
	Code() int
}

type errorHttp struct {
	msg  error
	code int
}

func (e errorHttp) Error() string {
	return e.msg.Error()
}

func (e errorHttp) Code() int {
	return e.code
}

func NewErrConflict(errToWrap error, loc context.Locale) ErrorHttp {
	return errorHttp{
		msg:  fmt.Errorf("%w: %w", errConflict[loc], errToWrap),
		code: http.StatusConflict,
	}
}

func NewErrUnauthorized(errToWrap error, loc context.Locale) ErrorHttp {
	return errorHttp{
		msg:  fmt.Errorf("%w: %w", errUnauthorized[loc], errToWrap),
		code: http.StatusUnauthorized,
	}
}

func NewErrForbidden(errToWrap error, loc context.Locale) ErrorHttp {
	return errorHttp{
		msg:  fmt.Errorf("%w: %w", errForbidden[loc], errToWrap),
		code: http.StatusForbidden,
	}
}

func NewErrBadRequest(errToWrap error, loc context.Locale) ErrorHttp {
	return errorHttp{
		msg:  fmt.Errorf("%w: %w", errBadRequest[loc], errToWrap),
		code: http.StatusBadRequest,
	}
}

func NewErrInternal(errToWrap error, loc context.Locale) ErrorHttp {
	return errorHttp{
		msg:  fmt.Errorf("%w: %w", errInternal[loc], errToWrap),
		code: http.StatusInternalServerError,
	}
}

func NewErrNotFound(errToWrap error, loc context.Locale) ErrorHttp {
	return errorHttp{
		msg:  fmt.Errorf("%w: %w", errNotFound[loc], errToWrap),
		code: http.StatusNotFound,
	}
}
