package lastfmclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
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

// GetRecentTracks contains the full response from the getrecenttracks-method
type GetRecentTracks struct {
	RecentTracks `json:"recenttracks"`
}

// RecentTracks is the name of the list of tracks returned
type RecentTracks struct {
	RecentTracks []RecentTrack `json:"track"`
}

// RecentTrack is a track-struct
type RecentTrack struct {
	Artist     `json:"artist"`
	Name       string `json:"name"`
	Streamable int    `json:"streamable,string"`
	Mbid       string `json:"mbid"`
	Album      `json:"album"`
	URL        string  `json:"url"`
	Image      []Image `json:"image"`
	Date       Date    `json:"date"`
}

// Artist is the mbid and name of an artist
type Artist struct {
	Text string `json:"#text"`
	Mbid string `json:"mbid"`
}

// Album is the mbid and name of an album
type Album struct {
	Text string `json:"#text"`
	Mbid string `json:"mbid"`
}

// Image is the images for the album together with their sizes
type Image struct {
	Text string `json:"#text"`
	Sixe string `json:"size"`
}

// Date is the date when the track was played
type Date struct {
	Uts Uts `json:"uts"`
}

// Uts contains just the time-stamp
type Uts struct {
	Time time.Time
}

// UnmarshalJSON converts the Uts-field of the Date-struct to a time.Time
func (u *Uts) UnmarshalJSON(b []byte) error {
	timestampString := strings.Trim(string(b), "\"")
	timestampInt, err := strconv.Atoi(timestampString)
	if err != nil {
		return err
	}
	u.Time = time.Unix(int64(timestampInt), 0)
	return nil

}

// GetRecentTracks gets the recent tracks for a user
func (c *Client) GetRecentTracks(user string) []RecentTrack {
	url := fmt.Sprintf("method=user.getrecenttracks&user=%v&api_key=%v&format=json&limit=3", user, c.Apikey)
	reqURL := c.BaseURL + url
	response, err := c.httpClient.Get(reqURL)
	if err != nil {
		panic(err)
	}
	responseSlice, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	var ret GetRecentTracks
	var tracks []RecentTrack
	err = json.Unmarshal(responseSlice, &ret)
	if err != nil {
		panic(err)
	}
	for _, track := range ret.RecentTracks.RecentTracks {
		tracks = append(tracks, track)
	}

	return tracks
}
