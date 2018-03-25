package radio

import (
	"encoding/json"
	"fmt"
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

// TopStations - top of stations
// List of top stations:
// Post : http://shoutcast.com/Home/Top
func TopStations() {
	/*
		var buf bytes.Buffer
		res, err := http.Post("http://shoutcast.com/Home/Top", "", &buf)
		if err != nil {
			fmt.Println("err1 = >>> ", err)
		}
		robots, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("err2 = >>> ", err)
		}
		defer res.Body.Close()
		fmt.Printf("%s", robots)
	*/

	var err error
	robots := []byte(`[
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
	      },
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
	      ]`)

	var stations []station
	err = json.Unmarshal(robots, &stations)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v\n", stations)
}

// List of stations by genrename:
// Post : http://shoutcast.com/Home/BrowseByGenre?genrename=50s

// Search stations:
// Post : http://shoutcast.com/Search/UpdateSearch
