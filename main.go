package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os" // Important: This allows Go to read your PC's environment
)

type weatherData struct {
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity float64 `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

func main() {
	// Reads the key from your PC's environment variables
	apiKey := os.Getenv("WEATHER_API_KEY")

	if apiKey == "" {
		fmt.Println("[!] Error: API key not found. Please set the WEATHER_API_KEY.")
		return
	}

	city := "Nairobi"
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)

	fmt.Println("--- Hydroponics Weather Advisor ---")
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	var data weatherData
	json.NewDecoder(resp.Body).Decode(&data)

	fmt.Printf("Current Temp: %.1f°C\n", data.Main.Temp)
	fmt.Printf("Humidity:     %.0f%%\n", data.Main.Humidity)
}
