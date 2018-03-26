package radio

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Example of JSON output
/*
   {
       "ID": 65504,
       "Name": "Deleted 346",
       "Format": "audio/mpeg",
       "Bitrate": 128,
       "Genre": "50s",
       "CurrentTrack": "Mitch Miller - March From The River Kwai And Colonel Bogey ",
       "Listeners": 0,
       "IsRadionomy": false,
       "IceUrl": "",
       "StreamUrl": null,
       "AACEnabled": 0,
       "IsPlaying": false,
       "IsAACEnabled": false
   }
*/
type station struct {
	ID    int
	Name  string
	Genre string
}

// getStations return stations
// List of top stations:
// Post : http://shoutcast.com/Home/Top
func getStations() (stations []station, err error) {
	var buf bytes.Buffer
	res, err := http.Post("http://shoutcast.com/Home/Top", "", &buf)
	if err != nil {
		err = fmt.Errorf("Cannot create post request. %v", err)
		return
	}
	jsonList, err := ioutil.ReadAll(res.Body)
	if err != nil {
		err = fmt.Errorf("Cannot read json. %v", err)
		return
	}
	defer res.Body.Close()

	err = json.Unmarshal(jsonList, &stations)
	if err != nil {
		return
	}

	fmt.Printf("Found : %d stations\n", len(stations))

	return
}

// List of stations by genrename:
// Post : http://shoutcast.com/Home/BrowseByGenre?genrename=50s

// Search stations:
// Post : http://shoutcast.com/Search/UpdateSearch
