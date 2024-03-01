package response

import (
	domConn "sola-test-task/internal/domain/connector"
	shareConn "sola-test-task/internal/domain/connector/share"
	domSt "sola-test-task/internal/domain/station"
	shareSt "sola-test-task/internal/domain/station/share"
	"sola-test-task/pkg/util"
)

type Station struct {
	UID        string         `json:"uid"`
	Public     bool           `json:"public"`
	Title      *string        `json:"title"`
	Desc       *string        `json:"description"`
	Address    *string        `json:"address"`
	Coords     *shareSt.Point `json:"coordinates"`
	Connectors []connector    `json:"connectors"`
}

type connector struct {
	ID         string                  `json:"id"`
	Type       shareConn.ConnectorType `json:"type"`
	MaxPowerKw float64                 `json:"max_power_kw"`
}

func NewStationFromDomain(domSt *domSt.Station) *Station {
	return &Station{
		UID:        domSt.ExternalUID,
		Public:     domSt.Public,
		Title:      domSt.Title,
		Desc:       domSt.Desc,
		Address:    domSt.Address,
		Coords:     domSt.Coords,
		Connectors: util.Map(domSt.Connectors, newConnectorFromDomain),
	}
}

func newConnectorFromDomain(domConn domConn.Connector) connector {
	return connector{
		ID:         domConn.ExternalID,
		Type:       domConn.Type,
		MaxPowerKw: domConn.MaxPowerKw,
	}
}
