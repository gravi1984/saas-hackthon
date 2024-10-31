package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type City struct {
	Name    string
	Country string
}

type GeoCodingResponse struct {
	GeoCodingResults []GeoCodingResult `json:"geocoding_results"`
}

type GeoCodingResult struct {
	Name      string `json:"name"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Country   string `json:"country"`
}

type Location struct {
	Longitude string
	Latitude  string
}

type Forecast struct {
	Temperature string
	Location    Location
}

func GetWeather(loc Location) []byte {
	// formattedURL := fmt.Sprintf("https://api.weather.gov/gridpoints/TOP/%s,%s/forecast", loc.Longitude, loc.Latitude)
	response, err := http.Get("https://api.weather.gov/gridpoints/TOP/31,80/forecast")

	if err != nil {
		fmt.Print((err.Error()))
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Print((err.Error()))
		os.Exit(1)
	}

	return responseData
}

func FindCityLocation(city City) (string, string) {
	url := fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=10&language=en&format=json", city.Name)
	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))

	var geocoding_response GeoCodingResponse
	json.Unmarshal(responseData, &geocoding_response)

	for i := 0; i < len(geocoding_response.GeoCodingResults); i++ {
		if geocoding_response.GeoCodingResults[i].Country == city.Country {
			return geocoding_response.GeoCodingResults[i].Latitude, geocoding_response.GeoCodingResults[i].Longitude
		}
	}
	return "", ""
}

func main() {
	// Take input
	city := flag.String("city", "", "Name of the city")
	country := flag.String("country", "", "Country of the city")
	day := flag.String("day", "", "Day for the weather forecast (e.g., '2024-10-31')")
	flag.Parse()

	if *city == "" || *country == "" || *day == "" {
		log.Fatal("City, country, and day must be provided.")
	}

	// Call logic func
	lat, lon := FindCityLocation(City{Name: *city, Country: *country})

	loc := Location{
		Latitude:  lat,
		Longitude: lon,
	}

	weather := GetWeather(loc)

	// Print weather
	fmt.Printf("Weather in %s, %s on %s: %s\n", *city, *country, *day, string(weather))
}
