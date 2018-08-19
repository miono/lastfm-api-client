package lastfmclient

import (
	"net/http"
	"net/url"
)

// Client is the api-client for the last.fm API
type Client struct {
	Apikey     string
	URL        *url.URL
	httpClient *http.Client
}

// NewClient creates a new API-client
func NewClient(apikey string) *Client {
	return &Client{
		Apikey: apikey,
		URL: &url.URL{
			Scheme: "https",
			Host:   "ws.audioscrobbler.com",
			Path:   "/2.0/",
		},
	}
}
