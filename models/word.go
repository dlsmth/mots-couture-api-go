package models

import "mc.com/rest-api/db"

// type Livre struct {
// 	Book string
// 	Page int64
// }

type Word struct {
	ID       int64
	Mot      string `binding:"required"`
	Livre    string `binding:"required"`
	Page     string `binding:"required"`
	DateTime string `binding:"required"`
}

var words = []Word{}

func (w *Word) Save() error {
	query := `
	INSERT INTO words(mot, livre, page, dateTime)
	VALUES(?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(w.Mot, w.Livre, w.Page, w.DateTime)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	w.ID = id
	// events = append(events, e)
	return nil
}

func GetAllWords() ([]Word, error) {
	query := "SELECT * FROM words"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var words []Word
	for rows.Next() {
		var word Word
		err := rows.Scan(&word.ID, &word.Mot, &word.Livre, &word.Page, &word.DateTime)
		if err != nil {
			return nil, err
		}
		words = append(words, word)
	}
	return words, nil
}

func GetWordById(id int64) (*Word, error) {
	query := "SELECT * FROM words WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var word Word
	err := row.Scan(&word.ID, &word.Mot, &word.Livre, &word.Page, &word.DateTime)
	if err != nil {
		return nil, err
	}
	return &word, nil
}

func (word Word) Update() error {
	query := `
	UPDATE words
	SET mot = ?, livre = ?, page = ?, dateTime = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(word.Mot, word.Livre, word.Page, word.DateTime, word.ID)
	return err
}

func (word Word) Delete() error {
	query := `
	DELETE FROM words
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(word.ID)
	return err
}
