package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/PM-Jab/RSS-aggregator/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	users, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:       uuid.New(),
		CreateAt: time.Now().UTC(),
		UpdateAt: time.Now().UTC(),
		Name:     params.Name,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error creating user: %v", err))
		return
	}

	responseWithJSON(w, 200, databaseUserToUser(users))
}
