package response

type GeocodeFindAddressCandidates struct {
	Candidates       []GeocodeCandidate `json:"candidates"`
	SpatialReference SpatialReference   `json:"spatialReference"`
}
type GeocodeCandidate struct {
	Address    string            `json:"address"`
	Location   Location          `json:"location"`
	Score      int               `json:"score"`
	Attributes AttributesGeocode `json:"attributes"`
	Extent     Extent            `json:"extent"`
}
type AttributesGeocode struct {
	// See https://developers.arcgis.com/rest/geocode/service-output/#output-fields
	/*
		SpatialReference
		Address
		Location
		ResultID
		LocName
		Status
		Score
		MatchAddr
		LongLabel
		ShortLabel
		AddrType
		Type
		PlaceName
		PlaceAddr
		Phone
		URL
		Rank
		AddBldg
		AddNum
		AddNumFrom
		AddNumTo
		AddRange
		Side
		StPreDir
		StPreType
		StName
		StType
		StdDir
		BldgComp
		BldgType
		BldgName
		LevelType
		LevelName
		UnitType
		UnitName
		RoomType
		RoomName
		WingType
		WingName
		SubAddr
		StAddr
		Block
		Sector
		Nbrhd
		Neighborhood
		District
		City
		MetroArea
		Subregion
		Region
		RegionAbbr
		Territory
		Postal
		PostalExt
		Country
		CountryCode
		CntryName
		LangCode
		Distance
		X
		Y
		InputX
		InputY
		DisplayX
		DisplayY
		Xmin
		Xmax
		Ymin
		Ymax
		ExInfo
		MatchID
		PotentialID
		StrucType
		StrucDet
		extent
	*/
}
