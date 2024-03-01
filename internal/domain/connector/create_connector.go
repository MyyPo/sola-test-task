package station

import "sola-test-task/internal/domain/connector/share"

type CreateConnector struct {
	ExternalID string
	Type       share.ConnectorType
	MaxPowerKw float64
}

func NewCreateConnector(extId string, typ share.ConnectorType, maxPow float64) CreateConnector {
	return CreateConnector{
		ExternalID: extId,
		Type:       typ,
		MaxPowerKw: maxPow,
	}
}
