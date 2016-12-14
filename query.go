package elasticsearch

type Query struct {
    q EsQuery
}

func (gq *Query) SetQuery(q EsQuery){
    gq.q = q
}

func (q *Query) Source() interface{}{
    m:=makeMap()
    m["query"] = q.q.Source()
    return m
}

func NewQuery() *Query{
    return &Query{}
}