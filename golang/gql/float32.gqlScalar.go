package gql

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"

	"github.com/rs/zerolog/log"
)

type Float32 float32

func (Float32) ImplementsGraphQLType(name string) bool {
	return name == "Float32"
}

// UnmarshalGraphQL converts GraphQL input to Float32
func (f *Float32) UnmarshalGraphQL(input interface{}) error {
	switch v := input.(type) {
	case float32:
		*f = Float32(v)
	case float64:
		if v < -math.MaxFloat32 || v > math.MaxFloat32 {
			return fmt.Errorf("value out of range for float32")
		}
		*f = Float32(v)
	case int:
		*f = Float32(v)
	case int8:
		*f = Float32(v)
	case int16:
		*f = Float32(v)
	case int32:
		*f = Float32(v)
	case int64:
		if float64(v) < -math.MaxFloat32 || float64(v) > math.MaxFloat32 {
			return fmt.Errorf("value out of range for float32")
		}

		if v < -16777216 || v > 16777216 {
			log.Warn().Int64("value", v).Msg("Warning: value may lose precision in float32")
		}

		*f = Float32(v)
	case uint8:
		*f = Float32(v)
	case uint16:
		*f = Float32(v)
	case uint32:
		*f = Float32(v)
	case uint64:
		if float64(v) > float64(math.MaxFloat32) {
			return fmt.Errorf("value out of range for float32")
		}

		if v > 16777216 {
			log.Warn().Uint64("value", v).Msg("Warning: value may lose precision in float32")
		}

		*f = Float32(v)
	case json.Number:
		val, err := v.Float64()
		if err != nil {
			return fmt.Errorf("invalid float32 value: %s", err)
		}
		if val < -math.MaxFloat32 || val > math.MaxFloat32 {
			return fmt.Errorf("value out of range for float32")
		}
		*f = Float32(val)
	case string:
		val, err := strconv.ParseFloat(v, 32)
		if err != nil {
			return fmt.Errorf("invalid float32 value: %s", err)
		}
		*f = Float32(val)
	default:
		return fmt.Errorf("invalid type for float32: %T", input)
	}
	return nil
}

// MarshalJSON converts Float32 to JSON
func (f Float32) MarshalJSON() ([]byte, error) {
	return json.Marshal(float32(f))
}
