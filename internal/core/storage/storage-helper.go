package storage

import (
	"database/sql"
	"discordbot/internal/logger"
)

func (m *SqlStorer) GetVersion() (s string) {
	s = "NULL"
	// Query for a value based on a single row.
	err := m.instance.QueryRow("SELECT @@VERSION").Scan(&s)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.LOGERR("GetVersion - Query Returned No Rows")
		}
	}
	return s
}
