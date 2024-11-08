package repositories

import (
	"database/sql"
	"mcs_bab_6/entities"
)

func GetCards(db *sql.DB) (result []entities.Card, err error) {
	sql := "SELECT * FROM card"

	rows, err := db.Query(sql)

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var data entities.Card
		err = rows.Scan(&data.ID)
		if err != nil {
			return
		}
		result = append(result, data)
	}
	return
}

func InsertCard(db *sql.DB, card entities.Card) (err error) {
	sql := "INSERT INTO card(id) values($1)"

	// errs := db.QueryRow(sql, 1)
	// return errs.Err()

	_, err = db.Exec(sql, card.ID)
	return err
}

func DeleteCard(db *sql.DB, card entities.Card) (err error) {
	sql := "DELETE FROM card WHERE id = $1"

	// errs := db.QueryRow(sql, 1)
	// return errs.Err()

	_, err = db.Exec(sql, card.ID)
	return err
}
