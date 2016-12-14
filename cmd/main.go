package main

import (
	"fmt"

	"github.com/sfi2k7/elasticsearch"
)

func main() {
	//   s:=elasticsearch.NewStatement()
	//   b:=elasticsearch.NewBoolQuery()
	//   b.AddShould(elasticsearch.NewTermQuery("serving_locations.zip", "78211"))
	//   b.AddShould(elasticsearch.NewGeoDistanceQuery(29.3425345, -98.5700519, 50))
	//   q:=elasticsearch.NewQuery()
	//   q.SetQuery(b)

	//   f:=elasticsearch.NewFilteredQuery()
	//   f.AddFilter(q)

	//   b2:=elasticsearch.NewBoolQuery()
	//   b2.AddMust(elasticsearch.NewTermQuery("types", 5025))
	//   f.AddQuery(b2)
	//   s.AddQuery(f)

	// statement:=elasticsearch.NewStatement()
	// b:=elasticsearch.NewBoolQuery()
	// t1:=elasticsearch.NewTermQuery("types", 307)
	// t2:=elasticsearch.NewTermQuery("serving_locations.zip", "78211")
	// b.AddShoulds(t1,t2)

	// statement.AddQuery(b)
	// fmt.Println(string(statement.Compile()))

	c := elasticsearch.NewClient("http://bmysql:9200/")
	// s := elasticsearch.NewStatement()
	// t := elasticsearch.NewTermQuery("types", 1594)
	// s.AddQuery(t)
	//s := elasticsearch.MathAllStatement()
	s := elasticsearch.NewStatement()
	f := elasticsearch.NewFilteredQuery()
	f.AddFilter(elasticsearch.NewTermQuery("serving_locations.zip", "78211"))
	f.AddQuery(elasticsearch.NewTermQuery("types", 307))
	s.AddQuery(f)
	body, err := c.Search(s, "idx0825/org")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

// POST idx0825/org/_search
// {"from":0,"query":{"filtered":{"filter":{"term":{"serving_locations.zip":{"value":"78211"}}},"query":{"term":{"types":{"value":307}}}}},"size":50}
