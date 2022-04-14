package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/wahyuadepratama/weather/src/models"

	"github.com/gin-gonic/gin"
)

func ShowAllWeather(context *gin.Context) {

	type Response struct {
		models.Weather
		WeatherDetail []models.WeatherDetail `json:"weather"`
	}

	// Show all weather data
	var weather models.Weather
	weatherData := db.Raw("select * from weather order by created_at desc limit 1").Scan(&weather)
	if weatherData.Error != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  "401",
			"message": "Load data failed",
		})
		return
	}

	// Show all weather data
	var weatherDetail []models.WeatherDetail
	weatherDetailData := db.Raw("select * from weather_detail where weather_id = ?", weather.ID).Scan(&weatherDetail)
	if weatherDetailData.Error != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  "401",
			"message": "Load data failed",
		})
		return
	}

	var response Response
	response.Weather = weather
	response.WeatherDetail = weatherDetail

	// Creating http response
	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Data loaded successfully",
		"data":    response,
	})

}

func UpdateWeather(context *gin.Context) {

	// Token validation
	if IsTokenValid(context) != 1 {
		return
	}

	// Get URL API
	url := os.Getenv("API_WEATHER")

	// Define api structure
	type WeatherDetail struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
	}

	type Current struct {
		Pressure  int             `json:"pressure"`
		Humidity  int             `json:"humidity"`
		WindSpeed float32         `json:"wind_speed"`
		Weather   []WeatherDetail `json:"weather"`
	}

	type Weather struct {
		Lat      float32 `json:"lat"`
		Lon      float32 `json:"lon"`
		Timezone string  `json:"timezone"`
		Current  Current `json:"current"`
	}

	// Define max timeout error
	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	// Validation API
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
		context.JSON(http.StatusRequestTimeout, gin.H{
			"status":  "408",
			"message": err,
		})
		return
	}

	req.Header.Set("User-Agent", "Weather Data")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
		context.JSON(http.StatusRequestTimeout, gin.H{
			"status":  "408",
			"message": getErr,
		})
		return
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
		context.JSON(http.StatusOK, gin.H{
			"status":  "401",
			"message": readErr,
		})
		return
	}

	data := Weather{}
	jsonErr := json.Unmarshal(body, &data)
	if jsonErr != nil {
		log.Fatal(jsonErr)
		context.JSON(http.StatusOK, gin.H{
			"status":  "401",
			"message": jsonErr,
		})
		return
	}

	// Insert data to weather table
	errInserWeather := db.Exec("INSERT INTO weather (lat, lon, timezone, pressure, humidity, wind_speed) VALUES (?,?,?,?,?,?)",
		data.Lat, data.Lon, data.Timezone, data.Current.Pressure, data.Current.Humidity, data.Current.WindSpeed)

	if errInserWeather.Error != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  "401",
			"message": "Insert data failed",
		})
		return
	}

	// Select latest data
	var latestWeather models.Weather
	weatherData := db.Raw("select * from weather order by created_at desc limit 1").Scan(&latestWeather)
	if weatherData.Error != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  "401",
			"message": "Load data failed",
		})
		return
	}

	for _, detail := range data.Current.Weather {
		// Insert data to weather_detail table
		errInserWeatherDetail := db.Exec("INSERT INTO weather_detail (id, weather_id, main, description) VALUES (?,?,?,?)",
			detail.ID, latestWeather.ID, detail.Main, detail.Description)

		if errInserWeatherDetail.Error != nil {
			context.JSON(http.StatusOK, gin.H{
				"status":  "401",
				"message": "Insert data failed",
			})
			return
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Data updated successfully from API",
		"data":    data,
	})
}
