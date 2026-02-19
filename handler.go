package go_auth_api_tests

import (
	"encoding/json"
	"net/http"
)

type Request struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Response struct {
	Status string `json:"status"`
}

func HandlerRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := Register(req.Email, req.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res := Response{Status: "registered"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}
