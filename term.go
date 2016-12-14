package elasticsearch

type TermQuery struct {
	Field string
	Value interface{}
}

func (tq *TermQuery) Source() interface{} {
	m := makeMap()
	v := makeMap()
	t := makeMap()
	v["value"] = tq.Value
	m[tq.Field] = v
	t["term"] = m
	return t
}

func NewTermQuery(field string, value interface{}) *TermQuery {
	return &TermQuery{Field: field, Value: value}
}


