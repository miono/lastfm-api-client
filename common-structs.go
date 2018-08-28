package lastfmclient

import (
	"strconv"
	"strings"
	"time"
)

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
