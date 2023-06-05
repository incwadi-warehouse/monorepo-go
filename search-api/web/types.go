package web

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type SearchBody struct {
	Facets []string `json:"facets"`
	Filter []string `json:"filter"`
	Limit  int      `json:"limit"`
	Q      string   `json:"q"`
}

type IndexBody struct {
	Uid        string `json:"uid"`
	PrimaryKey string `json:"primaryKey"`
}

type DocumentsBody struct {
	Id              string `json:"id"`
	Title           string `json:"title"`
	AuthorFirstname string `json:"authorFirstname"`
	AuthorSurname   string `json:"authorSurname"`
	Genre           string `json:"genre"`
	Price           int    `json:"price"`
	Currency        string `json:"Currency"`
}

type SettingsBody struct {
	DisplayedAttributes  []string                  `json:"displayedAttributes"`
	DistinctAttribute    string                    `json:"distinctAttribute"`
	Faceting             FacetingSettingsBody      `json:"faceting"`
	FilterableAttributes []string                  `json:"filterableAttributes"`
	Pagination           PaginationSettingsBody    `json:"pagination"`
	RankingRules         []string                  `json:"rankingRules"`
	SearchableAttributes []string                  `json:"searchableAttributes"`
	SortableAttributes   []string                  `json:"sortableAttributes"`
	StopWords            []string                  `json:"stopWords"`
	Synonyms             interface{}               `json:"synonyms"`
	TypoTolerance        TypoToleranceSettingsBody `json:"typoTolerance"`
}

type FacetingSettingsBody struct {
	MaxValuesPerFacet int `json:"maxValuesPerFacet"`
}

type PaginationSettingsBody struct {
	MaxTotalHits int `json:"maxTotalHits"`
}

type TypoToleranceSettingsBody struct {
	DisableOnAttributes []string                                     `json:"disableOnAttributes"`
	DisableOnWords      []string                                     `json:"disableOnWords"`
	Enabled             bool                                         `json:"enabled"`
	MinWordSizeForTypos MinWordSizeForTyposTypoToleranceSettingsBody `json:"minWordSizeForTypos"`
}

type MinWordSizeForTyposTypoToleranceSettingsBody struct {
	OneTypo  int `json:"oneTypo"`
	TwoTypos int `json:"twoTypos"`
}
