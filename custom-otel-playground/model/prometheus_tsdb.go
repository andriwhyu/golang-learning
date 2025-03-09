package model

type PromTSDBResponse struct {
	Status string   `json:"status"`
	Data   tsdbData `json:"data"`
}

type tsdbData struct {
	HeadStats                   headStats      `json:"headStats"`
	SeriesCountByMetricName     []keyPairValue `json:"seriesCountByMetricName"`
	LabelValueCountByLabelName  []keyPairValue `json:"labelValueCountByLabelName"`
	MemoryInBytesByLabelName    []keyPairValue `json:"memoryInBytesByLabelName"`
	SeriesCountByLabelValuePair []keyPairValue `json:"seriesCountByLabelValuePair"`
}

type headStats struct {
	NumSeries     int   `json:"numSeries"`
	NumLabelPairs int   `json:"numLabelPairs"`
	ChunkCount    int   `json:"chunkCount"`
	MinTime       int64 `json:"minTime"`
	MaxTime       int64 `json:"maxTime"`
}

type keyPairValue struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}
