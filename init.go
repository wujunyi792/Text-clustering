package clustering

type ClusterEngine struct {
	BadWords               []string
	ReBadWords             []Restruct
	LegitimateMatchingRate float64
}

type Restruct struct {
	Start     string
	End       string
	MaxLength int
}
