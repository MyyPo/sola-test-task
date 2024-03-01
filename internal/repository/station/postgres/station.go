package postgres

import (
	"database/sql"
	domain "sola-test-task/internal/domain/station"
	derr "sola-test-task/internal/error/data"
	"sola-test-task/pkg/context"
	"sola-test-task/pkg/util"

	. "sola-test-task/.gen/jet/sola/public/table"

	"go.uber.org/zap"

	_ "github.com/lib/pq"
)

type StationPostgresRepo struct {
	conn *sql.DB
}

func (r *StationPostgresRepo) CreateStation(
	c *context.Context,
	newSt *domain.CreateStation,
) (*domain.Station, derr.DataError) {
	tx, err := r.conn.Begin()
	if err != nil {
		c.Error("unexpectedly failed to start a new db transaction", zap.Error(err))
		tx.Rollback()
		return nil, r.newStationErr(c, err)
	}

	modSt := newStationFromDomain(newSt)
	stStmt := Stations.INSERT(Stations.ExternalUID, Stations.Public, Stations.Title, Stations.Description, Stations.Address, Stations.Latitude, Stations.Longitude).
		MODEL(modSt).
		RETURNING(Stations.AllColumns)
	if err = stStmt.Query(tx, &modSt); err != nil {
		c.Debug("failed to insert a new station record into db", zap.Error(err))
		tx.Rollback()
		return nil, r.newStationErr(c, err, derr.Conflict)
	}

	modConns := util.Map(newSt.Connectors, newConnectorFromDomainClosure(modSt.ID))
	connsStmt := Connectors.INSERT().MODELS(modConns).RETURNING(Connectors.AllColumns)
	if err = connsStmt.Query(tx, &modConns); err != nil {
		c.Debug("failed to insert new connectors into db", zap.Error(err))
		tx.Rollback()
		return nil, r.newConnectorErr(c, err, derr.Conflict)
	}

	if err := tx.Commit(); err != nil {
		c.Debug("failed to commit the station creation transaction", zap.Error(err))
		return nil, r.newStationErr(c, err)
	}

	return newStationFromModel(modSt, modConns), nil
}
