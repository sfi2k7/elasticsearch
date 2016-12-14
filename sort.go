package elasticsearch

type SortByField struct {
	fieldName string
	Order     string
}

func (sbf *SortByField) Source() interface{} {
	m := makeMap()
	if len(sbf.Order) == 0 {
		sbf.Order = "asc"
	}

	m[sbf.fieldName] = sbf.Order
	return m
}

func NewSortBy(field string, order string) *SortByField {
	return &SortByField{fieldName: field, Order: order}
}
