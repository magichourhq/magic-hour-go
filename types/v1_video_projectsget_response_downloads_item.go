package types

// The download url and expiration date of the image project
type V1VideoProjectsgetResponseDownloadsItem struct {
	ExpiresAt string `json:"expires_at"`
	Url       string `json:"url"`
}
