// Package cartodb provides a wrapper to the CartoDB API.
//
// You can create two kind of clients here: one authorised with OAuth or one
// authorised using the API Key provided by CartoDB.
//
// Included on the source of this packages you are going to find examples under
// the "examples/" folder.
package cartodb

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// BaseClient is the struct defined to interact with CartoDB API without any
// kind of authorization
type BaseClient struct {
	maxGetQueryLength int

	resourceURL string
}

// APIKeyClient is the struct defined to interact with CartoDB API being
// autheorised with the API Key.
type APIKeyClient struct {
	BaseClient

	apiKey string
}

// NewBaseClient returns a CartoDB client without any kind of authentication.
func NewBaseClient(username string, host string, protocol string, apiVersion string) *BaseClient {
	if host == "" {
		host = "cartodb.com"
	}
	if protocol == "" {
		protocol = "https"
	}
	if apiVersion == "" {
		apiVersion = "v2"
	}

	return &BaseClient{
		maxGetQueryLength: 2048,
		resourceURL:       fmt.Sprintf("%s://%s.%s/api/%s/sql", protocol, username, host, apiVersion),
	}
}

// NewAPIKeyClient returns a CartoDB client using the API Key authentication.
func NewAPIKeyClient(apiKey string, username string, host string, protocol string, apiVersion string) *APIKeyClient {
	if protocol != "" && protocol != "https" {
		log.Println("WAR: You are using this API key auth method with http")
	}
	return &APIKeyClient{
		BaseClient: *NewBaseClient(username, host, protocol, apiVersion),
		apiKey:     apiKey,
	}
}

// SQL is going to call CartoDB with the given SQL statement and return a response object.
// Is the user's responsability to know what to do with the response.Body.
func (c BaseClient) SQL(sql string, method string, format string) (response *http.Response, err error) {
	var params url.Values

	httpClient := &http.Client{}

	if method == "" {
		method = "GET"
	}
	// If the sql query is longer that the supported length we need to force the POST
	if len(sql) > c.maxGetQueryLength {
		method = "POST"
	}

	values := url.Values{"q": {sql}, "format": {format}}

	if method == "GET" {
		c.resourceURL = fmt.Sprintf("%s?%s", c.resourceURL, values.Encode())
	} else {
		params = values
	}

	req, _ := http.NewRequest(method, c.resourceURL, strings.NewReader(params.Encode()))

	response, err = httpClient.Do(req)
	return
}

func (c APIKeyClient) Req(httpURL string, httpMethod string, httpHeaders map[string]string, body string) (response *http.Response, err error) {
	httpClient := &http.Client{}

	params := url.Values{"api_key": {c.apiKey}, "body": {body}}
	req, _ := http.NewRequest(httpMethod, httpURL, strings.NewReader(params.Encode()))
	for k, v := range httpHeaders {
		req.Header.Add(k, v)
	}
	if httpMethod == "POST" {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	response, err = httpClient.Do(req)
	return
}
