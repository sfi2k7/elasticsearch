package elasticsearch

func makeMap() map[string]interface{} {
	return make(map[string]interface{})
}

func NewStatement() *Statement {
	return &Statement{}
}

func MatchAllStatement() *Statement {
	s := NewStatement()
	m := NewMatchAllQuery()
	s.AddQuery(m)
	return s
}
