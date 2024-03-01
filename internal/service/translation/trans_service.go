package translation

import (
	"errors"
	"fmt"
	errHttp "sola-test-task/internal/error/http"
	"sola-test-task/pkg/context"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/es"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	es_translations "github.com/go-playground/validator/v10/translations/es"
)

type TranslationsService interface {
	TranslateEN(errs error) errHttp.ErrorHttp
	TranslateES(errs error) errHttp.ErrorHttp
}

type errTranslationsService struct {
	enTrans ut.Translator
	esTrans ut.Translator
}

func NewTranslationService() (TranslationsService, error) {
	serv := errTranslationsService{}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		es := es.New()
		uni := ut.New(en, en, es)

		serv.enTrans, _ = uni.GetTranslator(context.English)
		en_translations.RegisterDefaultTranslations(v, serv.enTrans)

		serv.esTrans, _ = uni.GetTranslator(context.Spanish)
		es_translations.RegisterDefaultTranslations(v, serv.esTrans)
	}

	return &serv, nil
}

const FailedToValidate string = "unhandled input validation error has occured"

func (s *errTranslationsService) translate(
	errs error,
	ut ut.Translator,
) errHttp.ErrorHttp {
	valErrs, ok := errs.(validator.ValidationErrors)
	if !ok {
		return errHttp.NewErrBadRequest(ErrFailedToParse, ut.Locale())
	}

	transErrs := valErrs.Translate(ut)
	transErrsArr := make([]string, 0, len(transErrs))
	for _, v := range transErrs {
		transErrsArr = append(transErrsArr, v)
	}

	return errHttp.NewErrBadRequest(fmt.Errorf("%s", strings.Join(transErrsArr, ". ")), ut.Locale())
}

func (s *errTranslationsService) TranslateEN(
	errs error,
) errHttp.ErrorHttp {
	return s.translate(errs, s.enTrans)
}

func (s *errTranslationsService) TranslateES(
	errs error,
) errHttp.ErrorHttp {
	return s.translate(errs, s.esTrans)
}

var ErrFailedToParse error = errors.New("failed to parse incoming request, invalid json provided")
