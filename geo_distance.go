package elasticsearch

import (
	"strconv"
)

type GeoDistenceFilterQuery struct {
	Distance  string  
	Latitude  float64 
	Longitude float64 
}

func (g *GeoDistenceFilterQuery) Source() interface{} {
	m := make(map[string]interface{}, 0)
	v := make(map[string]interface{}, 0)
	d := make(map[string]interface{}, 0)
	d["lat"] = g.Latitude
	d["lon"] = g.Longitude
	v["locations.location"] = d
	v["distance"] = g.Distance
	v["distance_type"] = "arc"
	v["optimize_bbox"] = "memory"
	m["geo_distance"] = v
	return m
}

func NewGeoDistanceQuery(lat,lng,distance float64) *GeoDistenceFilterQuery{
	g:=&GeoDistenceFilterQuery{Latitude:lat,Longitude:lng}
	d:= strconv.Itoa(int(distance))+"miles"
	g.Distance = d
	return g
}