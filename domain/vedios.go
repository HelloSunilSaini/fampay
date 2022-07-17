package domain

import "google.golang.org/api/youtube/v3"

type VediosResponse struct {
	Vedios     []youtube.SearchResult `json:"vedios,omitempty"`
	TotalCount int64                  `json:"totalCount,omitempty"`
}
