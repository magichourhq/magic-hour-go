package types

// In the case of an error, this object will contain the error encountered during video render
type V1VideoProjectsgetResponseError struct {
	// An error code to indicate why a failure happened.
	Code string `json:"code"`
	// Details on the reason why a failure happened.
	Message string `json:"message"`
}
