package main

import (
	"net/http"
	"sort"
	"strconv"
)

func (cfg *apiConfig) handleChirpsRetrieve(w http.ResponseWriter, r *http.Request) {
	dbChirps, err := cfg.DB.GetChirps()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not retrieve chirps")
		return
	}

	chirps := []Chirp{}
	for _, dbChirp := range dbChirps {
		chirps = append(chirps, Chirp{
			ID:       dbChirp.ID,
			Body:     dbChirp.Body,
			AuthorID: dbChirp.AuthorID,
		})
	}

	sort.Slice(chirps, func(i, j int) bool {
		return chirps[i].ID < chirps[j].ID
	})

	respondWithJSON(w, http.StatusOK, chirps)
}

func (cfg *apiConfig) handleChirpsRetrieveById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "ID not found")
	}

	chirp, err := cfg.DB.GetChirpsById(intId)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Chirp not found")
	}

	respondWithJSON(w, http.StatusOK, chirp)
}
