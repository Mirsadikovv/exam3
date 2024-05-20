package pkg

import (
	"database/sql"
)

func NullStringToString(s sql.NullString) string {
	if s.Valid {
		return s.String
	}

	return ""
}

func NullTimeToString(s sql.NullTime) string {
	if s.Valid {
		return s.Time.String()
	}

	return ""
}
