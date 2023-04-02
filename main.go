package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type post struct {
	UserID int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type value struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	for range time.Tick(15 * time.Second) {
		userID := rand.Intn(100) + 1
		postData := post{
			UserID: userID,
			Title:  fmt.Sprintf(" %d", userID),
			Body:   fmt.Sprintf(" %d", userID),
		}

		postJSON, err := json.Marshal(postData)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		resp, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(postJSON))
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Println("POST request sent with data:", string(postJSON))
		fmt.Println("Response:", resp.Status)

		water := rand.Intn(100) + 1
		wind := rand.Intn(100) + 1

		waterStatus := getStatus(water, "water")
		windStatus := getStatus(wind, "wind")

		statusData := value{Water: water, Wind: wind}

		statusJSON, err := json.MarshalIndent(statusData, "", "  ")
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Println(string(statusJSON))

		fmt.Printf("Status water: %s\n", waterStatus)
		fmt.Printf("Status wind: %s\n", windStatus)
	}
}

func getStatus(value int, dataType string) string {
	var statusText string

	switch dataType {
	case "water":
		if value < 5 {
			statusText = "aman"
		} else if value >= 5 && value <= 8 {
			statusText = "siaga"
		} else {
			statusText = "bahaya"
		}
	case "wind":
		if value < 6 {
			statusText = "aman"
		} else if value >= 6 && value <= 15 {
			statusText = "siaga"
		} else {
			statusText = "bahaya"
		}
	default:
		statusText = "unknown"
	}

	return statusText
}
