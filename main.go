package main

import (
	"fmt"
	"io"
	"net/http"
)

// Structs to match the API responses
type Artist struct {
	Name       string   `json:"name"`
	Image      string   `json:"image"`
	StartYear  int      `json:"start_year"`
	FirstAlbum string   `json:"first_album_date"`
	Members    []string `json:"members"`
}

type Location struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

type Date struct {
	Date string `json:"date"`
}

type Relation struct {
	ArtistID   string `json:"artist_id"`
	DateID     string `json:"date_id"`
	LocationID string `json:"location_id"`
}

// Fetch and return data from API
func fetchData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil

}

func getArtists(w http.ResponseWriter, r *http.Request) {
	data, err := fetchData("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func getLocations(w http.ResponseWriter, r *http.Request) {
	data, err := fetchData("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func getDates(w http.ResponseWriter, r *http.Request) {
	data, err := fetchData("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func getRelations(w http.ResponseWriter, r *http.Request) {
	data, err := fetchData("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func main() {
	http.HandleFunc("/api/artists", getArtists)
	http.HandleFunc("/api/locations", getLocations)
	http.HandleFunc("/api/dates", getDates)
	http.HandleFunc("/api/relation", getRelations)

	http.ListenAndServe(":8080", nil)
	fmt.Println("Serve running on port http://localhost:8080")
}
