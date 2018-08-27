package lastfmclient

import (
	"net/http"
	"time"
)

// Client is the api-client for the last.fm API
type Client struct {
	Apikey     string
	BaseURL    string
	httpClient *http.Client
}

// NewClient creates a new API-client
func NewClient(apikey string) *Client {
	return &Client{
		Apikey:  apikey,
		BaseURL: "https://ws.audioscrobbler.com/2.0/?",
		httpClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}
