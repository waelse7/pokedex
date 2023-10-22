package api

import (
	"net/http"
	"time"
)

var httpClient *http.Client

func init(){
	httpClient = &http.Client{
		Timeout: 1 * time.Minute,
	}
}

func getHTTPClient() *http.Client {
	return httpClient
}