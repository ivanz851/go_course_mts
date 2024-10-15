package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	baseURL       = "http://localhost:8080"
	clientTimeout = 60 * time.Second
	hardOpTimeout = 15 * time.Second
)

func main() {
	client := &http.Client{
		Timeout: clientTimeout,
	}

	version, err := getVersion(client)
	if err != nil {
		log.Fatalf("Error getting version: %v", err)
	}
	fmt.Println("Version:", version)

	inputString := "aGVsbG8gd29ybGQ="
	decodedString, err := decodeString(client, inputString)
	if err != nil {
		log.Fatalf("Error decoding string: %v", err)
	}
	fmt.Println("Decoded string:", decodedString)

	err = performHardOp(client)
	if err != nil {
		log.Printf("Hard operation timed out or failed: %v", err)
	}
}

func getVersion(client *http.Client) (string, error) {
	resp, err := client.Get(baseURL + "/version")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func decodeString(client *http.Client, inputString string) (string, error) {
	reqBody, _ := json.Marshal(map[string]string{
		"inputString": inputString,
	})

	resp, err := client.Post(baseURL+"/decode", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var decodeResp map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&decodeResp); err != nil {
		return "", err
	}

	return decodeResp["outputString"], nil
}

func performHardOp(client *http.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), hardOpTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", baseURL+"/hard-op", nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Printf("Hard operation completed with status: %d\n", resp.StatusCode)
	return nil
}
