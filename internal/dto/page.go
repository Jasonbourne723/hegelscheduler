package dto

type PageRequest struct {
	Index int
	Size  int
}

type PageResponse[T any] struct {
	Total int64
	Items []*T
}
