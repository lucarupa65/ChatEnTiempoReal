package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/luisk6510/ChatEnTiempoReal/server"
)

type HomeResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func HomeHandlers(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(HomeResponse{
			Message: "Hola mundo",
			Status:  true,
		})
	}
}
