package types

type ChannelMetaInfo struct {
	ChannelVersion int `json:"channel_version"`
	Packages       map[string]struct {
		Subdirs []string `json:"subdirs"`
		Version string   `json:"version"`
	} `json:"packages"`
	Subdirs []string `json:"subdirs"`
}
