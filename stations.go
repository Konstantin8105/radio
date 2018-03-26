package radio

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Example of station JSON output
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

const (
	stationListFilename = ".station.list.json"
)

// getStations return stations
// List of top stations:
// Post : http://shoutcast.com/Home/Top
func getStations() (stations []station, err error) {
	var jsonList []byte

	jsonList, err = ioutil.ReadFile(stationListFilename)
	if err != nil {
		// ignore error
		err = nil
		// download list from internet
		var buf bytes.Buffer
		var res *http.Response
		res, err = http.Post("http://shoutcast.com/Home/Top", "", &buf)
		if err != nil {
			err = fmt.Errorf("Cannot create post request. %v", err)
			return
		}
		jsonList, err = ioutil.ReadAll(res.Body)
		if err != nil {
			err = fmt.Errorf("Cannot read json. %v", err)
			return
		}
		defer res.Body.Close()

		_ = ioutil.WriteFile(stationListFilename, jsonList, 0644)
	}

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
