package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// Connect to Redis
var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379", // Change if Redis runs on a different host
	Password: "",               // No password by default
	DB:       0,                // Default DB
})

// Simulated function to fetch player stats from an external API
func fetchPlayerSummary(battletag string) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://overfast-api.tekrop.fr/players/%s/summary", battletag)
	log.Println("Fetching player summary data from API:", url) // Debug statement
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&data)
	return data, nil
}

func fetchPlayerStats(battletag string) (map[string]interface{}, error) {

	url := fmt.Sprintf("https://overfast-api.tekrop.fr/players/%s/stats/summary", battletag)
	fmt.Println("Fetching player stats data from API:", url)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var data map[string]interface{}
	json.NewDecoder(res.Body).Decode(&data)
	return data, nil

}

// Cache Functions

// Check if data is present in Redis Cache. If not present, return false.
// If data is present, return true along with the cached data.
func cacheRead(key string) (bool, map[string]interface{}) {
	cachedData, err := rdb.Get(ctx, key).Result()
	if err == nil {
		// Cache hit
		log.Println("Cache hit for key:", key) // Debug statement
		var cachedStats map[string]interface{}
		json.Unmarshal([]byte(cachedData), &cachedStats)
		return true, cachedStats
	}
	return false, nil
}

// Write data to Redis Cache
// If data is successfully written, set a TTL of 10 minutes. If not, log the error.
func cacheWrite(key string, data map[string]interface{}) {
	jsonData, _ := json.Marshal(data)
	err := rdb.Set(ctx, key, jsonData, 10*time.Minute).Err()
	log.Println("Storing data in Redis:", key) // Debug statement
	if err != nil {
		log.Println("Error storing data in Redis:", err) // Debug statement
	}
}

// API Endpoint: Get player stats
// battletag_summary
func getPlayerSummary(c *gin.Context) {
	battletag := c.Param("battletag")

	// 1️ Check Redis Cache
	cached, cachedData := cacheRead(battletag + "_summary")
	if cached {
		c.JSON(200, gin.H{"source": "cache", "data": cachedData})
		return
	}
	//since we are at this point in the function, we are dealing with a cache miss

	log.Println("Cache miss for battletag:", battletag) // Debug statement

	//Fetch from External API if Cache Miss
	playerStats, err := fetchPlayerSummary(battletag)
	if err != nil {
		log.Println("Error fetching player stats from API:", err) // Debug statement
		c.JSON(500, gin.H{"error": "Failed to fetch player stats"})
		return
	}

	//Store in Cache (Set TTL for 10 minutes)
	cacheWrite(battletag+"_summary", playerStats)

	// Return the Data
	c.JSON(200, gin.H{"source": "API", "data": playerStats})
}

func getPlayerStats(c *gin.Context) {
	battletag := c.Param("battletag")

	// 1️ Check Redis Cache
	// 1️ Check Redis Cache
	cached, cachedData := cacheRead(battletag + "_stats")
	if cached {
		c.JSON(200, gin.H{"source": "cache", "data": cachedData})
		return
	}

	// Cache miss
	log.Println("Cache miss for player stats:", battletag) // Debug statement

	// 2 Fetch from External API
	playerStats, err := fetchPlayerStats(battletag)
	if err != nil {
		log.Println("Error fetching player stats from API:", err) // Debug statement
		c.JSON(500, gin.H{"error": "Failed to fetch player stats"})
		return
	}

	// 3 Store in Cache
	cacheWrite(battletag+"_stats", playerStats)

	// 4 Return the Data
	c.JSON(200, gin.H{"source": "API", "data": playerStats})
}

func main() {

	r := gin.Default()
	r.GET("/api/players/:battletag", getPlayerSummary)
	r.GET("/api/players/stats/:battletag", getPlayerStats)
	r.Run(":8080") // Running on localhost:8080
}
