package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dasotd/rssagg/internal/auth"
	"github.com/dasotd/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateuser(w http.ResponseWriter, r *http.Request){
	type parameters struct{
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err :=decoder.Decode(&params)
	if err !=nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing Json %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name :     params.Name,
	},
	)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error creating user %v", err))
		return
	}
	responseWithJSON(w, 201, databaseUserToUser(user))
}

func(apiCfg apiConfig)handlerGetUser(w http.ResponseWriter, r *http.Request){
	apikey, err := auth.GetApiKey(r.Header)
	if err != nil {
		responseWithError(w, 403, fmt.Sprintf("Error getting user api key  %v", err))
		return
	}
	user, err :=apiCfg.DB.GetUserByApiKey(r.Context(), apikey)
	if err != nil {
		responseWithError(w, 403, fmt.Sprintf("Error getting user with api key  %v", err))
		return
	}

	responseWithJSON(w, 200, databaseUserToUser(user))
}