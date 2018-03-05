package maps

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//List contains an array of maps
type List struct {
	Maps []Map `json:"maps"`
}

//Map defines map data
type Map struct {
	ID        int         `json:"id"`
	Center    MapPoint    `json:"center"`
	Zoom      int         `json:"zoom"`
	Interests []MapPoint  `json:"interests"`
	Metadata  MapMetadata `json:"metadata"`
}

//MapMetadata contains optional data about a particular map
type MapMetadata struct {
	Title string `json:"title,omitempty"`
}

//MapPoint defines a point on a map
type MapPoint struct {
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
	Maps List
}

//NewMapsHandler creates a new users handler
func NewMapsHandler() *Maps {
	return &Maps{
		Maps: List{
			Maps: []Map{
				Map{
					ID: 1,
					Center: MapPoint{
						Latitude:  41.3937688,
						Longitude: 2.128728,
					},
					Metadata: MapMetadata{
						Title: "Barcelona",
					},
					Zoom: 12,
					Interests: []MapPoint{
						MapPoint{
							Latitude:  41.3946688,
							Longitude: 2.179728,
							Metadata: MapPointMetadata{
								Title: "Barcelona is interesting...",
								Notes: "Suuuuuuper interesting...",
							},
						},
						MapPoint{
							Latitude:  41.3947698,
							Longitude: 2.078733,
							Metadata: MapPointMetadata{
								Title: "Barcelona is more interesting...",
								Notes: "Really suuuuuuper interesting...",
							},
						},
						MapPoint{
							Latitude:  41.3246688,
							Longitude: 2.129728,
							Metadata: MapPointMetadata{
								Title: "Barcelona is really more interesting...",
								Notes: "Really really suuuuuuper interesting...",
							},
						},
					},
				},
				Map{
					ID: 2,
					Center: MapPoint{
						Latitude:  4.3937688,
						Longitude: 22.128728,
					},
					Metadata: MapMetadata{
						Title: "Somewhere crap...",
					},
					Zoom:      12,
					Interests: []MapPoint{},
				},
			},
		},
	}

}

//ConfigureRouter to handle user related routes
func (m *Maps) ConfigureRouter(router *mux.Router) {
	router.HandleFunc("/", m.getMaps).Methods("GET")
	router.HandleFunc("/{id}", m.getMap).Methods("GET")
	router.HandleFunc("/{id}/pin", m.postPin).Methods("POST", "OPTIONS")
}

func (m *Maps) postPin(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	lat, _ := strconv.Atoi(r.FormValue("latitude"))
	lng, _ := strconv.Atoi(r.FormValue("longitude"))

	mp := MapPoint{
		Latitude:  float64(lat),
		Longitude: float64(lng),
		Metadata: MapPointMetadata{
			Title: r.FormValue("title"),
			Notes: r.FormValue("notes"),
		},
	}

	for _, mapData := range m.Maps.Maps {
		if mapData.ID == id {
			mapData.Interests = append(mapData.Interests, mp)
			w.Header().Add("Content-Type", "Application/JSON")
			json.NewEncoder(w).Encode(mapData)
		}
	}
}

func (m *Maps) getMaps(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "Application/JSON")
	json.NewEncoder(w).Encode(m.Maps)
}

func (m *Maps) getMap(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	w.Header().Add("Content-Type", "Application/JSON")
	json.NewEncoder(w).Encode(params)
}
