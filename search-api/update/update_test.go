package update

import (
	"reflect"
	"testing"
    "io"
    "net/http"
    "strings"
    "encoding/json"
)

func TestUpdateConf_getIndexes(t *testing.T) {
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
			[]string{"product_1"},
		},
	}
	for _, tt := range tests {
        jsonData,_ := json.Marshal(Indexes{
            Results: []Index{{Uid: "products_1"}},
        })
        buf := io.NopCloser(strings.NewReader(string(jsonData)))
        request = func(method, path string, requestBody io.Reader) *http.Response {
            return &http.Response{Body: buf}
        }

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
            if got := conf.Indexes; reflect.DeepEqual(got, tt.want) {
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
			[]string{"product_1"},
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
			if got := conf.ShouldExist; reflect.DeepEqual(got, tt.want) {
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
				AllowedBranches: []string{"1"},
			},
			[]string{"product_1"},
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
			if got := conf.MustCreate; reflect.DeepEqual(got, tt.want) {
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
				AllowedBranches: []string{"1"},
			},
			[]string{"product_1"},
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
			if got := conf.MustRemove; reflect.DeepEqual(got, tt.want) {
				t.Errorf("conf.MustRemove = %v, want %v", got, tt.want)
			}
		})
	}
}
