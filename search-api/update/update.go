package update

import (
	"log"
	"os"
	"strings"

	"github.com/incwadi-warehouse/monorepo-go/search-api/meili"
	"github.com/incwadi-warehouse/monorepo-go/search-api/util"
	"github.com/meilisearch/meilisearch-go"
)

type UpdateConf struct {
	AllowedBranches []string
	AllowedIndexes  []string
	Indexes         []string
	ShouldExist     []string
	MustCreate      []string
	MustRemove      []string
}

type Client interface {
    Index(uid string) *meilisearch.Index
	GetIndexes(*meilisearch.IndexesQuery) (*meilisearch.IndexesResults, error)
    CreateIndex(config *meilisearch.IndexConfig) (resp *meilisearch.TaskInfo, err error)
    DeleteIndex(uid string) (resp *meilisearch.TaskInfo, err error)
}

var client Client

func Run() {
	client = meili.NewClient()

	conf := &UpdateConf{
		AllowedBranches: strings.Split(os.Getenv("BRANCHES"), ","),
		AllowedIndexes:  strings.Split(os.Getenv("INDEXES"), ","),
	}

	conf.getIndexes()
	conf.getShouldExist()
	conf.getMustCreate()
	conf.getMustRemove()

	conf.doCreate()
	conf.doSettings()

	conf.doRemove()
}

func (conf *UpdateConf) getIndexes() {
	res, err := client.GetIndexes(nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range res.Results {
		conf.Indexes = append(conf.Indexes, v.UID)
	}
}

func (conf *UpdateConf) getShouldExist() {
	for _, name := range conf.AllowedIndexes {
		for _, branchId := range conf.AllowedBranches {
			conf.ShouldExist = append(conf.ShouldExist, name+"_"+branchId)
		}
	}
}

func (conf *UpdateConf) getMustCreate() {
	for _, name := range conf.ShouldExist {
		if !util.Contains(name, conf.Indexes) {
			conf.MustCreate = append(conf.MustCreate, name)
		}
	}
}

func (conf *UpdateConf) getMustRemove() {
	for _, name := range conf.Indexes {
		if !util.Contains(name, conf.ShouldExist) {
			conf.MustRemove = append(conf.MustRemove, name)
		}
	}
}

func (conf *UpdateConf) doCreate() {
	for _, name := range conf.MustCreate {
		client.CreateIndex(&meilisearch.IndexConfig{
			Uid:        name,
			PrimaryKey: "id",
		})
	}
}

func (conf *UpdateConf) doRemove() {
	for _, name := range conf.MustRemove {
		client.DeleteIndex(name)
	}
}

func (conf *UpdateConf) doSettings() {
	for _, name := range conf.Indexes {
		filterableAttributes := []string{
			"genre",
		}
		client.Index(name).UpdateFilterableAttributes(&filterableAttributes)
	}
}
