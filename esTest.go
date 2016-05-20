package main

import (
	"github.com/astaxie/beego/logs/es"
	"gopkg.in/olivere/elastic.v3"
	"fmt"
	"encoding/json"
)

type EsCfg struct {
	HostPort    []string
	IndexPrefix string
	Type        string
	Timeout     int
	Es          *elastic.Client
}

type DidiPoiInfo struct {
	Poiid            int64   `db:"poiid" json:"poiid"`
	Uid              string  `db:"uid" json:"uid"`
	Stype            int64   `db:"stype" json:"stype"`
	Area             int64   `db:"area" json:"area"`
	City             string  `db:"city" json:"city"`
	District         string  `db:"district" json:"district"`
	Displayname      string  `db:"displayname" json:"displayname"`
	Editname         string  `db:"editname" json:"editname"`
	DisplaynameLower string  `db:"displayname_lower" json:"displayname_lower"`
	DisplaynamePy    string  `db:"displayname_py" json:"displayname_py"`
	DisplaynameM     string  `db:"displayname_m" json:"displayname_m"`
	DisplaynameS     string  `db:"displayname_s" json:"displayname_s"`
	Address          string  `db:"address" json:"address"`
	Lng              float64 `db:"lng" json:"lng"`
	Lat              float64 `db:"lat" json:"lat"`
	Cotype           int64   `db:"cotype" json:"cotype"`
	RawTag           string  `db:"raw_tag" json:"raw_tag"`
	ManualRawTag     string  `db:"manual_raw_tag" json:"manual_raw_tag"`
	TagScore         int64   `db:"tag_score" json:"tag_score"`
	Cscore           float64 `db:"cscore" json:"cscore"`
	RawCscore        int64   `db:"raw_cscore" json:"raw_cscore"`
	Sameid           string  `db:"sameid" json:"sameid"`
	IsDel            int64   `db:"is_del" json:"is_del"`
	Status           int64   `db:"status" json:"status"`
	Updated          string  `db:"updated" json:"updated"`
}


func parse_es_result(r *elastic.SearchResult) () {
	if r.Hits.TotalHits == 0 {
		// 搜索无结果
		return
	}
	if r.Hits == nil || r.Hits.Hits == nil || len(r.Hits.Hits) == 0 {
		return
	}
	// 索引中的总召回数
	fmt.Println(int64(len(r.Hits.Hits)))
	// 在这里直接为es_result的结果数组分配空间。
	for _, hit := range r.Hits.Hits {
		var v DidiPoiInfo
		parse_err := json.Unmarshal(*hit.Source, &v)
		if parse_err != nil {
			fmt.Println("err" + parse_err)
		}

		out, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(out))
	}
	return
}

func main() {

	var es EsCfg = initEs()
	search_client := es.Es.Search().
	Index(es.IndexPrefix).
	Type(es.Type).
	Query("北")

	searchRes, _ := search_client.Do()
	parse_es_result(searchRes)
}