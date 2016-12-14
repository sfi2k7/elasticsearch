package elasticsearch

type GeoSort struct {
	FieldName string
	Latitude  float64
	Longitude float64
	Order     string
}

func (gs *GeoSort) Source() interface{} {
	b := makeMap()
	s := makeMap()
	l := makeMap()
	l["lat"] = gs.Latitude
	l["lon"] = gs.Longitude
	s[gs.FieldName] = l
	s["distance_type"] = "arc"

	if len(gs.Order) == 0 {
		gs.Order = "asc"
	}

	s["order"] = gs.Order
	s["unit"] = "miles"
	b["_geo_distance"] = s
	return b
}

func NewGeoDistanceSort(fieldName string, lat, lng float64, order string) *GeoSort {
	return &GeoSort{
		FieldName: fieldName, Latitude: lat, Longitude: lng, Order: order,
	}
}
