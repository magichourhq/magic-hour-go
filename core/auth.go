package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	http "net/http"
	"strings"
	"time"

	"github.com/qri-io/jsonpointer"
)

type AuthProvider interface {
	Apply(*http.Request) error
	SetValue(*string)
}

// --------- AUTH BASIC ---------

type AuthBasic struct {
	username string
	password string
}

func NewAuthBasic(username string, password string) *AuthBasic {
	return &AuthBasic{username: username, password: password}
}
func (a *AuthBasic) Apply(req *http.Request) error {
	req.SetBasicAuth(a.username, a.password)
	return nil
}
func (a *AuthBasic) SetValue(val *string) {
	if val != nil {
		a.username = *val
	}
}

// --------- AUTH BEARER TOKEN ---------

type AuthBearer struct {
	token string
}

func NewAuthBearer(token string) *AuthBearer {
	return &AuthBearer{token: token}
}
func (a *AuthBearer) Apply(req *http.Request) error {
	req.Header.Add("Authorization", "Bearer "+a.token)
	return nil
}
func (a *AuthBearer) SetValue(val *string) {
	if val != nil {
		a.token = *val
	}
}

// --------- AUTH KEY (header/query/cookie) ---------

type AuthKey struct {
	location string
	name     string
	value    string
}

func NewAuthKeyHeader(name string, value string) *AuthKey {
	return &AuthKey{location: "header", name: name, value: value}
}
func NewAuthKeyQuery(name string, value string) *AuthKey {
	return &AuthKey{location: "query", name: name, value: value}
}
func NewAuthKeyCookie(name string, value string) *AuthKey {
	return &AuthKey{location: "cookie", name: name, value: value}
}
func (a *AuthKey) Apply(req *http.Request) error {
	switch a.location {
	case "header":
		req.Header.Add(a.name, a.value)
	case "query":
		queryParams := req.URL.Query()
		queryParams.Add(a.name, a.value)
		req.URL.RawQuery = queryParams.Encode()
	case "cookie":
		authCookie := http.Cookie{Name: a.name, Value: a.value}
		req.AddCookie(&authCookie)
	default:
		fmt.Printf("Invalid auth key (%s) location %s, no auth addded to request", a.name, a.value)
	}

	return nil
}
func (a *AuthKey) SetValue(val *string) {
	if val != nil {
		a.value = *val
	}
}

// --------- OAUTH2 ---------

type OAuth2Password struct {
	Username     string
	Password     string
	ClientId     *string
	ClientSecret *string
	GrantType    *string
	Scope        *[]string
	TokenUrl     string
}

type OAuth2ClientCredentials struct {
	ClientId     string
	ClientSecret string
	GrantType    *string
	Scope        *[]string
	TokenUrl     string
}

type OAuth2 struct {
	// OAuth2 provider configuration
	baseUrl             string
	tokenUrl            string
	accessTokenPointer  string
	expiresInPointer    string
	credentialsLocation string
	bodyContent         string
	requestMutator      AuthProvider

	// OAuth2 access token request values
	username     *string
	password     *string
	clientId     *string
	clientSecret *string
	grantType    string
	scope        *[]string

	// access token retention
	accessToken *string
	expiresAt   *time.Time
}

