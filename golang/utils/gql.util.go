package utils

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/graph-gophers/graphql-go"
)

// GetStringOrDefault returns the value of the string pointer if not nil, otherwise returns "".
func GetStringOrDefault(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}

func GqlIdToInt32(i graphql.ID) (int32, error) {
	r, err := strconv.ParseInt(string(i), 10, 32)

	if err != nil {
		return 0, errors.New("GqlIDToUint")
	}

	return int32(r), nil
}

func GqlIDToUint32(i graphql.ID) (uint32, error) {
	r, err := strconv.ParseUint(string(i), 10, 32)
	if err != nil {
		return 0, errors.New("GqlIDToUint32")
	}

	return uint32(r), nil
}

func GqlTimePointer(val *time.Time) *graphql.Time {
	if val != nil {
		time := graphql.Time{Time: *val}

		return &time
	} else {
		return nil
	}
}

func GetInt32OrDefault(num *int32) int32 {
	if num == nil {
		return 0
	}
	return *num
}

func GetGqlUint32OrDefault(num *graphql.ID) uint32 {
	if num == nil {
		return 0
	}

	id, err := GqlIDToUint32(*num)
	if err != nil {
		return 0
	}

	return id
}

func GetFloat64OrDefault(num *float64) float64 {
	if num == nil {
		return 0.0
	}
	return *num
}

func GqlIdToUint32(id graphql.ID) (uint32, error) {
	if id == "" {
		return 0, errors.New("Invalid Id")
	}

	parsedID, err := strconv.ParseUint(string(id), 10, 32)
	if err != nil {
		return 0, err
	}

	return uint32(parsedID), nil
}

func Int32ToUint32Pointer(ptr *int32) uint32 {
	if ptr == nil {
		return 0
	}
	return uint32(*ptr)
}

// GetStringOrDefault returns the value of the bool pointer if not nil, otherwise returns false.
func GetBoolOrDefault(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

type GraphqlIDConvertible interface {
	int32 | uint32 | string
}

func ParseGraphqlID[T GraphqlIDConvertible](id graphql.ID) (T, error) {
	idString := string(id)
	var result T

	switch any(result).(type) {
	case int32:
		parsed, err := strconv.ParseInt(idString, 10, 32)
		if err != nil {
			return result, err
		}
		return T(parsed), nil
	case uint32:
		parsed, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			return result, err
		}
		return T(parsed), nil
	case string:
		return any(idString).(T), nil
	default:
		return result, fmt.Errorf("unsupported type")
	}
}
