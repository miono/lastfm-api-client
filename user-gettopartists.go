package lastfmclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// GetTopArtists contains the full JSON-response
type GetTopArtists struct {
	TopArtists `json:"topartists"`
}

// TopArtists is the name of the list of artists returned
type TopArtists struct {
	TopArtists []TopArtist `json:"artist"`
}

// TopArtist contains an instance of the artist on the list
type TopArtist struct {
	Name      string        `json:"name"`
	PlayCount int           `json:"playcount,string"`
	Mbid      string        `json:"mbid"`
	URL       string        `json:"url"`
	Image     []Image       `json:"image"`
	Attr      TopArtistAttr `json:"@attr"`
}

// TopArtistAttr contains the Rank of the artist for the user
type TopArtistAttr struct {
	Rank int `json:"rank,string"`
}

// GetTopArtists gets the Top Artists for a user. Supply user, limit and period (overall | 7day | 1month | 3month | 6month | 12month)
func (c *Client) GetTopArtists(user string, limit int, period string) []TopArtist {
	url := fmt.Sprintf("method=user.gettopartists&user=%v&api_key=%v&limit=%d&period=%v&format=json", user, c.Apikey, limit, period)
	reqURL := c.BaseURL + url
	response, err := c.httpClient.Get(reqURL)
	if err != nil {
		panic(err)
	}
	responseSlice, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	var ret GetTopArtists
	var artists []TopArtist
	err = json.Unmarshal(responseSlice, &ret)
	if err != nil {
		panic(err)
	}
	for _, artist := range ret.TopArtists.TopArtists {
		artists = append(artists, artist)
	}
	return artists
}
