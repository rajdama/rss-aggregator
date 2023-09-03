package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/rajdama/rss-aggregator/internal/database"
)

/*
We make this function/handler part of apiconfig because we have to also pass reference to aur datbase in order
to interact with it and the function signature of handler i.e number of arguments and return type cannot
be changed hence we pass our refernce to database by adding this function to apiconfig struct
*/

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing json: %s", err))
		return
	}

	user, errDB := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if errDB != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user: %s", errDB))
		return
	}

	respondWithJSON(w, 200, user)
}
