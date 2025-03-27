package gql

import (
	"encoding/json"
	"fmt"

	"gorm.io/datatypes"
)

type JSON datatypes.JSON

func (JSON) ImplementsGraphQLType(name string) bool {
	return name == "JSON"
}

func (j *JSON) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case JSON:
		*j = input
		return nil
	default:
		return fmt.Errorf("wrong type")
	}
}

func (j JSON) MarshalJSON() ([]byte, error) {
	if j == nil {
		return []byte("null"), nil
	}

	return json.RawMessage(j), nil
}
