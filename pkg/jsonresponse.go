package pkg

import (
	"encoding/json"
	"log"
	"net/http"
)

type JSONResponse struct {
	Success bool        `json:"success"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func writeJSON(w http.ResponseWriter, status int, resp JSONResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Println("Ошибка:", err.Error())
		return
	}
}

func WriteSuccess(w http.ResponseWriter, status int, data interface{}) {
	writeJSON(w, status, JSONResponse{
		Success: true,
		Status:  status,
		Data:    data,
	})
}

func WriteError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, JSONResponse{
		Success: false,
		Status:  status,
		Error:   msg,
	})
}
