package gt_search

type artistData struct {
	ID           int                 `json:"id"`
	Image        string              `json:"image"`
	Name         string              `json:"name"`
	Members      []string            `json:"members"`
	CreationDate int                 `json:"creationDate"`
	FirstAlbum   string              `json:"firstAlbum"`
	Relation     string              `json:"relations"`
	Concerts     map[string][]string `json:"datesLocations"`
}

type relation struct {
	ID       int                 `json:"id"`
	Concerts map[string][]string `json:"datesLocations"`
}

var allData []artistData


