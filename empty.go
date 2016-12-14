package elasticsearch

type Empty struct {
}

func (e *Empty) Source() interface{} {
	return makeMap()
}

func NewEmpty() *Empty {
	return &Empty{}
}
