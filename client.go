package lastfmclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

// Track defines a track
type Track struct {
	Artist string // `json:"artist"`
	Title  string // `json:"title"`
}

// GetRecentTracks gets the recent tracks for a user
func (c *Client) GetRecentTracks(user string) []Track {
	url := fmt.Sprintf("method=user.getrecenttracks&user=%v&api_key=%v&format=json", user, c.Apikey)
	reqURL := c.BaseURL + url
	response, err := c.httpClient.Get(reqURL)
	if err != nil {
		panic(err)
	}
	responseSlice, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	var out Track
	json.Unmarshal(responseSlice, &out)
	var tracks []Track
	track := Track{"kent", "747"}
	tracks = append(tracks, track)
	track = Track{"kent", "Våga vara rädd"}
	tracks = append(tracks, track)
	return tracks
}
