package lastfmclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// GetTopAlbums contains the full JSON-response
type GetTopAlbums struct {
	TopAlbums `json:"topalbums"`
}

// TopAlbums is the name of the list of albums returned
type TopAlbums struct {
	TopAlbums []TopAlbum `json:"album"`
}

// TopAlbum is an album on the top-list
type TopAlbum struct {
	Name         string         `json:"name"`
	PlayCount    int            `json:"playcount,string"`
	Mbid         string         `json:"mbid"`
	URL          string         `json:"url"`
	Artist       TopAlbumArtist `json:"artist"`
	Image        []Image        `json:"image"`
	TopAlbumAttr TopAlbumAttr   `json:"@attr"`
}

// TopAlbumAttr contains the rank of the album for the user
type TopAlbumAttr struct {
	Rank int `json:"rank,string"`
}

// TopAlbumArtist is an artist in the format given by GetTopAlbums
type TopAlbumArtist struct {
	Name string `json:"name"`
	Mbid string `json:"mbid"`
	URL  string `json:"url"`
}

// GetTopAlbums gets the Top Album for a user. Supply user, limit and period (overall | 7day | 1month | 3month | 6month | 12month)
func (c *Client) GetTopAlbums(user string, limit int, period string) []TopAlbum {
	url := fmt.Sprintf("method=user.gettopalbums&user=%v&api_key=%v&limit=%d&period=%v&format=json", user, c.Apikey, limit, period)
	reqURL := c.BaseURL + url
	response, err := c.httpClient.Get(reqURL)
	if err != nil {
		panic(err)
	}
	responseSlice, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	var ret GetTopAlbums
	var albums []TopAlbum
	err = json.Unmarshal(responseSlice, &ret)
	if err != nil {
		panic(err)
	}
	for _, album := range ret.TopAlbums.TopAlbums {
		albums = append(albums, album)
	}
	return albums
}
