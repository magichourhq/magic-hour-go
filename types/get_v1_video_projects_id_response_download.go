
package types

// Deprecated: Please use `.downloads` instead. The download url and expiration date of the video project
type GetV1VideoProjectsIdResponseDownload struct {
    ExpiresAt string `json:"expires_at"`
    Url string `json:"url"`
}



