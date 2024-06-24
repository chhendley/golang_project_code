// https://www.youtube.com/watch?v=ipDBFTj7R4Y 3:29 mark

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type apiConfigData struct {
	OpenWeatherMapApiKey string `json:"API_KEY"`
}
type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"` // Closing curly brace for Main struct
}

func loadApiConfig(filename string) (apiConfigData, error) {
	bytes, err := os.ReadFile(filename)

	if err != nil {
		return apiConfigData{}, err
	}

	var c apiConfigData

	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return apiConfigData{}, err
	}
	return c, nil
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from go!\n"))
}

func query(city string) (weatherData, error) {
	apiConfig, err := loadApiConfig(".apiConfig")
	if err != nil {
		return weatherData{}, err
	}

	fmt.Println("Constructed URL:", "http://api.openweathermap.org/data/2.5/weather?APPID="+apiConfig.OpenWeatherMapApiKey+"&q="+city)

	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=" + apiConfig.OpenWeatherMapApiKey + "&q=" + city)
	if err != nil {
		return weatherData{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return weatherData{}, fmt.Errorf("API error: %s", resp.Status)
	}

	var d weatherData
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return weatherData{}, err
	}

	// Use the decoded data in d here (e.g., access d.Name, d.Main.Kelvin)
	// You can add your logic to process the weather data
	fmt.Println("City:", d.Name)                        // Example: Print the city name
	fmt.Println("Temperature (Kelvin):", d.Main.Kelvin) // Example: Print the temperature

	// You can also return the decoded data from here:
	return d, nil
}

func main() {
	http.HandleFunc("/hello", hello)

	http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
		city := strings.SplitN(r.URL.Path, "/", 3)[2]
		data, err := query(city)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content.Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(data)
	})

	http.ListenAndServe(":8080", nil)
}
