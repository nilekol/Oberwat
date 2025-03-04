package main

/*

package main

import (
	"context"
	"encoding/json"
	"fmt"
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
func fetchPlayerStatsFromAPI(battletag string) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://some-overwatch-api.com/player/stats?battletag=%s", battletag)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&data)
	return data, nil
}

// API Endpoint: Get player stats
func getPlayerStats(c *gin.Context) {
	battletag := c.Param("battletag")

	// 1️⃣ Check Redis Cache
	cachedData, err := rdb.Get(ctx, battletag).Result()
	if err == nil {
		// Cache hit 🎯
		var cachedStats map[string]interface{}
		json.Unmarshal([]byte(cachedData), &cachedStats)
		c.JSON(200, gin.H{"source": "cache", "data": cachedStats})
		return
	}

	// 2️⃣ Fetch from External API if Cache Miss
	playerStats, err := fetchPlayerStatsFromAPI(battletag)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch player stats"})
		return
	}

	// 3️⃣ Store in Cache (Set TTL for 10 minutes)
	jsonData, _ := json.Marshal(playerStats)
	rdb.Set(ctx, battletag, jsonData, 10*time.Minute)

	// 4️⃣ Return the Data
	c.JSON(200, gin.H{"source": "API", "data": playerStats})
}

func main() {
	r := gin.Default()
	r.GET("/api/players/:battletag", getPlayerStats)
	r.Run(":8080") // Run on localhost:8080
}

*/
