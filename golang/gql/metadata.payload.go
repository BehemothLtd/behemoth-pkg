package gql

import (
	"github.com/BehemothLtd/behemoth-pkg/golang/pagination"
)

type MetadataPayload struct {
	Metadata pagination.Metadata
}

func (mt *MetadataPayload) Total() *Uint32 {
	return (*Uint32)(&mt.Metadata.Total)
}

func (mt *MetadataPayload) PerPage() *Uint32 {
	return (*Uint32)(&mt.Metadata.PerPage)
}

func (mt *MetadataPayload) Page() *Uint32 {
	return (*Uint32)(&mt.Metadata.Page)
}

func (mt *MetadataPayload) Pages() *Uint32 {
	return (*Uint32)(&mt.Metadata.Pages)
}

func (mt *MetadataPayload) Count() *Uint32 {
	return (*Uint32)(&mt.Metadata.Count)
}

func (mt *MetadataPayload) Next() *Uint32 {
	return (*Uint32)(&mt.Metadata.Next)
}

func (mt *MetadataPayload) Prev() *Uint32 {
	return (*Uint32)(&mt.Metadata.Prev)
}

func (mt *MetadataPayload) From() *Uint32 {
	return (*Uint32)(&mt.Metadata.From)
}

func (mt *MetadataPayload) To() *Uint32 {
	return (*Uint32)(&mt.Metadata.To)
}
