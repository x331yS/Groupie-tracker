package gt_search

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func Data() {
	//Use API for Searh Handler
	allData = GatherDataUp("https://groupietrackers.herokuapp.com/api/artists")
	if allData == nil {
		fmt.Println("Failed to gather Data from API")
		os.Exit(1)
	}
}

func GatherDataUp(link string) []artistData {
	//Parse JSON for Search Handler
	data1 := GetData(link)
	Artists := []artistData{}
	e := json.Unmarshal(data1, &Artists)
	if e != nil {
		log.Fatal(e)
		return nil
	}
	for i := 0; i < len(Artists); i++ {
		r := relation{}
		_ = json.Unmarshal(GetData(Artists[i].Relation), &r)
		Artists[i].Concerts = r.Concerts
	}
	return Artists
}

func GetData(link string) []byte {
	//Get API for Searh Handler
	data1, e1 := http.Get(link)
	if e1 != nil {
		log.Fatal(e1)
		return nil
	}
	data2, e2 := ioutil.ReadAll(data1.Body)
	if e2 != nil {
		log.Fatal(e2)
		return nil
	}
	return data2
}
