package core

import (
	fmt "fmt"
	io "io"
	http "net/http"
)

type ApiError struct {
	StatusCode int
	Method     string
	Url        string
	Data       []byte
	Request    http.Request
	Response   http.Response
}

func (e ApiError) Error() string {
	return fmt.Sprintf("Unexpected status code received %d from %s %s", e.StatusCode, e.Method, e.Url)
}

func NewApiError(req http.Request, res http.Response) ApiError {
	body, _ := io.ReadAll(res.Body)

	return ApiError{
		StatusCode: res.StatusCode,
		Method:     req.Method,
		Url:        req.URL.String(),
		Data:       body,
		Request:    req,
		Response:   res,
	}
}
