package pkg

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Artist []struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Relation struct {
	Index []Index `json:"index"`
}

type Index struct {
	ID             int64               `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type About struct {
	ID           int
	Name         string
	Image        string
	Members      []string
	CreationDate int
	FirstAlbum   string
	RelationData Relation
}

var artists Artist
var relations Relation
var apiURL = "https://groupietrackers.herokuapp.com/api/"

func ConnectToAPI(s string) ([]byte, error) {
	resp, err := http.Get(apiURL + s)
	if err != nil {
		return nil, err
	}
	body, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		return nil, err1
	}
	return body, nil
}

func ConnnectParseAPI(w http.ResponseWriter, r *http.Request, api string) {
	body, err := ConnectToAPI(api)
	if err != nil {
		InternalServerError(w, r)
	}
	if api == "artists" {
		_ = json.Unmarshal(body, &artists)
	} else if api == "relation" {
		_ = json.Unmarshal(body, &relations)
	}
}