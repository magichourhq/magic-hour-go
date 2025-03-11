package types

// The download url and expiration date of the image project
type GetV1ImageProjectsIdResponseDownloadsItem struct {
	ExpiresAt string `json:"expires_at"`
	Url       string `json:"url"`
}
