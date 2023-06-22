package update

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"

	"github.com/incwadi-warehouse/monorepo-go/search-api/api"
	"github.com/incwadi-warehouse/monorepo-go/search-api/util"
)

type UpdateConf struct {
	AllowedBranches []string
	AllowedIndexes  []string
	Indexes         []string
	ShouldExist     []string
	MustCreate      []string
	MustRemove      []string
}

type Indexes struct {
	Limit   int     `json:"limit"`
	Offset  int     `json:"offset"`
	Results []Index `json:"results"`
	Total   int     `json:"total"`
}

type Index struct {
	Uid        string `json:"uid"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
	PrimaryKey string `json:"primaryKey"`
}

type CreateIndex struct {
	Uid        string `json:"uid"`
	PrimaryKey string `json:"primaryKey"`
}

var request = api.NewRequestWithPlainRes

func Run() {
	conf := &UpdateConf{
		AllowedBranches: strings.Split(os.Getenv("BRANCHES"), ","),
		AllowedIndexes:  strings.Split(os.Getenv("INDEXES"), ","),
	}

	conf.getIndexes()
	conf.getShouldExist()
	conf.getMustCreate()
	conf.getMustRemove()

	conf.doCreate()
	conf.doRemove()
}

func (conf *UpdateConf) getIndexes() {
	res := request("GET", "/indexes", strings.NewReader(""))

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data Indexes
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	for _, v := range data.Results {
		conf.Indexes = append(conf.Indexes, v.Uid)
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
		jsonData, err := json.Marshal(CreateIndex{name, "id"})
		if err != nil {
			log.Fatal(err)
		}
		request("POST", "/indexes", strings.NewReader(string(jsonData)))
	}
}

func (conf *UpdateConf) doRemove() {
	for _, name := range conf.MustRemove {
		request("DELETE", "/indexes/"+name, strings.NewReader(""))
	}
}
