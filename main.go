package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// weatherData matches the JSON structure from OpenWeatherMap
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
	// Setup with the CORRECTED API key (Starts with 'd63d...')
	apiKey := "d63d9351ed14d2531571b2c61bc1461e"
	city := "Nairobi"
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)

	fmt.Println("--- Hydroponics Weather Advisor ---")
	fmt.Printf("Fetching current conditions for %s...\n", city)

	// Make the Request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error: Could not connect to the weather service.", err)
		return
	}
	defer resp.Body.Close()

	// Security & Activation Check
	// This prevents the "panic" crash if the key is wrong
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("\n[!] API Status Error: %s (Code: %d)\n", resp.Status, resp.StatusCode)
		return
	}

	// Decode the Data
	var data weatherData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println("Error: Could not read weather data.", err)
		return
	}

	// Output & Hydroponic Advice Logic
	fmt.Println("-----------------------------------")
	fmt.Printf("Current Temp: %.1f°C\n", data.Main.Temp)
	fmt.Printf("Humidity:     %.0f%%\n", data.Main.Humidity)
	fmt.Printf("Conditions:   %s\n", data.Weather[0].Description)
	fmt.Println("-----------------------------------")

	fmt.Println("ADVISORY:")
	if data.Main.Temp > 28 {
		fmt.Println("- High Heat: Increase nutrient solution oxygenation.")
		fmt.Println("- Check reservoir temperature to prevent root stress.")
	} else if data.Main.Temp < 18 {
		fmt.Println("- Cool Temp: Growth may slow. Ensure nutrient solution isn't too cold.")
	} else {
		fmt.Println("- Optimal Temperature: Standard nutrient schedule recommended.")
	}

	if data.Main.Humidity > 80 {
		fmt.Println("- High Humidity: Increase ventilation to prevent fungal growth.")
	}
}
