package elasticsearch

type MatchPhrasePrefixQuery struct {
	Field string
	Value interface{}
}

func (t *MatchPhrasePrefixQuery) Source() interface{} {
	m := make(map[string]interface{}, 0)
	v := make(map[string]interface{}, 0)
	v[t.Field] = t.Value
	m["match_phrase_prefix"] = v
	return m
}

func NewMatchPhrasePrefixQuery(field,value string) *MatchPhrasePrefixQuery{
    return &MatchPhrasePrefixQuery{Field:field, Value:value}
}