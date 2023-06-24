package update

import (
	"reflect"
	"testing"

	"github.com/meilisearch/meilisearch-go"
)

type MockClient struct {
	IndexFunc       func(uid string) *meilisearch.Index
	GetIndexesFunc  func(*meilisearch.IndexesQuery) (*meilisearch.IndexesResults, error)
	CreateIndexFunc func(config *meilisearch.IndexConfig) (resp *meilisearch.TaskInfo, err error)
	DeleteIndexFunc func(uid string) (resp *meilisearch.TaskInfo, err error)
}

func (m *MockClient) Index(uid string) *meilisearch.Index {
	return m.IndexFunc(uid)
}

func (m *MockClient) GetIndexes(indexes *meilisearch.IndexesQuery) (*meilisearch.IndexesResults, error) {
	return m.GetIndexesFunc(indexes)
}

func (m *MockClient) CreateIndex(config *meilisearch.IndexConfig) (resp *meilisearch.TaskInfo, err error) {
	return m.CreateIndexFunc(config)
}
func (m *MockClient) DeleteIndex(uid string) (resp *meilisearch.TaskInfo, err error) {
	return m.DeleteIndexFunc(uid)
}

func TestUpdateConf_getIndexes(t *testing.T) {
	client = &MockClient{GetIndexesFunc: func(*meilisearch.IndexesQuery) (*meilisearch.IndexesResults, error) {
		indexes := &meilisearch.IndexesResults{
			Results: []meilisearch.Index{{UID: "products_1"}},
			Offset:  0,
			Limit:   0,
			Total:   1,
		}
		return indexes, nil
	}}

	type fields struct {
		AllowedBranches []string
		AllowedIndexes  []string
		Indexes         []string
		ShouldExist     []string
		MustCreate      []string
		MustRemove      []string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			"Indexes should exist",
			fields{
				AllowedIndexes:  []string{"products"},
				AllowedBranches: []string{"1"},
			},
			[]string{"products_1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &UpdateConf{
				AllowedBranches: tt.fields.AllowedBranches,
				AllowedIndexes:  tt.fields.AllowedIndexes,
				Indexes:         tt.fields.Indexes,
				ShouldExist:     tt.fields.ShouldExist,
				MustCreate:      tt.fields.MustCreate,
				MustRemove:      tt.fields.MustRemove,
			}
			conf.getIndexes()
			if got := conf.Indexes; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("conf.Indexes = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateConf_getShouldExist(t *testing.T) {
	type fields struct {
		AllowedBranches []string
		AllowedIndexes  []string
		Indexes         []string
		ShouldExist     []string
		MustCreate      []string
		MustRemove      []string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			"Indexes should exist",
			fields{
				AllowedIndexes:  []string{"products"},
				AllowedBranches: []string{"1"},
			},
			[]string{"products_1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &UpdateConf{
				AllowedBranches: tt.fields.AllowedBranches,
				AllowedIndexes:  tt.fields.AllowedIndexes,
				Indexes:         tt.fields.Indexes,
				ShouldExist:     tt.fields.ShouldExist,
				MustCreate:      tt.fields.MustCreate,
				MustRemove:      tt.fields.MustRemove,
			}
			conf.getShouldExist()
			if got := conf.ShouldExist; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("conf.ShouldExist = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateConf_getMustCreate(t *testing.T) {
	type fields struct {
		AllowedBranches []string
		AllowedIndexes  []string
		Indexes         []string
		ShouldExist     []string
		MustCreate      []string
		MustRemove      []string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			"Indexes should exist",
			fields{
				AllowedIndexes:  []string{"products"},
				AllowedBranches: []string{"1", "3"},
				Indexes:         []string{"products_1"},
				ShouldExist:     []string{"products_1", "products_3"},
			},
			[]string{"products_3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &UpdateConf{
				AllowedBranches: tt.fields.AllowedBranches,
				AllowedIndexes:  tt.fields.AllowedIndexes,
				Indexes:         tt.fields.Indexes,
				ShouldExist:     tt.fields.ShouldExist,
				MustCreate:      tt.fields.MustCreate,
				MustRemove:      tt.fields.MustRemove,
			}
			conf.getMustCreate()
			if got := conf.MustCreate; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("conf.MustCreate = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateConf_getMustRemove(t *testing.T) {
	type fields struct {
		AllowedBranches []string
		AllowedIndexes  []string
		Indexes         []string
		ShouldExist     []string
		MustCreate      []string
		MustRemove      []string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			"Indexes should exist",
			fields{
				AllowedIndexes:  []string{"products"},
				AllowedBranches: []string{"1", "2"},
				Indexes:         []string{"products_1", "products_3"},
				ShouldExist:     []string{"products_1"},
			},
			[]string{"products_3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &UpdateConf{
				AllowedBranches: tt.fields.AllowedBranches,
				AllowedIndexes:  tt.fields.AllowedIndexes,
				Indexes:         tt.fields.Indexes,
				ShouldExist:     tt.fields.ShouldExist,
				MustCreate:      tt.fields.MustCreate,
				MustRemove:      tt.fields.MustRemove,
			}
			conf.getMustRemove()
			if got := conf.MustRemove; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("conf.MustRemove = %v, want %v", got, tt.want)
			}
		})
	}
}
