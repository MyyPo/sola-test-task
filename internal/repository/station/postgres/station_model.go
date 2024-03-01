package postgres

import (
	"database/sql"
	"sola-test-task/pkg/context"
	"sola-test-task/pkg/util"

	model "sola-test-task/.gen/jet/sola/public/model"

	derr "sola-test-task/internal/error/data"

	commJet "sola-test-task/internal/repository/common/postgres_jet"

	domConn "sola-test-task/internal/domain/connector"
	shareConn "sola-test-task/internal/domain/connector/share"
	domSt "sola-test-task/internal/domain/station"
	shareSt "sola-test-task/internal/domain/station/share"

	"github.com/google/uuid"
)

func (r *StationPostgresRepo) newStationErr(
	c *context.Context,
	err error,
	expTypes ...derr.DataErrorType,
) derr.DataError {
	errType, constrName := commJet.ErrSpec(err)
	var fieldName string
	if constrName == "stations_external_uid_key" {
		fieldName = "uid"
	}

	return derr.NewErr(errType, err, "station", fieldName, c.LocaleOrDefault(), expTypes...)
}

func (r *StationPostgresRepo) newConnectorErr(
	c *context.Context,
	err error,
	expTypes ...derr.DataErrorType,
) derr.DataError {
	errType, _ := commJet.ErrSpec(err)
	return derr.NewErr(errType, err, "connector", "", c.LocaleOrDefault(), expTypes...)
}

func NewStationPostgresRepo(conn *sql.DB) *StationPostgresRepo {
	return &StationPostgresRepo{conn}
}

func newStationFromDomain(domSt *domSt.CreateStation) model.Stations {
	var lat *float64
	var long *float64
	if domSt.Coords != nil {
		lat = util.Pointer(domSt.Coords[0])
		long = util.Pointer(domSt.Coords[1])
	}

	return model.Stations{
		ExternalUID: &domSt.ExternalUID,
		Public:      domSt.Public,
		Title:       domSt.Title,
		Description: domSt.Desc,
		Address:     domSt.Address,
		Latitude:    lat,
		Longitude:   long,
	}
}

func newConnectorFromDomainClosure(stId uuid.UUID) func(domConn.Connector) model.Connectors {
	return func(domConn domConn.Connector) model.Connectors {
		return model.Connectors{
			ID:         uuid.New(),
			ExternalID: &domConn.ExternalID,
			StationID:  stId,
			Type:       model.ConnectorType(domConn.Type),
			MaxPowerKw: domConn.MaxPowerKw,
		}
	}
}

func newStationFromModel(
	modSt model.Stations,
	modConns []model.Connectors,
) *domSt.Station {
	var coords *shareSt.Point
	if modSt.Latitude != nil && modSt.Longitude != nil {
		coords = util.Pointer(shareSt.Point{*modSt.Latitude, *modSt.Longitude})
	}

	return &domSt.Station{
		ExternalUID: *modSt.ExternalUID,
		Public:      modSt.Public,
		Title:       modSt.Title,
		Desc:        modSt.Description,
		Address:     modSt.Address,
		Coords:      coords,
		Connectors:  util.Map(modConns, newConnectorFromModel),
	}
}

func newConnectorFromModel(modConn model.Connectors) domConn.Connector {
	return domConn.Connector{
		ExternalID: *modConn.ExternalID,
		Type:       shareConn.ConnectorType(modConn.Type),
		MaxPowerKw: modConn.MaxPowerKw,
	}
}
