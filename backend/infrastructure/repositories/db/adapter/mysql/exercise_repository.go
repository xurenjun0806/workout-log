package mysql

import "database/sql"

type ExerciseRepository struct {
	Conn *sql.DB
}
