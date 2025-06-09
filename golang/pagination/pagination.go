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

type PaginationData struct {
	Metadata   Metadata
	Collection any
}

func Paginate(ctx context.Context, db *gorm.DB, p *PaginationData, tableName string) (func(db *gorm.DB) *gorm.DB, error) {
	log.Debug().Ctx(ctx).Msg("pagination.Paginate")

	// Create a new session to avoid side effects
	dbClone := db.Session(&gorm.Session{})

	// Calculate the total number of records
	var totalRecords uint32
	dbClone.Select(fmt.Sprintf("COUNT(DISTINCT %s.id)", tableName)).Scan(&totalRecords)

	// Update pagination metadata using the new method
	if err := p.Metadata.Calculate(totalRecords); err != nil {
		return nil, err
	}

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(int(p.Metadata.Offset())).Limit(int(p.Metadata.PerPage))
	}, nil
}

func (meta *Metadata) Calculate(totalRecords uint32) error {
	meta.Total = totalRecords
	meta.Pages = uint32(math.Max(1, math.Ceil(float64(totalRecords)/float64(meta.PerPage))))

	// Adjust current page if it exceeds the last page
	if meta.Page > meta.Pages {
		return exceptions.NewUnprocessableContentError(translator.Translate(nil, "errValidation_pagingOverflow"), nil)
	}

	// Update navigation metadata
	meta.Prev = 0
	meta.Next = 0
	if meta.Page > 1 {
		meta.Prev = meta.Page - 1
	}
	if meta.Page < meta.Pages {
		meta.Next = meta.Page + 1
	}

	// Calibrate the count for the current page
	meta.Count = meta.PerPage
	if meta.Page == meta.Pages {
		meta.Count = uint32(totalRecords) - (meta.PerPage * (meta.Page - 1))
	}

	// Calculate offset and range
	meta.From, meta.To = calculateRange(meta.Offset(), meta.Count, totalRecords)

	return nil
}

func (meta *Metadata) Offset() uint32 {
	if meta.Page == 0 {
		return 0
	}
	return meta.PerPage * (meta.Page - 1)
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
