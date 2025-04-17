package pagination

import (
	"context"
	"fmt"
	"math"

	"github.com/BehemothLtd/behemoth-pkg/golang/exceptions"
	translator "github.com/BehemothLtd/behemoth-pkg/golang/translators"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type Metadata struct {
	Total   uint32
	PerPage uint32
	Page    uint32
	Pages   uint32
	Count   uint32
	Next    uint32
	Prev    uint32
	From    uint32
	To      uint32
}

type PaginationData[T any] struct {
	Metadata   Metadata
	Collection T
}

func Paginate[T any](ctx context.Context, db *gorm.DB, p *PaginationData[T], tableName string) (func(db *gorm.DB) *gorm.DB, error) {
	log.Debug().Ctx(ctx).Msg("pagination.Paginate")

	// Create a new session to avoid side effects
	dbClone := db.Session(&gorm.Session{})

	// Calculate the total number of records
	var totalRecords uint32
	if err := dbClone.Select(fmt.Sprintf("COUNT(DISTINCT %s.id)", tableName)).Scan(&totalRecords).Error; err != nil {
		return nil, err
	}

	// Update pagination metadata
	p.Metadata.Total = totalRecords
	p.Metadata.Pages = uint32(math.Max(1, math.Ceil(float64(totalRecords)/float64(p.Metadata.PerPage))))

	// Adjust current page if it exceeds the last page
	if p.Metadata.Page > p.Metadata.Pages {
		return nil, exceptions.NewUnprocessableContentError(translator.Translate(nil, "errValidation_pagingOverflow"), nil)
	}

	// Update navigation metadata
	if p.Metadata.Page > 1 {
		p.Metadata.Prev = p.Metadata.Page - 1
	}
	if p.Metadata.Page < p.Metadata.Pages {
		p.Metadata.Next = p.Metadata.Page + 1
	}

	// Calibrate the count for the current page
	p.Metadata.Count = p.Metadata.PerPage
	if p.Metadata.Page == p.Metadata.Pages {
		p.Metadata.Count = uint32(totalRecords) - (p.Metadata.PerPage * (p.Metadata.Page - 1))
	}

	// Calculate offset and range
	offset := p.Metadata.PerPage * (p.Metadata.Page - 1)
	p.Metadata.From, p.Metadata.To = calculateRange(offset, p.Metadata.Count, totalRecords)

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(int(offset)).Limit(int(p.Metadata.PerPage))
	}, nil
}

// Helper function to calculate the range of items for the pagination metadata
func calculateRange(offset, count uint32, totalRecords uint32) (uint32, uint32) {
	if totalRecords == 0 {
		return 0, 0
	}
	from := offset + 1
	to := uint32(math.Min(float64(offset+count), float64(totalRecords)))
	return from, to
}
