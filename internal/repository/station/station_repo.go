package station

import (
	domain "sola-test-task/internal/domain/station"
	derr "sola-test-task/internal/error/data"
	"sola-test-task/pkg/context"
)

type StationRepo interface {
	CreateStation(c *context.Context, newSt *domain.CreateStation) (*domain.Station, derr.DataError)
}
