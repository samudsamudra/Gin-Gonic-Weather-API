package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	// Add the necessary import for database connection
	"weather-api/config"
)

const weatherAPIKey = "130e66833c4919c1423bbdc1322c75c8"

func GetWeather(c *gin.Context) {
	location := c.Query("location")
	if location == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Location is required"})
		return
	}

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", location, weatherAPIKey)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch weather data"})
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read weather data"})
		return
	}

	var weatherData map[string]interface{}
	if err := json.Unmarshal(body, &weatherData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse weather data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"location": location,
		"weather":  weatherData["weather"],
		"main":     weatherData["main"],
	})
}

func AddFavoriteLocation(c *gin.Context) {
	userID := c.GetInt("user_id") // Ambil user_id dari middleware

	type FavoriteInput struct {
		Location string `json:"location"`
	}

	var input FavoriteInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := config.DB.Exec("INSERT INTO favorites (user_id, location) VALUES (?, ?)", userID, input.Location)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add favorite location"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Favorite location added"})
}

func GetFavoriteLocations(c *gin.Context) {
	userID := c.GetInt("user_id") // Ambil user_id dari middleware

	rows, err := config.DB.Query("SELECT id, location FROM favorites WHERE user_id = ?", userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch favorite locations"})
		return
	}
	defer rows.Close()

	var favorites []map[string]interface{}
	for rows.Next() {
		var id int
		var location string
		rows.Scan(&id, &location)

		favorites = append(favorites, map[string]interface{}{
			"id":       id,
			"location": location,
		})
	}

	c.JSON(http.StatusOK, gin.H{"favorites": favorites})
}
