package station

import (
	domConn "sola-test-task/internal/domain/connector"
	"sola-test-task/internal/domain/station/share"
	request "sola-test-task/internal/dto/request/station"
	"sola-test-task/pkg/util"
)

type CreateStation struct {
	ExternalUID string
	Public      bool
	Title       *string
	Desc        *string
	Address     *string
	Coords      *share.Point
	Connectors  []domConn.Connector
}

func NewCreateStationFromRequest(req *request.CreateStation) *CreateStation {
	domConns := make([]domConn.Connector, 0, len(req.Connectors))
	for _, c := range req.Connectors {
		domConns = append(domConns, domConn.NewConnector(c.ID, c.Type, *c.MaxPowerKw))
	}
	return newCreateStation(
		req.UID,
		util.DerefOrDefault(req.Public),
		req.Title,
		req.Desc,
		req.Address,
		req.Coords,
		domConns,
	)
}

func newCreateStation(
	extUid string,
	pub bool,
	title, desc, addr *string,
	coords *share.Point,
	conns []domConn.Connector,
) *CreateStation {
	return &CreateStation{
		ExternalUID: extUid,
		Public:      pub,
		Title:       title,
		Desc:        desc,
		Address:     addr,
		Coords:      coords,
		Connectors:  conns,
	}
}
