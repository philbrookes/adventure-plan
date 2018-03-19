package maps

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//List contains an array of maps
type List struct {
	Maps []*Map `json:"maps"`
}

//Map defines map data
type Map struct {
	ID        int         `json:"id"`
	Center    MapPoint    `json:"center"`
	Zoom      int         `json:"zoom"`
	Interests []*MapPoint `json:"interests"`
	Metadata  MapMetadata `json:"metadata"`
}

//MapMetadata contains optional data about a particular map
type MapMetadata struct {
	Title string `json:"title,omitempty"`
}

//MapPoint defines a point on a map
type MapPoint struct {
	ID        int              `json:"id"`
	Latitude  float64          `json:"latitude"`
	Longitude float64          `json:"longitude"`
	Metadata  MapPointMetadata `json:"metadata"`
}

//MapPointMetadata contains optional data for a point of interest
type MapPointMetadata struct {
	Title string `json:"title,omitempty"`
	Notes string `json:"notes"`
}

//Maps Hander
type Maps struct {
	db *sql.DB
}

//NewMapsHandler creates a new users handler
func NewMapsHandler(db *sql.DB) *Maps {
	return &Maps{
		db: db,
	}
}

//ConfigureRouter to handle user related routes
func (m *Maps) ConfigureRouter(router *mux.Router) {
	router.HandleFunc("/", m.getMaps).Methods("GET")
	router.HandleFunc("/{id}", m.getMap).Methods("GET")
	router.HandleFunc("/{id}/pin", m.postPin).Methods("POST")
}

func (m *Maps) postPin(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	newPoint := &MapPoint{Metadata: MapPointMetadata{}}
	json.Unmarshal(b, newPoint)

	stmt, err := m.db.Prepare("INSERT INTO adventureplan.points (latitude, longitude, title, notes, map_id) VALUES (?, ?, ?, ?, ?);")
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	res, err := stmt.Exec(newPoint.Latitude, newPoint.Longitude, newPoint.Metadata.Title, newPoint.Metadata.Notes, id)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	ID, err := res.LastInsertId()
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	newPoint.ID = int(ID)
	json.NewEncoder(w).Encode(newPoint)
}

func (m *Maps) getMaps(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "Application/JSON")
	maps := List{Maps: []*Map{}}

	rows, err := m.db.Query("SELECT id, center_lat, center_lng, zoom, title FROM adventureplan.maps")
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	for rows.Next() {
		thisMap, err := loadMapFromRow(rows, m.db)
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}
		maps.Maps = append(maps.Maps, thisMap)
	}
	json.NewEncoder(w).Encode(maps)
}

func (m *Maps) getMap(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	w.Header().Add("Content-Type", "Application/JSON")

	rows, err := m.db.Query("SELECT id, center_lat, center_lng, zoom, title FROM adventureplan.maps WHERE id=?;", params["id"])
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	if !rows.Next() {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	thisMap, err := loadMapFromRow(rows, m.db)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(thisMap)

}

func loadMapFromRow(rows *sql.Rows, db *sql.DB) (*Map, error) {
	thisMap := Map{Center: MapPoint{}}
	err := rows.Scan(&thisMap.ID, &thisMap.Center.Latitude, &thisMap.Center.Longitude, &thisMap.Zoom, &thisMap.Metadata.Title)
	if err != nil {
		return &Map{}, err
	}

	err = loadMapPoints(db, &thisMap)
	if err != nil {
		return &Map{}, err
	}
	return &thisMap, nil
}

func loadMapPoints(db *sql.DB, mapObject *Map) error {
	pointsRS, err := db.Query("SELECT id, latitude, longitude, title, notes FROM adventureplan.points WHERE map_id = ?;", mapObject.ID)
	if err != nil {
		return err
	}
	for pointsRS.Next() {
		var thisPoint = MapPoint{Metadata: MapPointMetadata{}}
		err = pointsRS.Scan(&thisPoint.ID, &thisPoint.Latitude, &thisPoint.Longitude, &thisPoint.Metadata.Title, &thisPoint.Metadata.Notes)
		mapObject.Interests = append(mapObject.Interests, &thisPoint)
	}
	return nil
}
