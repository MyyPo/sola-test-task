package request

import (
	shareConn "sola-test-task/internal/domain/connector/share"
	shareSt "sola-test-task/internal/domain/station/share"

	"github.com/go-playground/validator/v10"
)

type CreateStation struct {
	UID        string            `json:"uid"         binding:"required,min=5,max=128"`
	Public     *bool             `json:"public"      binding:"required"`
	Title      *string           `json:"title"       binding:"required_if=Public true,omitempty,min=3,max=255"`
	Desc       *string           `json:"description" binding:"required_if=Public true,omitempty,min=3,max=2048"`
	Address    *string           `json:"address"     binding:"required_if=Public true,omitempty,min=10,max=1024"`
	Coords     *shareSt.Point    `json:"coordinates" binding:"required_if=Public true,omitempty,dive"`
	Connectors []createConnector `json:"connectors"  binding:"required,min=1,max=8,dive,required"`
}

type createConnector struct {
	ID         string                  `json:"id"           binding:"required,min=5,max=128"`
	Type       shareConn.ConnectorType `json:"type"         binding:"required,oneof=CCS CHAdeMO Type1 Type2"`
	MaxPowerKw *float64                `json:"max_power_kw" binding:"required,gte=1,lte=500"`
}

func CreateStationValidation(sl validator.StructLevel) {
	cs := sl.Current().Interface().(CreateStation)

	if cs.Coords != nil {
		if len(cs.Coords) != 2 {
			sl.ReportError(cs.Coords, "coordinates", "coordinates", "required", "")
		}
		lat := cs.Coords[0]
		if lat > 90 || lat < -90 {
			sl.ReportError(cs.Coords, "coordinates", "coordinates", "latitude", "")
		}
		long := cs.Coords[1]
		if long > 180 || long < -180 {
			sl.ReportError(cs.Coords, "coordinates", "coordinates", "longitude", "")
		}
	}
}
