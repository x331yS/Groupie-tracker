package gt_filter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//Main struct for Search
var All API

//Parse JSON from URL and write to All API struct
func ParseJSON() {
	urlArtists := "https://groupietrackers.herokuapp.com/api/artists"
	urlLocations := "https://groupietrackers.herokuapp.com/api/locations"
	urlDates := "https://groupietrackers.herokuapp.com/api/dates"
	urlRelation := "https://groupietrackers.herokuapp.com/api/relation"

	ParseInfo(urlArtists, &All.Artists)
	ParseInfo(urlLocations, &All.Locations)
	ParseInfo(urlDates, &All.Dates)
	ParseInfo(urlRelation, &All.Relation)
}

func ParseInfo(url string, temp interface{}) {
	//Get Data for Search API
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	json.Unmarshal(body, &temp)
}
