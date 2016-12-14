package elasticsearch

type MatchAll struct {
}

func (ma *MatchAll) Source() interface{} {
	m := makeMap()
	empty := NewEmpty()
	m["match_all"] = empty.Source()
	return m
}

func NewMatchAllQuery() *MatchAll {
	return &MatchAll{}
}
