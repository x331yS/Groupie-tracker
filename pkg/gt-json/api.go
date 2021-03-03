package gt_json

import (
	"../gt-error"
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

type Search struct {
	Artists   Artist
	RelationS Relation
}

var Artists Artist
var Relations Relation
var ApiURL = "https://groupietrackers.herokuapp.com/api/"

func ConnectToAPI(s string) ([]byte, error) {
	resp, err := http.Get(ApiURL + s)
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
		gt_error.InternalServerError(w, r)
	}
	if api == "artists" {
		_ = json.Unmarshal(body, &Artists)
	} else if api == "relation" {
		_ = json.Unmarshal(body, &Relations)
	}
}