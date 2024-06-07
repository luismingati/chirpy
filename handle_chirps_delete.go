package main

import (
	"net/http"
	"strconv"

	"github.com/luismingati/chirpy/internal/auth"
)

func (cfg *apiConfig) HandleChirpsDelete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	idInt, _ := strconv.Atoi(id)

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find JWT")
		return
	}

	subject, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't validate JWT")
		return
	}

	userIDInt, err := strconv.Atoi(subject)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't parse user ID")
		return
	}

	chirp, err := cfg.DB.GetChirpsById(userIDInt)
	if err != nil {
		respondWithError(w, http.StatusForbidden, "Couldn't get chirp")
		return
	}

	if chirp.AuthorID != userIDInt {
		respondWithError(w, http.StatusForbidden, "You can't delete this chirp")
		return
	}

	err = cfg.DB.DeleteChirp(idInt)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusNoContent, nil)

}
