package station

import (
	domConn "sola-test-task/internal/domain/connector"
	shareConn "sola-test-task/internal/domain/connector/share"
	"sola-test-task/internal/domain/station/share"
)

type Station struct {
	ExternalUID string
	Public      bool
	Title       *string
	Desc        *string
	Address     *string
	Coords      *share.Point
	Connectors  []domConn.Connector
}

func NewStation(
	extUid string,
	pub bool,
	title, desc, addr *string,
	coords *share.Point,
	conns []domConn.Connector,
) *Station {
	return &Station{
		ExternalUID: extUid,
		Public:      pub,
		Title:       title,
		Desc:        desc,
		Address:     addr,
		Coords:      coords,
		Connectors:  conns,
	}
}

func NewConnector(extId string, typ shareConn.ConnectorType, maxPow float64) domConn.Connector {
	return domConn.Connector{
		ExternalID: extId,
		Type:       typ,
		MaxPowerKw: maxPow,
	}
}
