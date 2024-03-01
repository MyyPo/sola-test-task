package context

import (
	"context"

	"go.uber.org/zap"
)

type Context struct {
	context.Context
	*zap.Logger
}

func NewContext(ctx context.Context, log *zap.Logger) *Context {
	return &Context{
		Context: ctx,
		Logger:  log,
	}
}

func (c *Context) RequestIDOrDefault() string {
	reqId, ok := c.Value(RequestIDKey).(string)
	if !ok {
		reqId = unknownReqID
	}
	return reqId
}

func (c *Context) LocaleOrDefault() Locale {
	reqLoc, ok := c.Value(RequestLocaleKey).(string)
	if !ok {
		reqLoc = English
	}

	switch reqLoc {
	case English:
		return English
	case Spanish:
		return Spanish
	default:
		return English
	}
}

const (
	RequestIDKey string = "RequestIDKey"
	unknownReqID string = "Unknown request ID"
)

const (
	RequestLocaleKey string = "RequestLocaleKey"
)

type Locale = string

const (
	English Locale = "en"
	Spanish Locale = "es"
)

var ImplementedLocales []Locale = []Locale{English, Spanish}

const (
	RequestLoggerKey              = "RequestLoggerKey"
	StructuredLoggingRequestIDKey = "request_id"
)
