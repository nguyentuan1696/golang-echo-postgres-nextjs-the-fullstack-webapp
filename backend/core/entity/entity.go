package entity

type Pagination[T any] struct {
	Items       []T
	TotalItems  int
	TotalPages  int
	CurrentPage int
	PageSize    int
}
