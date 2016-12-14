package elasticsearch

type MatchQuery struct {
	Field string
	Value interface{}
}

func (mq *MatchQuery) Source() interface{} {
	m := makeMap()
	ma := makeMap()
	m[mq.Field] = mq.Value
	ma["match"] = m
	return ma
}
