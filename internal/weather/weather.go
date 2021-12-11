package weather

import (
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

func errorCheck(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func trimLastInput(s string) string {
    r, size := utf8.DecodeLastRuneInString(s)
    if r == utf8.RuneError && (size == 0 || size == 2) {
        size = 0
    }
    return s[:len(s)-size]
}

func trimFirstInput(s string) string {
    _, i := utf8.DecodeRuneInString(s)
    return s[i:]
}

func formatGeoProp(prop string) string {
	prop = trimFirstInput(prop)
	prop = trimLastInput(prop)
	return prop
}

func getGeo(address string) GeoCoords {
	USGeoEndpoint := "https://geocoding.geo.census.gov/geocoder/locations/onelineaddress?address="
	USGeoParams := "&benchmark=2020&format=json"
	zip := trimLastInput(address)
	getURL := USGeoEndpoint + url.QueryEscape(zip) + USGeoParams
	res, err := http.Get(getURL)
	
	errorCheck(err)

    body, err := ioutil.ReadAll(res.Body)
	errorCheck(err)

	data, err := gabs.ParseJSON(body)
	errorCheck(err)

	return GeoCoords{data.Path("result.addressMatches.coordinates.x").String(), data.Path("result.addressMatches.coordinates.y").String()}
}

func getWeatherZone(longitude string, latitude string) string {
	weatherEndpoint := "https://api.weather.gov/points/"
	getURL := weatherEndpoint + longitude + "," + latitude
	res, err := http.Get(getURL)
	errorCheck(err)

	body, err := ioutil.ReadAll(res.Body)
	errorCheck(err)

	data, err := gabs.ParseJSON(body)
	errorCheck(err)

	return data.Path("properties.forecast").String()
}

func getForecast(zoneURL string) Forecast {

	res, err := http.Get(formatGeoProp(zoneURL))
	errorCheck(err)

	body, err := ioutil.ReadAll(res.Body)
	errorCheck(err)

	data, err := gabs.ParseJSON(body)
	errorCheck(err)

	result, err := data.Path("properties.periods").Children()
	errorCheck(err)

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

func fetchWeather(address string) {
	geo := getGeo(address)
	zone := getWeatherZone(formatGeoProp(geo.Y), formatGeoProp(geo.X))
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