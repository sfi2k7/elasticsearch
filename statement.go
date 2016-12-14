package elasticsearch

import (
	"encoding/json"
)

type Statement struct {
    From int `json:"from"`
    Size int `json:"size"`
    Sorts []interface{}
    Query EsQuery
}

func (s *Statement) AddQuery(query EsQuery){
    s.Query = query
}

func (s *Statement) AddSort(eq EsQuery){
    s.Sorts = append(s.Sorts,eq.Source())
}

func (s *Statement) Compile() []byte{
    m:=makeMap()
    if s.Size == 0{
        s.Size = 50
    }
    
    m["from"]= s.From
    m["size"] = s.Size
    m["query"] = s.Query.Source()
    if len(s.Sorts) > 0{
        m["sort"] = s.Sorts
    }
    
    jsoned,_:=json.Marshal(m)
    return jsoned 
}



