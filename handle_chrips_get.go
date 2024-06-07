package main

import (
	"net/http"
	"sort"
	"strconv"
)

func (cfg *apiConfig) handleChirpsRetrieve(w http.ResponseWriter, r *http.Request) {
	authorId := r.URL.Query().Get("author_id")
	sortOrder := r.URL.Query().Get("sort")

	if authorId != "" {
		authorID, err := strconv.Atoi(authorId)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid author_id")
			return
		}

		dbChirps, err := cfg.DB.GetChirpsByAuthor(authorID)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Could not retrieve chirps")
			return
		}

		chirps := []Chirp{}
		if sortOrder == "desc" {
			sort.Slice(chirps, func(i, j int) bool {
				return chirps[i].ID > chirps[j].ID
			})
		} else {
			sort.Slice(chirps, func(i, j int) bool {
				return chirps[i].ID < chirps[j].ID
			})
		}

		for _, dbChirp := range dbChirps {
			chirps = append(chirps, Chirp{
				ID:       dbChirp.ID,
				Body:     dbChirp.Body,
				AuthorID: dbChirp.AuthorID,
			})
		}

		respondWithJSON(w, http.StatusOK, chirps)
		return
	}

	dbChirps, err := cfg.DB.GetChirps()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not retrieve chirps")
		return
	}

	chirps := []Chirp{}
	if sortOrder == "desc" {
		sort.Slice(chirps, func(i, j int) bool {
			return chirps[i].ID > chirps[j].ID
		})
	} else {
		sort.Slice(chirps, func(i, j int) bool {
			return chirps[i].ID < chirps[j].ID
		})
	}
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
