package controller

import (
	stDom "sola-test-task/internal/domain/station"
	request "sola-test-task/internal/dto/request/station"
	response "sola-test-task/internal/dto/response/station"
	errHttp "sola-test-task/internal/error/http"
	stationRepo "sola-test-task/internal/repository/station"
	"sola-test-task/pkg/context"

	"go.uber.org/zap"
)

type StationController interface {
	CreateStation(
		c *context.Context,
		req *request.CreateStation,
	) (*response.Station, errHttp.ErrorHttp)
}

func (sc *stationController) CreateStation(
	c *context.Context,
	req *request.CreateStation,
) (*response.Station, errHttp.ErrorHttp) {
	domSt, derr := sc.stRepo.CreateStation(c, stDom.NewCreateStationFromRequest(req))
	if derr != nil {
		if derr.Conflict() {
			c.Debug("attempted to create a station with duplicate data", zap.Error(derr.Verbose()))
			return nil, errHttp.NewErrConflict(derr, c.LocaleOrDefault())
		}
		c.Error("unexpectedly failed to create a new station", zap.Error(derr.Verbose()))
		return nil, errHttp.NewErrInternal(derr, c.LocaleOrDefault())
	}

	return response.NewStationFromDomain(domSt), nil
}

type stationController struct {
	stRepo stationRepo.StationRepo
}

func NewStationController(stRepo stationRepo.StationRepo) StationController {
	return &stationController{stRepo}
}
