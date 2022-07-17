package dbrepository

import (
	"fampay/domain"

	"google.golang.org/api/youtube/v3"
)

type IRepo interface {
	AddOrUpdateVediosBulk(vedios []*youtube.SearchResult) error
	GetVedioDetailsBySearchTerm(searchTerm string, offset, size int) (*domain.VediosResponse, error)
}
