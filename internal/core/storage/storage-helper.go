package storage

import (
	"database/sql"
	"fmt"
)

func (m *SqlStorer) GetVersion() (s string) {
	s = "NULL"
	// Query for a value based on a single row.
	err := m.instance.QueryRow("SELECT @@VERSION").Scan(&s)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("bad")
		}
	}
	return s
}
