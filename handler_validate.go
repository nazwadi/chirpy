package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	type returnVals struct {
		CleanedBody string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	badWords := [3]string{"kerfuffle", "sharbert", "fornax"}
	bodyParse := strings.Split(params.Body, " ")
	for idx, word := range bodyParse {
		for _, badword := range badWords {
			if strings.ToLower(word) == badword {
				bodyParse[idx] = "****"
			}
		}
	}
	body := strings.Join(bodyParse, " ")

	respondWithJSON(w, http.StatusOK, returnVals{
		CleanedBody: body,
	})
}
