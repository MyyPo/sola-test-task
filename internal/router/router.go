package router

import (
	stdCtx "context"
	"fmt"
	"sola-test-task/internal/config"
	"sola-test-task/internal/controller"
	"sola-test-task/internal/dto/response"
	errHttp "sola-test-task/internal/error/http"
	"sola-test-task/internal/middleware"
	transServ "sola-test-task/internal/service/translation"
	valServ "sola-test-task/internal/service/validation"
	"sola-test-task/pkg/context"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Router struct {
	log  *zap.Logger
	conf *config.Config

	reqCtxMid middleware.RequestCtxMiddleware

	valServ   valServ.ValidationService
	transServ transServ.TranslationsService

	stCont controller.StationController
}

func NewRouter(
	log *zap.Logger,
	conf *config.Config,

	reqCtxMid middleware.RequestCtxMiddleware,

	valServ valServ.ValidationService,
	transServ transServ.TranslationsService,

	stCont controller.StationController,
) Router {
	return Router{
		log,
		conf,

		reqCtxMid,

		valServ,
		transServ,

		stCont,
	}
}

func (r *Router) SetupHandlers() error {
	h := gin.Default()

	if err := r.valServ.RegisterValidations(); err != nil {
		return err
	}

	stationsGroup := h.Group("/stations")
	stationsGroup.Use(r.reqCtxMid.RequestCtx())

	{
		stationsGroup.POST("", r.CreateStation)
	}

	return h.Run(r.conf.Server.Address())
}

func transBindJSON[T any](
	c *context.Context,
	ginCtx *gin.Context,
	req *T,
	tsServ transServ.TranslationsService,
) errHttp.ErrorHttp {
	if err := ginCtx.BindJSON(req); err != nil {
		switch c.LocaleOrDefault() {
		case context.English:
			fmt.Println("english locale: ", c.LocaleOrDefault())
			return tsServ.TranslateEN(err)
		case context.Spanish:
			fmt.Println("spanish locale: ", c.LocaleOrDefault())
			return tsServ.TranslateES(err)
		default:
			fmt.Println("default locale: ", c.LocaleOrDefault())
			return tsServ.TranslateEN(err)
		}
	}
	return nil
}

func sendResponse(c *context.Context, ginCtx *gin.Context, statusCode int, data any, err error) {
	var errString string
	if err != nil {
		errString = err.Error()
	}
	resp := response.NewBaseResponse(data, errString, c.RequestIDOrDefault())
	ginCtx.JSON(statusCode, resp)
}

func (r *Router) context(stdCtx stdCtx.Context) *context.Context {
	val := stdCtx.Value(middleware.LoggerKey)
	if val == nil {
		return context.NewContext(stdCtx, r.log)
	}
	if log, ok := val.(*zap.Logger); !ok {
		return context.NewContext(stdCtx, r.log)
	} else {
		return context.NewContext(stdCtx, log)
	}
}
