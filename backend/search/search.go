package search

import (
	"context"
	"github.com/anda-ai/anda/models"
)

// Searcher is the interface for searching google bing or other search outer search engine ....
type Searcher interface {
	Search(ctx context.Context, query string, pageSize int) (*models.SearchResult, error)
}
