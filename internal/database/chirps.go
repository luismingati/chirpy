package database

import "errors"

type Chirp struct {
	ID       int    `json:"id"`
	Body     string `json:"body"`
	AuthorID int    `json:"author_id"`
}

func (db *DB) CreateChirp(body string, authorId int) (Chirp, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return Chirp{}, err
	}

	id := len(dbStructure.Chirps) + 1
	chirp := Chirp{
		ID:       id,
		Body:     body,
		AuthorID: authorId,
	}

	dbStructure.Chirps[id] = chirp

	err = db.writeDB(dbStructure)
	if err != nil {
		return Chirp{}, nil
	}

	return chirp, nil
}

func (db *DB) GetChirps() ([]Chirp, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return nil, nil
	}

	chirps := make([]Chirp, 0, len(dbStructure.Chirps))
	for _, chirp := range dbStructure.Chirps {
		chirps = append(chirps, chirp)
	}

	return chirps, nil
}

func (db *DB) GetChirpsById(id int) (*Chirp, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return nil, nil
	}

	for _, chirp := range dbStructure.Chirps {
		if chirp.ID == id {
			return &chirp, nil
		}
	}
	return nil, errors.New("Chirp not found")
}

func (db *DB) DeleteChirp(id int) error {
	dbStructure, err := db.loadDB()
	if err != nil {
		return err
	}

	if _, ok := dbStructure.Chirps[id]; !ok {
		return errors.New("Chirp not found")
	}

	delete(dbStructure.Chirps, id)

	err = db.writeDB(dbStructure)
	if err != nil {
		return err
	}
	db.resetChirpIds()

	return nil
}

func (db *DB) resetChirpIds() error {
	dbStructure, err := db.loadDB()
	if err != nil {
		return err
	}

	chirps := dbStructure.Chirps
	dbStructure.Chirps = make(map[int]Chirp)

	for i, chirp := range chirps {
		chirp.ID = i + 1
		dbStructure.Chirps[i+1] = chirp
	}

	err = db.writeDB(dbStructure)
	if err != nil {
		return err
	}

	return nil
}
