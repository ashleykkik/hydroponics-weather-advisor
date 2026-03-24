# Hydroponics Weather Advisor

**Author:** Ashley Kinyanjui

**Institution:** Moringa School — AI Essentials for Software Developers

**Date:** 24th March 2026


## Project Overview

The Hydroponics Weather Advisor is a command-line tool built in Go (Golang). It is specifically designed for urban farmers in Nairobi to help manage hydroponic systems by fetching real time weather data and providing automated nutrient solution advice.


##  Features

 **Live Data:** Fetches current temperature and humidity for Nairobi via the OpenWeatherMap API

**Smart Advisory:** Provides specific warnings for high heat, cold temperatures, and humidity spikes to protect crop health

 **Modern Backend:** Built using Go to explore static typing, efficient compilation, and robust error handling


##  Setup & Installation

### 1. Requirements
- Go (Golang) installed on your system
- A valid OpenWeatherMap API Key

### 2. Configuration (Security)
To follow professional security standards, this project uses environment variables to protect the API key.

Create a `.env` file in the root directory and add your key:
```
WEATHER_API_KEY=your_actual_key_here
```

Note: The `.gitignore` file ensures your secret key is never uploaded to the public repository.

### 3. Running the Program

**Windows (PowerShell):**
```
$env:WEATHER_API_KEY="your_api_key_here"
go run main.go
```

**Mac/Linux:**
```
export WEATHER_API_KEY="your_api_key_here"
go run main.go
```


## 💻 Expected Output
```
--- Hydroponics Weather Advisor ---
Current Temp: 17.1°C
Humidity:     93%
```


## 📊 Advisory Logic

- **High Heat (>28°C):** Advice to increase oxygenation and check reservoir temperature
- **Cold (<18°C):** Warning that plant growth may slow down
- **High Humidity (>80%):** Recommendation to increase ventilation to prevent fungal growth


##  Project Structure
```
hydroponics-weather-advisor/
├── main.go                                      # Main application code
├── go.mod                                       # Go module file
├── .gitignore                                   # Keeps .env out of GitHub
├── README.md                                    # This file
└── AshleyKinyanjui_HydroponicsToolkitFINAL.pdf # Full toolkit document
```


##  Resources

- [Go Official Docs](https://go.dev/doc/)
- [OpenWeatherMap API](https://openweathermap.org/api)
- [Full Toolkit Document](./AshleyKinyanjui_HydroponicsToolkitFINAL.pdf)
