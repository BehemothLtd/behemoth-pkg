package utils

import gqlScalar "behemoth-pkg/golang/gqlScalars"

type AllowedTypes interface {
	int | int32 | uint32 | float64 | string | gqlScalar.Uint32
}

// returns a pointer to the given value.
func ToPointer[T AllowedTypes](v T) *T {
	return &v
}

func PointerSlice(input *[]*string) *[]string {
	if input == nil {
		return nil
	}
	result := make([]string, len(*input))
	for i, s := range *input {
		if s != nil {
			result[i] = *s
		}
	}
	return &result
}

func PtrToValue[T any](ptr *T, defaultValue T) T {
	if ptr != nil {
		return *ptr
	}
	return defaultValue
}
