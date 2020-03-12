package models

type Source struct {
	ID         uint   `json:"id"`
	VersionID  int    `json:"versionId"`
	ProviderID int    `json:"providerId"`
	Url        string `json:"url"`
	Secret     string `json:"secret"`
}
