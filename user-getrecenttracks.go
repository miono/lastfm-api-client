package lastfmclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

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
	PlayDate   time.Time
	TempDate   Date `json:"date"`
}

// GetRecentTracks gets the recent tracks for a user
func (c *Client) GetRecentTracks(user string, limit int) []RecentTrack {
	url := fmt.Sprintf("method=user.getrecenttracks&user=%v&api_key=%v&format=json&limit=%d", user, c.Apikey, limit)
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
		track.PlayDate = track.TempDate.Uts.Time
		tracks = append(tracks, track)
	}

	return tracks
}
