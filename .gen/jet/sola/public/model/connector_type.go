//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import "errors"

type ConnectorType string

const (
	ConnectorType_Ccs     ConnectorType = "CCS"
	ConnectorType_ChadeMO ConnectorType = "CHAdeMO"
	ConnectorType_Type1   ConnectorType = "Type1"
	ConnectorType_Type2   ConnectorType = "Type2"
)

func (e *ConnectorType) Scan(value interface{}) error {
	var enumValue string
	switch val := value.(type) {
	case string:
		enumValue = val
	case []byte:
		enumValue = string(val)
	default:
		return errors.New("jet: Invalid scan value for AllTypesEnum enum. Enum value has to be of type string or []byte")
	}

	switch enumValue {
	case "CCS":
		*e = ConnectorType_Ccs
	case "CHAdeMO":
		*e = ConnectorType_ChadeMO
	case "Type1":
		*e = ConnectorType_Type1
	case "Type2":
		*e = ConnectorType_Type2
	default:
		return errors.New("jet: Invalid scan value '" + enumValue + "' for ConnectorType enum")
	}

	return nil
}

func (e ConnectorType) String() string {
	return string(e)
}