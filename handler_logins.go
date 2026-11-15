package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/nazwadi/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerLogins(w http.ResponseWriter, r *http.Request) {
	type params struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	type returnVals struct {
		ID        uuid.UUID `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Email     string    `json:"email"`
	}

	decoder := json.NewDecoder(r.Body)
	parameters := params{}
	err := decoder.Decode(&parameters)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}
	user, err := cfg.dbQueries.GetUser(r.Context(), parameters.Email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "User not found", err)
		return
	}
	matches, err := auth.CheckPasswordHash(parameters.Password, user.HashedPassword)
	if err != nil || !matches {
		fmt.Println(user.HashedPassword)
		respondWithError(w, http.StatusUnauthorized, "password does not match", err)
		return
	}
	respondWithJSON(w, http.StatusOK, returnVals{ID: user.ID, CreatedAt: user.CreatedAt, UpdatedAt: user.UpdatedAt, Email: user.Email})
}
