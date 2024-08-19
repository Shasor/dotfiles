package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func init() {
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Data structures
type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type CombinedData struct {
	Artist   Artist
	Location Location
	Date     Date
	Relation Relation
}

var combinedData []CombinedData

const (
	artistsURL   = "https://groupietrackers.herokuapp.com/api/artists"
	locationsURL = "https://groupietrackers.herokuapp.com/api/locations"
	datesURL     = "https://groupietrackers.herokuapp.com/api/dates"
	relationsURL = "https://groupietrackers.herokuapp.com/api/relation"
)

var (
	artists []Artist
	// locations []Location
	// dates     []Date
	// relations []Relation
)

func main() {
	// Load data from APIs
	if err := loadData(); err != nil {
		ErrorLogger.Fatalf("Failed to load data: %v", err)
	}

	// Set up HTTP routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api/artists", artistsHandler)
	http.HandleFunc("/api/artist/", singleArtistHandler)
	// http.HandleFunc("/api/search", searchHandler)

	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Start the server
	port := ":8080"
	InfoLogger.Printf("Server starting on %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		ErrorLogger.Fatalf("Server failed to start: %v", err)
	}
}

func loadData() error {
	var wg sync.WaitGroup
	wg.Add(4)

	errChan := make(chan error, 4)

	go func() {
		defer wg.Done()
		if err := fetchJSON(artistsURL, &artists.Members); err != nil {
			errChan <- fmt.Errorf("error fetching artists: %v", err)
		}
	}()
	fmt.Println(artists[0])

	// go func() {
	// 	defer wg.Done()
	// 	if err := fetchJSON(locationsURL, &locations); err != nil {
	// 		errChan <- fmt.Errorf("error fetching locations: %v", err)
	// 	}
	// }()

	// go func() {
	// 	defer wg.Done()
	// 	if err := fetchJSON(datesURL, &dates); err != nil {
	// 		errChan <- fmt.Errorf("error fetching dates: %v", err)
	// 	}
	// }()

	// go func() {
	// 	defer wg.Done()
	// 	if err := fetchJSON(relationsURL, &relations); err != nil {
	// 		errChan <- fmt.Errorf("error fetching relations: %v", err)
	// 	}
	// }()

	// wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			return err
		}
	}

	// Combine data
	combinedData = make([]CombinedData, len(artists))
	for i, artist := range artists {
		combinedData[i] = CombinedData{
			Artist: artist,
		}
	}

	return nil
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// Serve your HTML homepage
	http.ServeFile(w, r, "templates/index.html")
}

func artistsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(combinedData)
}

func singleArtistHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/api/artist/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	for _, artist := range combinedData {
		if artist.Artist.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(artist)
			return
		}
	}

	http.NotFound(w, r)
}

// func searchHandler(w http.ResponseWriter, r *http.Request) {
// 	query := r.URL.Query().Get("q")
// 	results := []CombinedData{}

// 	for _, artist := range combinedData {
// 		if strings.Contains(strings.ToLower(artist.Artist.Name), strings.ToLower(query)) {
// 			results = append(results, artist)
// 		}
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(results)
// }

// Helper functions for fetching data from APIs
func fetchJSON(url string, target any) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	err = json.NewDecoder(resp.Body).Decode(&target)
	if err != nil {
		return err
	}
	return nil
}