func NewOAuth2Password(
	baseUrl string,
	defaultTokenUrl string,
	accessTokenPointer string,
	expiresInPointer string,
	credentialsLocation string,
	bodyContent string,
	requestMutator AuthProvider,
	form OAuth2Password) *OAuth2 {

	grantType := "password"
	if form.GrantType != nil {
		grantType = *form.GrantType
	}

	tokenUrl := defaultTokenUrl
	if form.TokenUrl != "" {
		tokenUrl = form.TokenUrl
	}

	return &OAuth2{
		baseUrl:             baseUrl,
		tokenUrl:            tokenUrl,
		accessTokenPointer:  accessTokenPointer,
		expiresInPointer:    expiresInPointer,
		credentialsLocation: credentialsLocation,
		bodyContent:         bodyContent,
		requestMutator:      requestMutator,

		username:     &form.Username,
		password:     &form.Password,
		clientId:     form.ClientId,
		clientSecret: form.ClientSecret,
		grantType:    grantType,
		scope:        form.Scope,

		accessToken: nil,
		expiresAt:   nil,
	}
}
func NewOAuth2ClientCredentials(
	baseUrl string,
	defaultTokenUrl string,
	accessTokenPointer string,
	expiresInPointer string,
	credentialsLocation string,
	bodyContent string,
	requestMutator AuthProvider,
	form OAuth2ClientCredentials) *OAuth2 {

	grantType := "client_credentials"
	if form.GrantType != nil {
		grantType = *form.GrantType
	}
	tokenUrl := defaultTokenUrl
	if form.TokenUrl != "" {
		tokenUrl = form.TokenUrl
	}

	return &OAuth2{
		baseUrl:             baseUrl,
		tokenUrl:            tokenUrl,
		accessTokenPointer:  accessTokenPointer,
		expiresInPointer:    expiresInPointer,
		credentialsLocation: credentialsLocation,
		bodyContent:         bodyContent,
		requestMutator:      requestMutator,

		username:     nil,
		password:     nil,
		clientId:     &form.ClientId,
		clientSecret: &form.ClientSecret,
		grantType:    grantType,
		scope:        form.Scope,

		accessToken: nil,
		expiresAt:   nil,
	}
}
func (a *OAuth2) Refresh() error {
	url := a.tokenUrl
	if strings.HasPrefix(a.tokenUrl, "/") {
		// tokenUrl is relative
		base := strings.TrimRight(a.baseUrl, "/")
		path := strings.TrimLeft(a.tokenUrl, "/")
		url = strings.TrimRight((base + "/" + path), "/")
	}

	// create data
	data := map[string]string{"grant_type": a.grantType}

	if a.clientId != nil && a.credentialsLocation == "request_body" {
		data["client_id"] = *a.clientId
	}
	if a.clientSecret != nil && a.credentialsLocation == "request_body" {
		data["client_secret"] = *a.clientSecret
	}
	if a.username != nil {
		data["username"] = *a.username
	}
	if a.password != nil {
		data["password"] = *a.password
	}
	if a.scope != nil {
		data["scope"] = strings.Join(*a.scope, " ")
	}

	var reqBody io.Reader
	var contentType string
	if a.bodyContent == "json" {
		jsonBody, err := json.Marshal(data)
		if err != nil {
			return err
		}
		reqBody = bytes.NewBuffer([]byte(jsonBody))
		contentType = "application/json"
	} else {
		formBody, err := FormUrlEncodedBody(data, map[string]string{}, map[string]bool{})
		if err != nil {
			return err
		}
		reqBody = formBody
		contentType = "application/x-www-form-urlencoded"
	}

	// init request
	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", contentType)

	// optionally add client creds as basic auth
	if a.credentialsLocation == "basic_authorization_header" && (a.clientId != nil || a.clientSecret != nil) {
		username := ""
		if a.clientId != nil {
			username = *a.clientId
		}
		password := ""
		if a.clientSecret != nil {
			password = *a.clientSecret
		}
		req.SetBasicAuth(username, password)
	}

	// send req
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode >= 300 {
		return NewApiError(*req, *res)
	}

	// extract expiry and access token
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	var resBody map[string]interface{}
	err = json.Unmarshal(body, &resBody)

	tokenPtr, err := jsonpointer.Parse(a.accessTokenPointer)
	if err != nil {
		return err
	}
	tokenVal, err := tokenPtr.Eval(resBody)
	if err != nil {
		return err
	}
	if strVal, ok := tokenVal.(string); ok {
		a.accessToken = &strVal
	}

	expiresPtr, err := jsonpointer.Parse(a.expiresInPointer)
	if err != nil {
		return err
	}
	expiresVal, err := expiresPtr.Eval(resBody)
	if err != nil {
		return err
	}
	if floatVal, ok := expiresVal.(float64); ok {
		expire := time.Now().UTC()
		expire = expire.Add(time.Duration(floatVal) * time.Second)
		a.expiresAt = &expire
	}

	return nil
}

func (a *OAuth2) Apply(req *http.Request) error {
	tokenExpired := false
	if a.expiresAt != nil {
		tokenExpired = a.expiresAt.Before(time.Now().UTC())
	}

	if a.accessToken == nil || tokenExpired {
		err := a.Refresh()
		if err != nil {
			return err
		}
	}

	a.requestMutator.SetValue(a.accessToken)
	a.requestMutator.Apply(req)

	return nil
}
func (a *OAuth2) SetValue(val *string) {
	panic("an OAuth2 auth provider cannot be a requestMutator")
}
