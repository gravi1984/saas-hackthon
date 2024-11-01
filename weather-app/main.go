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
	formattedParams.WriteString("&daily=temperature_2m_max,temperature_2m_min")
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
		"latitude=%s&longitude=%s",
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
	MaxTemps []float64 `json:"temperature_2m_max"`
	MinTemps []float64 `json:"temperature_2m_min"`
	UVIndex  []float64 `json:"uv_index_max"`
	Sunrise  []string  `json:"sunrise"`
	Sunset   []string  `json:"sunset"`
	Precip   []float64 `json:"precipitation_sum"`
	World    []string  `json:"time"`
}

func createPattern(n int) string {
	asterisks := strings.Repeat("*", n)
	spaces := strings.Repeat(" ", 5-n)
	return asterisks + spaces
}

func processJsonData(jsonData []byte, fah bool, showPrecip bool, showUV bool) {
	var resp Response

	err := json.Unmarshal(jsonData, &resp)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var minTemp, maxTemp float64
	for _, temp := range resp.History.MaxTemps {
		if minTemp == 0 || temp < minTemp {
			minTemp = temp
		}
		if temp > maxTemp {
			maxTemp = temp
		}
	}

	for i := 0; i < len(resp.History.MaxTemps); i++ {
		var temp float64
		if fah {
			temp = (resp.History.MaxTemps[i] * 9 / 5) + 32 // Convert to Fahrenheit
		} else {
			temp = resp.History.MaxTemps[i]
		}

		stars := int(((temp - minTemp) / (maxTemp - minTemp)) * 5)
		if stars <= 0 {
			stars = 1
		}

		sunrise := "_"
		if len(resp.History.Sunrise) > 0 {
			sunrise = resp.History.Sunrise[i]
		}

		sunset := "_"
		if len(resp.History.Sunset) > 0 {
			sunset = resp.History.Sunset[i]
		}

		var precipitation string
		if showPrecip {
			if len(resp.History.Precip) > 0 {
				precipitation = fmt.Sprintf("Precip: %.2f mm", resp.History.Precip[i])
			} else {
				precipitation = "Precip: _"
			}
		} else {
			precipitation = "Precip: _"
		}

		var uvIndex string
		if showUV {
			if len(resp.History.UVIndex) > 0 {
				uvIndex = fmt.Sprintf("UV Index: %.1f", resp.History.UVIndex[i])
			} else {
				uvIndex = "UV Index: _"
			}
		} else {
			uvIndex = "UV Index: _"
		}

		if fah {
			fmt.Printf("%s %02d °F | %s | %s | %s | %s | %s\n",
				createPattern(stars), int(temp), resp.History.World[i], sunrise, sunset, precipitation, uvIndex)
		} else {
			fmt.Printf("%s %02d °C | %s | %s | %s | %s | %s\n",
				createPattern(stars), int(temp), resp.History.World[i], sunrise, sunset, precipitation, uvIndex)
		}
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
			return geocodingResponse.GeoCodingResults[i].Latitude.String(), geocodingResponse.GeoCodingResults[i].Longitude.String(), nil
		}
	}

	return "", "", fmt.Errorf("Could not find a proper location match for %s of country %s", city.Name, city.Country)
}

func main() {
	city := flag.String("city", "", "Name of the city (e.g., 'The Hague') - *Mandatory")
	country := flag.String("country", "", "Country of the city (e.g., 'Netherlands') - *Mandatory")
	prec := flag.Bool("p", false, "Get precipitation - Optional")
	uv := flag.Bool("uv", false, "Get UV index - Optional")
	sunrise := flag.Bool("sunrise", false, "Get sunrise time - Optional")
	sunset := flag.Bool("sunset", false, "Get sunset time - Optional")
	fahrenheit := flag.Bool("f", false, "Use fahrenheit - Optional")

	flag.Usage = func() {
		fmt.Println("Weather Forecast Tool")
		fmt.Println("Weekly weather forecast for a city.")
		fmt.Println("Usage:")
		fmt.Println()
		fmt.Println("Mandatory Flags:")
		fmt.Println("  -city     Name of the city (e.g., 'The Hague')")
		fmt.Println("  -country  Country of the city (e.g., 'Netherlands')")
		fmt.Println()
		fmt.Println("Optional Flags:")
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

	processJsonData(weather, *fahrenheit, *prec, *uv)
}
