package arcgis

import (
	"strconv"
)

type Query struct {
	Limit             int
	ObjectIDs         string
	OutFields         string
	ResultRecordCount uint
	ResultOffset      uint
	SpatialReference  string // Should eventually make an enum, probably
	Where             string
}

func NewQuery() *Query {
	q := new(Query)
	return q
}

func (query Query) toParams() map[string]string {
	params := make(map[string]string)
	if query.Limit > 0 {
		params["limit"] = strconv.Itoa(query.Limit)
	}
	if query.ObjectIDs != "" {
		params["objectIds"] = query.ObjectIDs
	}
	if query.OutFields != "" {
		params["outFields"] = query.OutFields
	}
	if query.ResultOffset > 0 {
		params["resultOffset"] = strconv.Itoa(int(query.ResultOffset))
	}
	if query.Where != "" {
		params["where"] = query.Where
	}
	if len(query.SpatialReference) > 0 {
		params["outSR"] = query.SpatialReference
	}
	return params
}
