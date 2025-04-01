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
	case map[string]interface{}, []interface{}:
		data, err := json.Marshal(input)
		if err != nil {
			return err
		}
		*j = JSON(data)
		return nil
	case string:
		*j = JSON([]byte(input))
		return nil
	case nil:
		*j = nil
		return nil
	default:
		return fmt.Errorf("invalid JSON input type: %T", input)
	}
}

func (j JSON) MarshalJSON() ([]byte, error) {
	if j == nil {
		return []byte("null"), nil
	}

	return json.RawMessage(j), nil
}
