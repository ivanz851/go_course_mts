package server

import (
	"encoding/base64"
	"encoding/json"
	"math/rand"
	"net/http"
	"task2/config"
	"task2/internal/models"
	"time"
)

func VersionHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(config.Version))
	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
		writer.Write([]byte("Method not allowed"))
	}
}

func DecodeHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		var req models.DecodeRequest
		if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
			http.Error(writer, "Bad request", http.StatusBadRequest)
			return
		}

		decoded, err := base64.StdEncoding.DecodeString(req.InputString)
		if err != nil {
			http.Error(writer, "Failed to decode base64 string", http.StatusInternalServerError)
			return
		}

		resp := models.DecodeResponse{
			OutputString: string(decoded),
		}

		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(resp)
	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
		writer.Write([]byte("Method not allowed"))
	}
}

func HardOpHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		sleepTime := rand.Intn(16) + 18
		time.Sleep(time.Duration(sleepTime) * time.Second)

		status := http.StatusOK
		if rand.Intn(2) == 0 {
			status = http.StatusInternalServerError
		}

		writer.WriteHeader(status)
	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
		writer.Write([]byte("Method not allowed"))
	}
}
