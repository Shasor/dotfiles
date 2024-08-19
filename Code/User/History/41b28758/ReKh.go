package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const port = ":8080"

// Every structures that datas from JSON files will be put into

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type ArtistLocation struct {
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
	} `json:"index"`
}

type ArtistDates struct {
	Dates []string `json:"dates"`
}

type ArtistRelation struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}

type ArtistPageData struct {
	Artist          Artist
	Locations       ArtistLocation
	Dates           ArtistDates
	Relations       ArtistRelation
	ArtistLocations []string
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/artist/", ArtistPage)
	http.HandleFunc("/search", SearchArtists)
	http.HandleFunc("/filtered-results", FilteredResults)
	http.HandleFunc("/locations-autocomplete", LocationsAutocomplete)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	fmt.Println("(http://localhost:8080) - Server started on port", port)
	server := &http.Server{
		Addr:              port,              //adresse du server (le port choisi est à titre d'exemple)
		ReadHeaderTimeout: 10 * time.Second,  // temps autorisé pour lire les headers
		WriteTimeout:      10 * time.Second,  // temps maximum d'écriture de la réponse
		IdleTimeout:       120 * time.Second, // temps maximum entre deux rêquetes
		MaxHeaderBytes:    1 << 20,           // 1 MB // maxinmum de bytes que le serveur va lire
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

// Will render any template .html from the templates directory
func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {

	funcMap := template.FuncMap{
		"json": func(v interface{}) template.JS {
			a, err := json.Marshal(v)
			if err != nil {
				log.Printf("Error marshaling JSON: %v", err)
				return template.JS("{}")
			}
			return template.JS(a)
		},
	}
	// include in map js func
	page, err := template.New(tmpl + ".html").Funcs(funcMap).ParseFiles("templates/" + tmpl + ".html")

	if err != nil {
		w.WriteHeader(404)
		http.Error(w, "error 404", http.StatusNotFound)
		log.Printf("error template %v", err)
		return
	}
	err = page.Execute(w, data)
	if err != nil {
		http.Error(w, "Error 500, Internal server error", http.StatusInternalServerError)
		log.Printf("error template %v", err)
		return
	}
}

// Page where all artists are displayed
func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		RenderTemplate(w, "error404", nil)
	} else {
		var artists []Artist
		urlAPI := "https://groupietrackers.herokuapp.com/api/artists"

		err := FetchData(urlAPI, &artists)
		if err != nil {
			Error(w, http.StatusInternalServerError, "Error fetching artist data")
			return
		}

		RenderTemplate(w, "index", artists)
	}
}

// Iterates through all artists and append into a new tab of strings, according to the filter values used
func filterArtists(artists []Artist, creationDateMin, creationDateMax, firstAlbumMin, firstAlbumMax int, memberCounts []string, locations string) []Artist {
	var filtered []Artist
	var locationsData ArtistLocation
	err := FetchData("https://groupietrackers.herokuapp.com/api/locations", &locationsData)
	if err != nil {
		log.Printf("Error fetching locations: %v", err)
		return filtered
	}

	for _, artist := range artists {
		if artist.CreationDate < creationDateMin || artist.CreationDate > creationDateMax {
			continue
		}

		firstAlbumYear, _ := strconv.Atoi(strings.Split(artist.FirstAlbum, "-")[2])
		if firstAlbumYear < firstAlbumMin || firstAlbumYear > firstAlbumMax {
			continue
		}

		if !contains(memberCounts, strconv.Itoa(len(artist.Members))) {
			continue
		}

		//handle the use of location filter
		if locations != "" {
			artistLocations := getArtistLocations(locationsData, artist.ID)
			if !containsLocation(artistLocations, locations) {
				continue
			}
		}

		filtered = append(filtered, artist)
	}

	return filtered
}

// Iterates through all of the locations API and verify if the Artist ID is the same in both APIs
func getArtistLocations(locationsData ArtistLocation, artistID int) []string {
	for _, item := range locationsData.Index {
		if item.ID == artistID {
			return item.Locations
		}
	}
	return nil
}

