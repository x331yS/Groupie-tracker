package gt_index

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
	Suiv         int
	Prec         int
	Name         string
	Image        string
	Members      []string
	CreationDate int
	FirstAlbum   string
	RelationData Relation
}

var Artists Artist
var Relations Relation
