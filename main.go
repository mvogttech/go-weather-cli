package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"unicode/utf8"

	"github.com/Jeffail/gabs"
	"github.com/olekukonko/tablewriter"
)


type Zip struct {
	Result []struct {
		AddressMatches []struct {
			Coordinates struct {
				X float64 `json:"x"`
				Y float64 `json:"y"`
			} `json:"coordinates"`
		} `json:"addressMatches"`
	} `json:"result"`
}

type GeoCoords struct {
	X string
	Y string
}

type Forecast struct {
	Periods []Period  `json:"periods"`
}

type Period struct {
		Number float64 `json:"number"`
		Name string `json:"name"`
		StartTime string `json:"startTime"`
		EndTime string `json:"endTime"`
		IsDaytime bool `json:"isDaytime"`
		Temperature float64 `json:"temperature"`
		TemperatureUnit string `json:"temperatureUnit"`
		WindSpeed string `json:"windSpeed"`
		WindDirection string `json:"windDirection"`
		Icon string `json:"icon"`
		ShortForecast string `json:"shortForecast"`
		DetailedForecast string `json:"detailedForecast"`
}

func trimLastInput(i string) string {
    r, s := utf8.DecodeLastRuneInString(i)
    if r == utf8.RuneError && (s == 0 || s == 2) {
        s = 0
    }
    return i[:len(i)-s]
}

func trimFirstInput(s string) string {
    _, i := utf8.DecodeRuneInString(s)
    return s[i:]
}

func formatStrings(address string) string {
	address = trimFirstInput(address)
	address = trimLastInput(address)
	return address
}

func getGeo(address string) GeoCoords {
	USGeoEndpoint := "https://geocoding.geo.census.gov/geocoder/locations/onelineaddress?address="
	USGeoParams := "&benchmark=2020&format=json"
	zip := trimLastInput(address)
	getURL := USGeoEndpoint + url.QueryEscape(zip) + USGeoParams
	res, _ := http.Get(getURL)
	
    body, _ := ioutil.ReadAll(res.Body)

	data, _ := gabs.ParseJSON(body)

	return GeoCoords{data.Path("result.addressMatches.coordinates.x").String(), data.Path("result.addressMatches.coordinates.y").String()}
}

func getWeatherZone(longitude string, latitude string) string {
	weatherEndpoint := "https://api.weather.gov/points/"
	getURL := weatherEndpoint + longitude + "," + latitude
	res, err := http.Get(getURL)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	data, err := gabs.ParseJSON(body)
	if err != nil {
		fmt.Println(err)
	}

	return data.Path("properties.forecast").String()
}

func getForecast(zoneURL string) Forecast {
	res, err := http.Get(formatStrings(zoneURL))
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	data, err := gabs.ParseJSON(body)
	if err != nil {
		fmt.Println(err)
	}

	result, err := data.Path("properties.periods").Children()
	if err != nil {
		fmt.Println(err)
	}

	forecast := Forecast{}

	for _, v := range result {
		forecast.Periods = append(forecast.Periods, Period{
			Number: v.Path("number").Data().(float64),
			Name: v.Path("name").Data().(string),
			StartTime: v.Path("startTime").Data().(string),
			EndTime: v.Path("endTime").Data().(string),
			IsDaytime: v.Path("isDaytime").Data().(bool),
			Temperature: v.Path("temperature").Data().(float64),
			TemperatureUnit: v.Path("temperatureUnit").Data().(string),
			WindSpeed: v.Path("windSpeed").Data().(string),
			WindDirection: v.Path("windDirection").Data().(string),
			Icon: v.Path("icon").Data().(string),
			ShortForecast: v.Path("shortForecast").Data().(string),
			DetailedForecast: v.Path("detailedForecast").Data().(string),
		})
	}

	return forecast
}

func FetchWeather(address string) {
	geo := getGeo(address)
	zone := getWeatherZone(formatStrings(geo.Y), formatStrings(geo.X))
	forecast := getForecast(zone)

	data := [][]string{}

	for _, v := range forecast.Periods {
		data = append(data, []string{v.Name, fmt.Sprint(v.Temperature)  + " " + v.TemperatureUnit, v.WindDirection + " " + v.WindSpeed, v.ShortForecast})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Day", "Temperature", "Wind", "Summary"})
	table.SetBorder(true)
	table.SetRowLine(true) 
	table.AppendBulk(data)
	fmt.Println("The weather for " + address)
	table.Render()
}


func main() {
	var address string
	if len(os.Args[1:]) > 0 {
		address = os.Args[1]
	}
	if address != "" {
		fmt.Println("Fetching weather for: " + address)
	}
	if address == "" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter your address: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		address = input
	}
	
	FetchWeather(address)

}