Hydroponics Weather Advisor
Author: Ashley Kinyanjui
Institution: Moringa School — AI Essentials for Software Developers

Date: March 2026

📌 Project Overview
The Hydroponics Weather Advisor is a command-line tool built in Go (Golang). It is specifically designed for urban farmers in Nairobi to help manage hydroponic systems by fetching real-time weather data and providing automated nutrient solution advice.

🛠️ Features
Live Data: Fetches current temperature and humidity for Nairobi via the OpenWeatherMap API.

Smart Advisory: Provides specific warnings for high heat, cold temperatures, and humidity spikes to protect crop health.

Modern Backend: Built using Go to explore static typing, efficient compilation, and robust error handling.

🚀 Setup & Installation
1. Requirements
Go (Golang) installed on your system.

A valid OpenWeatherMap API Key.

2. Configuration (Security)
To follow professional security standards, this project uses Environment Variables to protect the API key.

Create a .env file in the root directory.

Add your key to the file:

WEATHER_API_KEY=your_actual_key_here

Note: The .gitignore file ensures your secret key is never uploaded to the public repository.

3. Running the Program
Open your terminal in the project folder and run:

On Windows (PowerShell):

PowerShell
$env:WEATHER_API_KEY="your_api_key_here"
go run main.go
On Mac/Linux:

Bash
export WEATHER_API_KEY="your_api_key_here"
go run main.go
📊 Advisory Logic
High Heat (>28°C): Advice to increase oxygenation and check reservoir temperature.

Cold (<18°C): Warning that plant growth may slow down.

High Humidity (>80%): Recommendation to increase ventilation to prevent fungal growth.