package gt_index

import (
	"../gt-error"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//URL
var ApiURL = "https://groupietrackers.herokuapp.com/api/"

func ConnectToAPI(s string) ([]byte, error) {
	//Get API for Index Handler
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
	//Parse JSON for IndexHandler
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