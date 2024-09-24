package content

import (
	"time"
)

// contentRoot is the root directory for content files.
const contentRoot = "data/content"

// IndexEntry represents an entry in the content index.
type IndexEntry struct {
	Slug    string    `json:"slug"`
	Date    time.Time `json:"date"`
	Summary string    `json:"summary"`
}