func containsLocation(locations []string, query string) bool {
	query = strings.ToLower(query)
	for _, location := range locations {
		if strings.Contains(strings.ToLower(location), query) {
			return true
		}
	}
	return false
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func FilteredResults(w http.ResponseWriter, r *http.Request) {
	// Parse filter parameters
	creationDateMin, _ := strconv.Atoi(r.URL.Query().Get("creationDateMin"))
	creationDateMax, _ := strconv.Atoi(r.URL.Query().Get("creationDateMax"))
	firstAlbumMin, _ := strconv.Atoi(r.URL.Query().Get("firstAlbumMin"))
	firstAlbumMax, _ := strconv.Atoi(r.URL.Query().Get("firstAlbumMax"))
	memberCounts := strings.Split(r.URL.Query().Get("memberCounts"), ",")
	locations := r.URL.Query().Get("locations")

	// Fetch all artists
	var allArtists []Artist
	err := FetchData("https://groupietrackers.herokuapp.com/api/artists", &allArtists)
	if err != nil {
		Error(w, http.StatusInternalServerError, "Error fetching artist data")
		return
	}

	// Filter artists
	filteredArtists := filterArtists(allArtists, creationDateMin, creationDateMax, firstAlbumMin, firstAlbumMax, memberCounts, locations)

	// Render the template with filtered artists
	RenderTemplate(w, "index", filteredArtists)
}

func ArtistPage(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	if len(path) < 3 {
		Error(w, http.StatusBadRequest, "Invalid artist ID")
		return
	}
	artistID, err := strconv.Atoi(path[2])
	if err != nil {
		Error(w, http.StatusBadRequest, "Invalid artist ID")
		return
	}

	if artistID <= 0 || artistID > 52 {
		Error(w, http.StatusBadRequest, "Invalid artist ID")
		return
	}

	var pageData ArtistPageData

	// Fetch artist data
	artistURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%d", artistID)
	err = FetchData(artistURL, &pageData.Artist)
	if err != nil {
		Error(w, http.StatusInternalServerError, "Error fetching artist data")
		return
	}

	// Fetch relations
	relationsURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%d", artistID)
	err = FetchData(relationsURL, &pageData.Relations)
	if err != nil {
		Error(w, http.StatusInternalServerError, "Error fetching relations data")
		return
	}

	locations := make([]string, 0)
	for location := range pageData.Relations.DatesLocations {
		locations = append(locations, location)
	}
	pageData.ArtistLocations = locations

	RenderTemplate(w, "artist", pageData)
}

func LocationsAutocomplete(w http.ResponseWriter, r *http.Request) {
	locations, err := getAllUniqueLocations()
	if err != nil {
		http.Error(w, "Error fetching locations", http.StatusInternalServerError)
		return
	}

	query := strings.ToLower(r.URL.Query().Get("q"))
	var matches []string
	for _, location := range locations {
		if strings.Contains(strings.ToLower(location), query) {
			matches = append(matches, location)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(matches)
}

func getAllUniqueLocations() ([]string, error) {
	var locationsData ArtistLocation
	err := FetchData("https://groupietrackers.herokuapp.com/api/locations", &locationsData)
	if err != nil {
		return nil, err
	}

	uniqueLocations := make(map[string]bool)
	for _, item := range locationsData.Index {
		for _, location := range item.Locations {
			uniqueLocations[location] = true
		}
	}

	var result []string
	for location := range uniqueLocations {
		result = append(result, location)
	}

	return result, nil
}

// Will Fetch Data from the JSON file
func FetchData(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("HTTP request error: %v", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Unexpected status code: %d", resp.StatusCode)
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		log.Printf("JSON decode error: %v", err)
		return err
	}

	return nil
}

func SearchArtists(w http.ResponseWriter, r *http.Request) {
	query := strings.ToLower(r.URL.Query().Get("q"))

	var artists []Artist
	var locationsData struct {
		Index []struct {
			ID        int      `json:"id"`
			Locations []string `json:"locations"`
		} `json:"index"`
	}

	// Fetch artists data
	err := FetchData("https://groupietrackers.herokuapp.com/api/artists", &artists)
	if err != nil {
		Error(w, http.StatusInternalServerError, "Error fetching artist data")
		return
	}

	// Fetch locations data
	err = FetchData("https://groupietrackers.herokuapp.com/api/locations", &locationsData)
	if err != nil {
		Error(w, http.StatusInternalServerError, "Error fetching location data")
		return
	}

	suggestions := []SearchSuggestion{}

	for i, artist := range artists {
		// Check artist/band name
		if strings.Contains(strings.ToLower(artist.Name), query) {
			suggestions = append(suggestions, SearchSuggestion{Value: artist.Name, Type: "artist/band", ID: artist.ID})
		}

		// Check creation date
		creationDateStr := strconv.Itoa(artist.CreationDate)
		if strings.Contains(creationDateStr, query) {
			suggestions = append(suggestions, SearchSuggestion{Value: creationDateStr, Type: "creation date", ID: artist.ID})
		}

		// Check members (not visible on index page, but searchable)
		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), query) {
				suggestions = append(suggestions, SearchSuggestion{Value: member, Type: "member", ID: artist.ID})
			}
		}

		// Check first album (not visible on index page, but searchable)
		if strings.Contains(strings.ToLower(artist.FirstAlbum), query) {
			suggestions = append(suggestions, SearchSuggestion{Value: artist.FirstAlbum, Type: "first album", ID: artist.ID})
		}

		// Check locations (not visible on index page, but searchable)
		if i < len(locationsData.Index) {
			for _, location := range locationsData.Index[i].Locations {
				if strings.Contains(strings.ToLower(location), query) {
					suggestions = append(suggestions, SearchSuggestion{Value: location, Type: "location", ID: artist.ID})
				}
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(suggestions)
}

type SearchSuggestion struct {
	Value string `json:"value"`
	Type  string `json:"type"`
	ID    int    `json:"id"`
}

func Error(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	var tmpl string
	switch status {
	case http.StatusBadRequest:
		tmpl = "error400"
	case http.StatusNotFound:
		tmpl = "error404"
	case http.StatusInternalServerError:
		tmpl = "error500"
	default:
		tmpl = "unexpected_error"
		log.Printf("Unexpected error status: %d", status)
	}
	page, err := template.ParseFiles("templates/" + tmpl + ".html")
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		log.Printf("Error parsing template: %v", err)
		return
	}
	data := struct {
		Status  int
		Message string
	}{
		Status:  status,
		Message: message,
	}
	err = page.Execute(w, data)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
	}
}

// func formatLocation(locationStr string) string {
// 	regionAndCountry := strings.Split(locationStr, "-")
// 	regionWords := strings.Split(regionAndCountry[0], "_")
// 	for i := range regionWords {
// 		regionWords[i] = strings.ToUpper(regionWords[i][:1]) + strings.ToLower(regionWords[i][1:])
// 	}
// 	region := strings.Join(regionWords, "-")
// 	countryWords := strings.Split(regionAndCountry[1], "_")
// 	for i := range countryWords {
// 		countryWords[i] = strings.ToUpper(countryWords[i][:1]) + strings.ToLower(countryWords[i][1:])
// 	}
// 	country := strings.Join(countryWords, " ")
// 	return fmt.Sprintf("%s (%s)", region, country)
// }
