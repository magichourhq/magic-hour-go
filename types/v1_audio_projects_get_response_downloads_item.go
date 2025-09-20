package types

// The download url and expiration date of the audio project
type V1AudioProjectsGetResponseDownloadsItem struct {
	ExpiresAt string `json:"expires_at"`
	Url       string `json:"url"`
}
