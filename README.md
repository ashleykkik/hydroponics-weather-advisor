# Hydroponics Weather Advisor

## Project Overview
This project is a Go-based toolkit designed for hydroponic farmers in Nairobi. It connects to the OpenWeatherMap API to fetch real time weather data and provides actionable advice for nutrient solution management based on the current temperature.

## Features
- **Real-time Data:** Fetches live temperature from Nairobi using a REST API.
- **Automated Advice:** Provides specific alerts if temperatures are too high or too low for optimal plant growth.
- **Lightweight Logic:** Built with Go for high performance and efficiency.

## Setup Instructions
1. **Initialize the module:** `go mod init hydroponics_advisor`
2. **Add your API Key:** Replace the `apiKey` variable in `main.go` with a valid OpenWeatherMap key.
3. **Run the application:**
   `go run main.go`

## Technical Stack
- **Language:** Go 1.26.1
- **API:** OpenWeatherMap API (One Call 3.0)
- **Environment:** Windows / VS Code

## Author
**Ashley Kinyanjui** *Moringa School: AI Essentials for Software Developers Capstone Project*