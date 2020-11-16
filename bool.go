package elasticsearch

import "fmt"

type BoolQuery struct {
	Shoulds []interface{}
	Musts   []interface{}
}

func (bq *BoolQuery) AddShoulds(qs ...EsQuery) {
	for _, q := range qs {
		bq.Shoulds = append(bq.Shoulds, q.Source())
	}
}

func (bq *BoolQuery) AddMusts(qs ...EsQuery) {
	fmt.Println("LEN MUST", len(qs))
	for _, q := range qs {
		bq.Musts = append(bq.Musts, q.Source())
	}
}

func (bq *BoolQuery) Source() interface{} {
	b := makeMap()
	m := makeMap()
	if len(bq.Shoulds) > 0 {
		m["should"] = bq.Shoulds
	}
	if len(bq.Musts) > 0 {
		m["must"] = bq.Musts
	}
	b["bool"] = m
	return b
}

func NewBoolQuery() *BoolQuery {
	return &BoolQuery{}
}
