package gql

import (
	"github.com/BehemothLtd/behemoth-pkg/golang/pagination"
)

type PagyInput struct {
	PerPage *int32
	Page    *int32
}

// ToPaginationInput converts PagyInput to models.PaginationData.
func ToPaginationInput[T any](input *PagyInput) pagination.PaginationData[T] {
	paginationInput := pagination.PaginationData[T]{
		Metadata: pagination.Metadata{
			Page:    1,
			PerPage: 10,
		},
	}

	if input != nil {
		if input.Page != nil && *input.Page >= 1 {
			paginationInput.Metadata.Page = uint32(*input.Page)
		}

		if input.PerPage != nil && *input.PerPage >= 1 {
			paginationInput.Metadata.PerPage = uint32(*input.PerPage)
		}
	}

	return paginationInput
}
