package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type countryData struct {
	Name      string
	Languages []string
	Phone     string
	Continent string
	Capital   string
	Currency  string
}

type countryMap map[string]countryData

func main() {

	allCountries := make(countryMap)
	countriesBytes, err := ioutil.ReadFile("countries.json")
	if err != nil {
		log.Fatal("Could not load countries data file")
	}

	err = json.Unmarshal(countriesBytes, &allCountries)
	if err != nil {
		log.Fatal("Could not unmarshal data file")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s", string(countriesBytes))
	})

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		term := r.FormValue("term")
		var results []countryData
		for countryCode, countryData := range allCountries {
			if strings.Contains(strings.ToLower(countryData.Name), strings.ToLower(term)) {
				results = append(results, countryData)
			} else if strings.Contains(strings.ToLower(countryCode), strings.ToLower(term)) {
				results = append(results, countryData)
			}
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(results)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
