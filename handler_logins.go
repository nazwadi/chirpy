package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/nazwadi/chirpy/internal/auth"
	"github.com/nazwadi/chirpy/internal/database"
)

func (cfg *apiConfig) handlerLogins(w http.ResponseWriter, r *http.Request) {
	type params struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	type returnVals struct {
		ID           uuid.UUID `json:"id"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
		Email        string    `json:"email"`
		Token        string    `json:"token"`
		RefreshToken string    `json:"refresh_token"`
	}

	decoder := json.NewDecoder(r.Body)
	parameters := params{}
	err := decoder.Decode(&parameters)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't decode parameters", err)
		return
	}

	user, err := cfg.dbQueries.GetUser(r.Context(), parameters.Email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "User not found", err)
		return
	}
	matches, err := auth.CheckPasswordHash(parameters.Password, user.HashedPassword)
	if err != nil || !matches {
		respondWithError(w, http.StatusUnauthorized, "password does not match", err)
		return
	}
	expiresIn := time.Hour
	token, err := auth.MakeJWT(user.ID, cfg.secret, expiresIn)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't generate JWT", err)
		return
	}
	rawTokenStr := auth.MakeRefreshToken()
	refreshToken, err := cfg.dbQueries.CreateToken(r.Context(), database.CreateTokenParams{Token: rawTokenStr, UserID: user.ID, ExpiresAt: time.Now().Add(60 * 24 * time.Hour)})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create refresh token", err)
		return
	}
	respondWithJSON(w, http.StatusOK, returnVals{ID: user.ID, CreatedAt: user.CreatedAt, UpdatedAt: user.UpdatedAt, Email: user.Email, Token: token, RefreshToken: refreshToken.Token})
}
