package meili

import (
	"os"

	"github.com/meilisearch/meilisearch-go"
)

func NewClient() meilisearch.ServiceManager {
	return meilisearch.New(os.Getenv("MEILI"), meilisearch.WithAPIKey(os.Getenv("MEILI_TOKEN")))
}
