package condatypes

type ChannelData struct {
	ChannelDataVersion int                    `json:"channeldata_version"`
	Packages           map[string]PackageData `json:"packages"`
	Subdirs            []string               `json:"subdirs"`
}

type PackageData struct {
	Subdirs      []string `json:"subdirs"`
	Version      *string  `json:"version"`
	ActivateD    bool     `json:"activate.d"`
	BinaryPrefix bool     `json:"binary_prefix"`
	DeactivateD  bool     `json:"deactivate.d"`
	Description  *string  `json:"description"`
	DevUrl       *string  `json:"dev_url"`
	DocSourceUrl *string  `json:"doc_source_url"`
	DocUrl       *string  `json:"doc_url"`
	Home         *string  `json:"home"`
	IconHash     *string  `json:"icon_hash"`
	IconUrl      *string  `json:"icon_url"`
	Identifiers  *string  `json:"identifiers"`
	Keywords     []string `json:"keywords"`
	License      *string  `json:"license"`
	PostLink     bool     `json:"post_link"`
	PreLink      bool     `json:"pre_link"`
	PreUnlink    bool     `json:"pre_unlink"`
	RecipeOrigin *string  `json:"recipe_origin"`
	SourceGitUrl *string  `json:"source_git_url"`
	SourceUrl    *string  `json:"source_url"`
	Summary      *string  `json:"summary"`
	Tags         []string `json:"tags"`
	TextPrefix   bool     `json:"text_prefix"`
	Timestamp    uint64   `json:"timestamp"`
}
