package dto

type ApiMetaInfo struct {
	Indexer    string `json:"indexer"`
	Image      string `json:"image"`
	Registry   string `json:"registry"`
	Repository string `json:"repository"`
}
