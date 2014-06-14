package cartodb_go

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("shit")
}

type CartoDBBase struct {
	maxGetQueryLength int

	accessTokenURL string
	resourceURL    string
}

func (c CartoDBBase) Sql(sql string, doPost bool, format string) (response *http.Response, err error) {
	var params url.Values

	httpClient := &http.Client{}

	method := "GET"
	if doPost || len(sql) > c.maxGetQueryLength {
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

func NewCartoDBBase(cartodbDomain string, host string, protocol string, apiVersion string) *CartoDBBase {
	if host == "" {
		host = "cartodb.com"
	}
	if protocol == "" {
		protocol = "https"
	}
	if apiVersion == "" {
		apiVersion = "v2"
	}

	return &CartoDBBase{
		maxGetQueryLength: 2048,
		accessTokenURL:    fmt.Sprintf("%s://%s.%s/oauth/access_token", protocol, cartodbDomain, host),
		resourceURL:       fmt.Sprintf("%s://%s.%s/api/%s/sql", protocol, cartodbDomain, host, apiVersion),
	}
}

type CartoDBOAuth struct {
	CartoDBBase
}

func NewCartoDBOAuth(cartodbDomain string, host string, protocol string, apiVersion string) CartoDBOAuth {
	return CartoDBOAuth{*NewCartoDBBase(cartodbDomain, host, protocol, apiVersion)}
}

type CartoDBAPIKey struct {
	CartoDBBase

	apiKey string
}

func NewCartoDBAPIKey(apiKey string, cartodbDomain string, host string, protocol string, apiVersion string) *CartoDBAPIKey {
	if protocol != "" && protocol != "https" {
		log.Println("WAR: You are using this API key auth method with http")
	}
	return &CartoDBAPIKey{
		CartoDBBase: *NewCartoDBBase(cartodbDomain, host, protocol, apiVersion),
		apiKey:      apiKey,
	}
}

func (c CartoDBAPIKey) Req(httpUrl string, httpMethod string, httpHeaders map[string]string, body string) (response *http.Response, err error) {
	httpClient := &http.Client{}

	params := url.Values{"api_key": {c.apiKey}, "body": {body}}
	req, _ := http.NewRequest(httpMethod, httpUrl, strings.NewReader(params.Encode()))
	for k, v := range httpHeaders {
		req.Header.Add(k, v)
	}
	if httpMethod == "POST" {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	response, err = httpClient.Do(req)
	return
}
