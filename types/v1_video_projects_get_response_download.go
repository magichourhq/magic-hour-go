package types

// Deprecated: Please use `.downloads` instead. The download url and expiration date of the video project
type V1VideoProjectsGetResponseDownload struct {
	ExpiresAt string `json:"expires_at"`
	Url       string `json:"url"`
}
