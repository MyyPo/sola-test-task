package station

import "sola-test-task/internal/domain/connector/share"

type Connector struct {
	ExternalID string
	Type       share.ConnectorType
	MaxPowerKw float64
}

func NewConnector(extId string, typ share.ConnectorType, maxPow float64) Connector {
	return Connector{
		ExternalID: extId,
		Type:       typ,
		MaxPowerKw: maxPow,
	}
}
