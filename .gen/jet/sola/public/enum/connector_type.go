//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package enum

import "github.com/go-jet/jet/v2/postgres"

var ConnectorType = &struct {
	Ccs     postgres.StringExpression
	ChadeMO postgres.StringExpression
	Type1   postgres.StringExpression
	Type2   postgres.StringExpression
}{
	Ccs:     postgres.NewEnumValue("CCS"),
	ChadeMO: postgres.NewEnumValue("CHAdeMO"),
	Type1:   postgres.NewEnumValue("Type1"),
	Type2:   postgres.NewEnumValue("Type2"),
}
