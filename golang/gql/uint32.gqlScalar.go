package gql

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Uint32 uint32

func (Uint32) ImplementsGraphQLType(name string) bool {
	return name == "Uint32"
}

// UnmarshalGraphQL converts GraphQL input to Uint32
func (u *Uint32) UnmarshalGraphQL(input interface{}) error {
	switch v := input.(type) {
	case int:
		if v < 0 {
			return fmt.Errorf("uint32 cannot be negative")
		}
		*u = Uint32(v)
	case int8:
		if v < 0 {
			return fmt.Errorf("uint32 cannot be negative")
		}
		*u = Uint32(v)
	case int16:
		if v < 0 {
			return fmt.Errorf("uint32 cannot be negative")
		}
		*u = Uint32(v)
	case int32:
		if v < 0 {
			return fmt.Errorf("uint32 cannot be negative")
		}
		*u = Uint32(v)
	case int64:
		if v < 0 || v > int64(^uint32(0)) {
			return fmt.Errorf("value out of range for uint32")
		}
		*u = Uint32(v)
	case uint8:
		*u = Uint32(v)
	case uint16:
		*u = Uint32(v)
	case uint32:
		*u = Uint32(v)
	case uint64:
		if v > uint64(^uint32(0)) {
			return fmt.Errorf("value out of range for uint32")
		}
		*u = Uint32(v)
	case float32:
		if v < 0 || v > float32(^uint32(0)) {
			return fmt.Errorf("value out of range for uint32")
		}
		*u = Uint32(uint32(v))
	case float64:
		if v < 0 || v > float64(^uint32(0)) {
			return fmt.Errorf("value out of range for uint32")
		}
		*u = Uint32(v)
	case json.Number:
		val, err := v.Int64()
		if err != nil || val < 0 || val > int64(^uint32(0)) {
			return fmt.Errorf("invalid uint32 value")
		}
		*u = Uint32(val)
	case string:
		val, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return fmt.Errorf("invalid uint32 value: %s", err)
		}
		*u = Uint32(val)
	default:
		return fmt.Errorf("invalid type for uint32: %T", input)
	}
	return nil
}

// MarshalJSON converts Uint32 to JSON
func (u Uint32) MarshalJSON() ([]byte, error) {
	return json.Marshal(uint32(u))
}
