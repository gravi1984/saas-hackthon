package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type City struct {
	Name    string
	Country string
}

type GeoCodingResponse struct {
	GeoCodingResults []GeoCodingResult `json:"results"`
}

type GeoCodingResult struct {
	Name      string      `json:"name"`
	Latitude  json.Number `json:"latitude"`
	Longitude json.Number `json:"longitude"`
	Country   string      `json:"country"`
}

type Location struct {
	Longitude string
	Latitude  string
}

type ForecastParams struct {
	Precipitation bool
	Sunrise       bool
	Sunset        bool
	UVIndex       bool
	Fahr          bool
}

func formatExtraForecastParams(f ForecastParams) string {
	var formattedParams strings.Builder
	if f.Precipitation {
		formattedParams.WriteString(",precipitation_sum")
	}
	if f.Sunrise {
		formattedParams.WriteString(",sunrise")
	}
	if f.Sunset {
		formattedParams.WriteString(",sunset")
	}
	if f.UVIndex {
		formattedParams.WriteString(",uv_index_max")
	}
	if f.Fahr {
		formattedParams.WriteString("&temperature_unit=fahrenheit")
	}
	return formattedParams.String()
}

func GetWeather(loc Location, forecast_params ForecastParams) ([]byte, error) {
	var formattedUrl strings.Builder
	formattedUrl.WriteString("https://api.open-meteo.com/v1/forecast?")
	formattedUrl.WriteString(fmt.Sprintf(
		"latitude=%s&longitude=%s&daily=temperature_2m_max,temperature_2m_min",
		loc.Latitude,
		loc.Longitude,
	))
	formattedUrl.WriteString(formatExtraForecastParams(forecast_params))

	response, err := http.Get(formattedUrl.String())

	if err != nil {
		return []byte{}, err
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return []byte{}, err
	}

	return responseData, nil
}

type Response struct {
	History History `json:"daily"`
}
type History struct {
	Hello []float64 `json:"temperature_2m_max"`
	World []string  `json:"time"`
}

func createPattern(n int) string {
	asterisks := strings.Repeat("*", n)
	spaces := strings.Repeat(" ", 5-n)
	return asterisks + spaces
}
func processJsonData(jsonData []byte) {
	var resp Response

	err := json.Unmarshal(jsonData, &resp)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var min = 1000000000000000.0
	var max = 0.0

	for i := 0; i < len(resp.History.Hello); i++ {
		var temp = resp.History.Hello[i]
		temp += 70.0
		if temp < min {
			min = temp
		}
		if temp > max {
			max = temp
		}
	}

	//fmt.Println(min)
	//fmt.Println(max)

	for i := 0; i < len(resp.History.Hello); i++ {
		var temp = resp.History.Hello[i] + 70
		var stars = int((temp - min) / (max - min) * 5)
		if stars == 0 {
			stars = 1
		}
		t, err := time.Parse("2006-01-02", resp.History.World[i])
		if err != nil {
			panic(err)
		}
		fmt.Println(createPattern(stars), fmt.Sprintf("%02d", int(temp)-70), "°C", resp.History.World[i], t.Weekday())

	}

}

func FindCityLocation(city City) (string, string, error) {
	api_params := url.PathEscape(fmt.Sprintf("name=%s&count=10&language=en&format=json", city.Name))
	api_url := fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?%s", api_params)
	response, err := http.Get(api_url)

	if err != nil {
		return "", "", err
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return "", "", err
	}

	var geocodingResponse GeoCodingResponse
	err = json.Unmarshal(responseData, &geocodingResponse)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	for i := 0; i < len(geocodingResponse.GeoCodingResults); i++ {
		if geocodingResponse.GeoCodingResults[i].Country == city.Country {
			log.Println("Matched country location: ", geocodingResponse.GeoCodingResults[i])
			return geocodingResponse.GeoCodingResults[i].Latitude.String(), geocodingResponse.GeoCodingResults[i].Longitude.String(), nil
		}
	}

	return "", "", fmt.Errorf("Could not find a proper location match for %s of country %s", city.Name, city.Country)
}

func main() {
	city := flag.String("city", "", "Name of the city (e.g., 'The Hague') - *Mandatory")
	country := flag.String("country", "", "Country of the city (e.g., 'Netherlands') - *Mandatory")
	day := flag.String("day", "", "Day for the weather forecast (e.g., '2024-10-31') - Optional (default is today)")
	prec := flag.Bool("p", false, "Get precipitation - Optional")
	uv := flag.Bool("uv", false, "Get UV index - Optional")
	sunrise := flag.Bool("sunrise", false, "Get sunrise time - Optional")
	sunset := flag.Bool("sunset", false, "Get sunset time - Optional")
	fahrenheit := flag.Bool("f", false, "Use fahrenheit - Optional")

	flag.Usage = func() {
		fmt.Println("Weather Forecast Tool")
		fmt.Println("Weekly weather forecast for a city.")
		fmt.Println("Usage:")
		fmt.Println("  go run main.go -city=\"CityName\" -country=\"CountryName\" [-day=\"YYYY-MM-DD\"] [-uv] [-sunrise] [-p] [-sunset] [-f]")
		fmt.Println()
		fmt.Println("Mandatory Flags:")
		fmt.Println("  -city     Name of the city (e.g., 'The Hague')")
		fmt.Println("  -country  Country of the city (e.g., 'Netherlands')")
		fmt.Println()
		fmt.Println("Optional Flags:")
		fmt.Println("  -day      Day for the weather forecast (default is today)")
		fmt.Println("  -p        Get precipitation")
		fmt.Println("  -uv       Get UV index")
		fmt.Println("  -sunrise  Get sunrise time")
		fmt.Println("  -sunset   Get sunset time")
		fmt.Println("  -f        Use fahrenheit")
	}

	flag.Parse()

	if *city == "" || *country == "" {
		flag.Usage()
		os.Exit(1)
	}

	if *day == "" {
		*day = time.Now().Format("2006-01-02")
	}

	lat, lon, err := FindCityLocation(City{Name: *city, Country: *country})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	loc := Location{
		Latitude:  lat,
		Longitude: lon,
	}
	params := ForecastParams{
		Precipitation: *prec,
		Sunrise:       *sunrise,
		Sunset:        *sunset,
		UVIndex:       *uv,
		Fahr:          *fahrenheit,
	}

	weather, err := GetWeather(loc, params)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// TODO: Print weather
	fmt.Printf("Weather in %s, %s on %s: %s\n", *city, *country, *day, string(weather))
	processJsonData(weather)
}
