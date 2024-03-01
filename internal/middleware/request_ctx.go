package middleware

import (
	stdCtx "context"
	"slices"
	"sola-test-task/pkg/context"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type RequestCtxMiddleware interface {
	RequestCtx() gin.HandlerFunc
}

type requestCtx struct {
	log *zap.Logger
}

func NewRequestCtx(log *zap.Logger) RequestCtxMiddleware {
	return &requestCtx{
		log: log,
	}
}

func (m *requestCtx) RequestCtx() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reqId := ctx.GetHeader(HeaderRequestID)
		if len(reqId) == 0 {
			reqId = uuid.NewString()
		}
		langHead := ctx.GetHeader(HeaderLanguage)
		if len(langHead) == 0 {
			langHead = defaultLocale
		}

		loc := prioritizedLocalization(langHead)

		ctx.Set(context.RequestIDKey, reqId)
		ctx.Set(context.RequestLocaleKey, loc)
		ctx.Set(
			context.RequestLoggerKey,
			m.log.With(zap.String(context.StructuredLoggingRequestIDKey, reqId)),
		)

		ctx.Next()
	}
}

const (
	HeaderRequestID string = "X-Request-ID"
	HeaderLanguage  string = "Accept-Language"

	LoggerKey string = "LoggerKey"
	LocaleKey string = "LocaleKey"

	defaultLocale string = "en"
)

func AppCtx(ctx stdCtx.Context) *context.Context {
	val := ctx.Value(LoggerKey)
	if val != nil {
		log, ok := val.(*zap.Logger)
		if ok {
			return context.NewContext(ctx, log)
		}
	}
	return nil
}

func prioritizedLocalization(langHeader string) context.Locale {
	locPr := parseAcceptLanguage(langHeader)
	loc := defaultLocale
	var currMaxPr float64

	for lkey, lpr := range locPr {
		if slices.Contains(context.ImplementedLocales, lkey) && lpr > currMaxPr {
			loc = lkey
			currMaxPr = lpr
		}
	}

	return loc
}

type localePriority map[string]float64

func parseAcceptLanguage(langHeader string) localePriority {
	locPriors := make(localePriority)

	langStrs := strings.Split(langHeader, ",")
	for _, langStr := range langStrs {
		trimedLangStr := strings.Trim(langStr, " ")

		lang := strings.Split(trimedLangStr, ";")
		langKey := strings.Split(lang[0], "-")[0]

		if len(lang) == 1 {
			locPriors[langKey] = 1
		} else {
			langVal := lang[1]

			locPriorValStr := strings.Split(langVal, "=")
			locPriorVal, err := strconv.ParseFloat(locPriorValStr[1], 64)
			if err != nil {
				continue
			}
			locPriors[langKey] = locPriorVal
		}
	}
	return locPriors
}
