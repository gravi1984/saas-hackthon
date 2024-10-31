package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// // Struct to represent the JSON data
// type Day struct {
// 	Day  string `json:"name"`
// 	Temp int    `json:"Temperature"`
// }

// type Response struct {
// 	History History `json:"properties"`
// }

// type History struct {
// 	Day []Day `json:"periods"`
// }

// // // main function
// // func main() {
// // 	// JSON data as a byte slice
// // 	jsonData := []byte(`

// {
//     "@context": [
//         "https://geojson.org/geojson-ld/geojson-context.jsonld",
//         {
//             "@version": "1.1",
//             "wx": "https://api.weather.gov/ontology#",
//             "geo": "http://www.opengis.net/ont/geosparql#",
//             "unit": "http://codes.wmo.int/common/unit/",
//             "@vocab": "https://api.weather.gov/ontology#"
//         }
//     ],
//     "type": "Feature",
//     "geometry": {
//         "type": "Polygon",
//         "coordinates": [
//             [
//                 [
//                     -97.137207000000004,
//                     39.7444372
//                 ],
//                 [
//                     -97.1367549,
//                     39.7223799
//                 ],
//                 [
//                     -97.108080900000004,
//                     39.722725199999999
//                 ],
//                 [
//                     -97.108527000000009,
//                     39.744782499999999
//                 ],
//                 [
//                     -97.137207000000004,
//                     39.7444372
//                 ]
//             ]
//         ]
//     },
//     "properties": {
//         "units": "us",
//         "forecastGenerator": "BaselineForecastGenerator",
//         "generatedAt": "2024-10-31T10:50:03+00:00",
//         "updateTime": "2024-10-31T08:11:16+00:00",
//         "validTimes": "2024-10-31T02:00:00+00:00/P6DT23H",
//         "elevation": {
//             "unitCode": "wmoUnit:m",
//             "value": 456.89519999999999
//         },
//         "periods": [
//             {
//                 "number": 1,
//                 "name": "Overnight",
//                 "startTime": "2024-10-31T05:00:00-05:00",
//                 "endTime": "2024-10-31T06:00:00-05:00",
//                 "isDaytime": false,
//                 "temperature": 39,
//                 "temperatureUnit": "F",
//                 "temperatureTrend": "",
//                 "probabilityOfPrecipitation": {
//                     "unitCode": "wmoUnit:percent",
//                     "value": null
//                 },
//                 "windSpeed": "15 mph",
//                 "windDirection": "NW",
//                 "icon": "https://api.weather.gov/icons/land/night/bkn?size=medium",
//                 "shortForecast": "Mostly Cloudy",
//                 "detailedForecast": "Mostly cloudy, with a low around 39. Northwest wind around 15 mph, with gusts as high as 25 mph."
//             },
//             {
//                 "number": 2,
//                 "name": "Thursday",
//                 "startTime": "2024-10-31T06:00:00-05:00",
//                 "endTime": "2024-10-31T18:00:00-05:00",
//                 "isDaytime": true,
//                 "temperature": 55,
//                 "temperatureUnit": "F",
//                 "temperatureTrend": "",
//                 "probabilityOfPrecipitation": {
//                     "unitCode": "wmoUnit:percent",
//                     "value": null
//                 },
//                 "windSpeed": "5 to 15 mph",
//                 "windDirection": "W",
//                 "icon": "https://api.weather.gov/icons/land/day/few?size=medium",
//                 "shortForecast": "Sunny",
//                 "detailedForecast": "Sunny, with a high near 55. West wind 5 to 15 mph, with gusts as high as 25 mph."
//             },
//             {
//                 "number": 3,
//                 "name": "Thursday Night",
//                 "startTime": "2024-10-31T18:00:00-05:00",
//                 "endTime": "2024-11-01T06:00:00-05:00",
//                 "isDaytime": false,
//                 "temperature": 37,
//                 "temperatureUnit": "F",
//                 "temperatureTrend": "",
//                 "probabilityOfPrecipitation": {
//                     "unitCode": "wmoUnit:percent",
//                     "value": null
//                 },
//                 "windSpeed": "0 to 5 mph",
//                 "windDirection": "S",
//                 "icon": "https://api.weather.gov/icons/land/night/few?size=medium",
//                 "shortForecast": "Mostly Clear",
//                 "detailedForecast": "Mostly clear, with a low around 37. South wind 0 to 5 mph."
//             },
//             {
//                 "number": 4,
//                 "name": "Friday",
//                 "startTime": "2024-11-01T06:00:00-05:00",
//                 "endTime": "2024-11-01T18:00:00-05:00",
//                 "isDaytime": true,
//                 "temperature": 60,
//                 "temperatureUnit": "F",
//                 "temperatureTrend": "",
//                 "probabilityOfPrecipitation": {
//                     "unitCode": "wmoUnit:percent",
//                     "value": null
//                 },
//                 "windSpeed": "5 to 15 mph",
//                 "windDirection": "S",
//                 "icon": "https://api.weather.gov/icons/land/day/few?size=medium",
//                 "shortForecast": "Sunny",
//                 "detailedForecast": "Sunny, with a high near 60. South wind 5 to 15 mph, with gusts as high as 25 mph."
//             },
//             {
//                 "number": 5,
//                 "name": "Friday Night",
//                 "startTime": "2024-11-01T18:00:00-05:00",
//                 "endTime": "2024-11-02T06:00:00-05:00",
//                 "isDaytime": false,
//                 "temperature": 42,
//                 "temperatureUnit": "F",
//                 "temperatureTrend": "",
//                 "probabilityOfPrecipitation": {
//                     "unitCode": "wmoUnit:percent",
//                     "value": 40
//                 },
//                 "windSpeed": "5 mph",
//                 "windDirection": "SE",
//                 "icon": "https://api.weather.gov/icons/land/night/sct/rain_showers,40?size=medium",
//                 "shortForecast": "Partly Cloudy then Chance Rain Showers",
//                 "detailedForecast": "A chance of rain showers after 1am. Partly cloudy, with a low around 42. Southeast wind around 5 mph. Chance of precipitation is 40%."
//             },
//             {
//                 "number": 6,
//                 "name": "Saturday",
//                 "startTime": "2024-11-02T06:00:00-05:00",
//                 "endTime": "2024-11-02T18:00:00-05:00",
//                 "isDaytime": true,
//                 "temperature": 61,
//                 "temperatureUnit": "F",
//                 "temperatureTrend": "",
//                 "probabilityOfPrecipitation": {
//                     "unitCode": "wmoUnit:percent",
//                     "value": 90
//                 },
//                 "windSpeed": "5 to 10 mph",
//                 "windDirection": "SE",
//                 "icon": "https://api.weather.gov/icons/land/day/tsra,70/tsra,90?size=medium",
//                 "shortForecast": "Showers And Thunderstorms",
//                 "detailedForecast": "A chance of rain showers before 7am, then showers and thunderstorms. Mostly cloudy, with a high near 61. Southeast wind 5 to 10 mph. Chance of precipitation is 90%."
//             },
//             {
//                 "number": 7,
//                 "name": "Saturday Night",
//                 "startTime": "2024-11-02T18:00:00-05:00",
//                 "endTime": "2024-11-03T06:00:00-06:00",
//                 "isDaytime": false,
//                 "temperature": 54,
//                 "temperatureUnit": "F",
//                 "temperatureTrend": "",
//                 "probabilityOfPrecipitation": {
//                     "unitCode": "wmoUnit:percent",
//                     "value": 90
//                 },
//                 "windSpeed": "10 mph",
//                 "windDirection": "SE",
//                 "icon": "https://api.weather.gov/icons/land/night/tsra,90?size=medium",
//                 "shortForecast": "Showers And Thunderstorms",
//                 "detailedForecast": "Showers and thunderstorms. Cloudy, with a low around 54. Southeast wind around 10 mph, with gusts as high as 20 mph. Chance of precipitation is 90%."
//             },
//             {
//                 "number": 8,
//                 "name": "Sunday",
//                 "startTime": "2024-11-03T06:00:00-06:00",
//                 "endTime": "2024-11-03T18:00:00-06:00",
//                 "isDaytime": true,
//                 "temperature": 68,
//                 "temperatureUnit": "F",
//                 "temperatureTrend": "",
//                 "probabilityOfPrecipitation": {
//                     "unitCode": "wmoUnit:percent",
//                     "value": 90
//                 },
//                 "windSpeed": "10 to 15 mph",
//                 "windDirection": "S",
//                 "icon": "https://api.weather.gov/icons/land/day/tsra,90/tsra,40?size=medium",
//                 "shortForecast": "Showers And Thunderstorms then Chance Showers And Thunderstorms",
//                 "detailedForecast": "Showers and thunderstorms before noon, then a chance of showers and thunderstorms. Mostly cloudy, with a high near 68. South wind 10 to 15 mph, with gusts as high as 25 mph. Chance of precipitation is 90%."
//             },
//             {
//                 "number": 9,
//                 "name": "Sunday Night",
//                 "startTime": "2024-11-03T18:00:00-06:00",
//                 "endTime": "2024-11-04T06:00:00-06:00",
//                 "isDaytime": false,
//                 "temperature": 54,
//                 "temperatureUnit": "F",
//                 "temperatureTrend": "",
//                 "probabilityOfPrecipitation": {
//                     "unitCode": "wmoUnit:percent",
//                     "value": 70
//                 },
//                 "windSpeed": "10 mph",
//                 "windDirection": "S",
//                 "icon": "https://api.weather.gov/icons/land/night/tsra,60/tsra,70?size=medium",
//                 "shortForecast": "Showers And Thunderstorms Likely",
//                 "detailedForecast": "Showers and thunderstorms likely. Mostly cloudy, with a low around 54. South wind around 10 mph, with gusts as high as 20 mph. Chance of precipitation is 70%."
//             },
//             {
//                 "number": 10,
//                 "name": "Monday",
//                 "startTime": "2024-11-04T06:00:00-06:00",
//                 "endTime": "2024-11-04T18:00:00-06:00",
//                 "isDaytime": true,
//                 "temperature": 64,
//                 "temperatureUnit": "F",
//                 "temperatureTrend": "",
//                 "probabilityOfPrecipitation": {
//                     "unitCode": "wmoUnit:percent",
//                     "value": 80
//                 },
//                 "windSpeed": "10 mph",
//                 "windDirection": "S",
//                 "icon": "https://api.weather.gov/icons/land/day/tsra,60/tsra,80?size=medium",
//                 "shortForecast": "Showers And Thunderstorms",
//                 "detailedForecast": "Showers and thunderstorms. Mostly cloudy, with a high near 64. Chance of precipitation is 80%."
//             },
//             {
//                 "number": 11,
//                 "name": "Monday Night",
//                 "startTime": "2024-11-04T18:00:00-06:00",
//                 "endTime": "2024-11-05T06:00:00-06:00",
//                 "isDaytime": false,
//                 "temperature": 40,
//                 "temperatureUnit": "F",
//                 "temperatureTrend": "",
//                 "probabilityOfPrecipitation": {
//                     "unitCode": "wmoUnit:percent",
//                     "value": 70
//                 },
//                 "windSpeed": "10 mph",
//                 "windDirection": "SW",
//                 "icon": "https://api.weather.gov/icons/land/night/tsra_sct,70/tsra_sct,30?size=medium",
//                 "shortForecast": "Showers And Thunderstorms Likely",
//                 "detailedForecast": "Showers and thunderstorms likely. Mostly cloudy, with a low around 40. Chance of precipitation is 70%."
//             },
//             {
//                 "number": 12,
//                 "name": "Tuesday",
//                 "startTime": "2024-11-05T06:00:00-06:00",
//                 "endTime": "2024-11-05T18:00:00-06:00",
//                 "isDaytime": true,
//                 "temperature": 56,
//                 "temperatureUnit": "F",
//                 "temperatureTrend": "",
//                 "probabilityOfPrecipitation": {
//                     "unitCode": "wmoUnit:percent",
//                     "value": 30
//                 },
//                 "windSpeed": "5 to 10 mph",
//                 "windDirection": "W",
//                 "icon": "https://api.weather.gov/icons/land/day/rain_showers,30/rain_showers,20?size=medium",
//                 "shortForecast": "Chance Rain Showers",
//                 "detailedForecast": "A chance of rain showers. Mostly sunny, with a high near 56. Chance of precipitation is 30%."
//             },
//             {
//                 "number": 13,
//                 "name": "Tuesday Night",
//                 "startTime": "2024-11-05T18:00:00-06:00",
//                 "endTime": "2024-11-06T06:00:00-06:00",
//                 "isDaytime": false,
//                 "temperature": 34,
//                 "temperatureUnit": "F",
//                 "temperatureTrend": "",
//                 "probabilityOfPrecipitation": {
//                     "unitCode": "wmoUnit:percent",
//                     "value": 20
//                 },
//                 "windSpeed": "5 mph",
//                 "windDirection": "W",
//                 "icon": "https://api.weather.gov/icons/land/night/rain_showers,20/sct?size=medium",
//                 "shortForecast": "Slight Chance Rain Showers then Partly Cloudy",
//                 "detailedForecast": "A slight chance of rain showers before midnight. Partly cloudy, with a low around 34. Chance of precipitation is 20%."
//             },
//             {
//                 "number": 14,
//                 "name": "Wednesday",
//                 "startTime": "2024-11-06T06:00:00-06:00",
//                 "endTime": "2024-11-06T18:00:00-06:00",
//                 "isDaytime": true,
//                 "temperature": 55,
//                 "temperatureUnit": "F",
//                 "temperatureTrend": "",
//                 "probabilityOfPrecipitation": {
//                     "unitCode": "wmoUnit:percent",
//                     "value": null
//                 },
//                 "windSpeed": "5 to 10 mph",
//                 "windDirection": "SE",
//                 "icon": "https://api.weather.gov/icons/land/day/sct?size=medium",
//                 "shortForecast": "Mostly Sunny",
//                 "detailedForecast": "Mostly sunny, with a high near 55."
//             }
//         ]
//     }
// }

//     `)

// 	// Create an instance of the struct to hold the parsed data
// 	var resp Response

// 	// Parse the JSON data into the struct
// 	err := json.Unmarshal(jsonData, &resp)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	// Access the parsed data
// 	fmt.Println("Name:", resp.History.Day)
// 	fmt.Println("Age:", resp.History.Day)

// }
