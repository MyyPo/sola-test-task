package validation

import (
	stReq "sola-test-task/internal/dto/request/station"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type ValidationService interface {
	RegisterValidations() error
}

type validationService struct{}

func NewValidationService() ValidationService {
	return &validationService{}
}

func (r *validationService) RegisterValidations() (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterStructValidation(stReq.CreateStationValidation, stReq.CreateStation{})
	}
	return
}
