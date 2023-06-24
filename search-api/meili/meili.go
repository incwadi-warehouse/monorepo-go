package meili

import (
	"os"

	"github.com/meilisearch/meilisearch-go"
)

func NewClient() *meilisearch.Client {
	return meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   os.Getenv("MEILI"),
		APIKey: os.Getenv("MEILI_TOKEN"),
	})
}
