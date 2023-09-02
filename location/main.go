package main

import (
	"fmt"
	"github.com/goccy/go-json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		ip, err := getIP()
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal Server Error",
			})
			return
		}

		country, region, city, err := getLocation(ip)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal Server Error",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"ip":      ip,
			"country": country,
			"region":  region,
			"city":    city,
		})
	})

	err := router.Run(":8080")
	if err != nil {
		return
	}
}

func getIP() (string, error) {
	resp, err := http.Get("https://api.ipify.org?format=text&apiKey=YOUR_API_KEY")
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			println(err)
		}
	}(resp.Body)

	ip := ""
	_, err = fmt.Fscanf(resp.Body, "%s", &ip)
	if err != nil {
		return "", err
	}

	return ip, nil
}

func getLocation(ip string) (string, string, string, error) {
	url := fmt.Sprintf("https://ipapi.com/%s?access_key=YOUR_API_KEY", ip)
	resp, err := http.Get(url)
	if err != nil {
		return "", "", "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			println(err)
		}
	}(resp.Body)

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", "", "", err
	}

	country := data["country_name"].(string)
	region := data["region"].(string)
	city := data["city"].(string)

	return country, region, city, nil
}
