package elasticsearch

type Filtered struct {
    Query interface{}
    Filter interface{}
}

func (f *Filtered) AddQuery(q EsQuery){
    f.Query = q.Source()
}

func (f *Filtered) AddFilter(q EsQuery){
    f.Filter  = q.Source()
}

func (f *Filtered) Source() interface{}{
    m:=makeMap()
    m["query"] = f.Query
    m["filter"] = f.Filter
    v:=makeMap()
    v["filtered"] = m
    return v
}

func NewFilteredQuery() *Filtered{
    return &Filtered{}
}