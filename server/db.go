package server

import (
	"database/sql" 
	_ "github.com/mattn/go-sqlite3"
	"../common"
)

func GetData(timestampLowBound int) ([]common.Data, error) {
	db, err := getDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return getData(db, timestampLowBound)
}

func AddData(data common.Data) error {
	db, err := getDb()
	if err != nil {
		return err
	}
	defer db.Close()
	return addData(db, data)
}

func getData(db *sql.DB, timestamp int) ([]common.Data, error) {
	rows, err := db.Query(
		"SELECT * FROM data WHERE timestamp>=? ORDER BY timestamp DESC",
		timestamp)

	defer rows.Close()
	if err != nil {
		return nil, err
	}

	result := make([]common.Data, 0)
	for rows.Next() {
		var t common.Data
		err = rows.Scan(&t.Id, &t.Temperature, &t.Timestamp)
		if err != nil {
			return nil, err
		}
		result = append(result, t)
	}
	return result, nil
}

func addData(db *sql.DB, data common.Data) error {
	stmt, err := db.Prepare("INSERT INTO data(id, temperature, timestamp) values(?,?,?)")
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(data.Id, data.Temperature, data.Timestamp)
	return err
}

func getDb() (*sql.DB, error) {
	return sql.Open("sqlite3", "foggo.db")
}

