package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	// Your OpenWeather API Key from your dashboard
	apiKey := "dd63d9351ed14d2531571b2c61bc1461e"
	city := "Nairobi"
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)

	fmt.Println("--- Hydroponics Weather Advisor ---")
	fmt.Printf("Fetching current conditions for %s...\n", city)

	// Making the API call to get live weather
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error: Could not connect to weather service.")
		return
	}
	defer resp.Body.Close()

	// Reading the results from the weather station
	var data map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&data)

	// Extracting the temperature from that messy text we saw in the browser
	mainData := data["main"].(map[string]interface{})
	temp := mainData["temp"].(float64)

	fmt.Printf("Current Temperature: %.2f°C\n", temp)

	// Hydroponic Advice Logic
	fmt.Println("-----------------------------------")
	if temp > 28 {
		fmt.Println("ADVICE: It's quite hot! Increase oxygenation in your nutrient solution.")
	} else if temp < 15 {
		fmt.Println("ADVICE: It's a bit cold. Monitor your plant growth rate.")
	} else {
		fmt.Println("ADVICE: Conditions are optimal for your hydroponic system. Happy growing!")
	}
	fmt.Println("-----------------------------------")
}
