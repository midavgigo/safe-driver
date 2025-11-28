package dbentry

import (
	"app/utils"
	"database/sql"

	_ "github.com/lib/pq"
)

type DBManager struct {
	db *sql.DB
}

func New() (DBManager, error) {
	connStr := "postgresql://doc:password@db:5432/db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return DBManager{nil}, utils.ReasonableError{
			Reason:  err,
			Message: "Error in opening sql connection",
		}
	}

	return DBManager{db}, nil
}

func (dbman DBManager) Close() error {
	return utils.ReasonableError{
		Reason:  dbman.db.Close(),
		Message: "Error in closing sql connection",
	}
}

func (dbman DBManager) Ping() bool {
	return dbman.db.Ping() == nil
}

func (dbman DBManager) Exec(command string, args ...any) (sql.Result, error) {
	res, err := dbman.db.Exec(command, args...)
	if err != nil {
		return nil, utils.ReasonableError{
			Reason:  err,
			Message: "Error in executing command'" + command + "'",
		}
	}
	return res, nil
}

func (dbman DBManager) Query(command string, args ...any) (*sql.Rows, error) {
	rows, err := dbman.db.Query(command, args...)
	if err != nil {
		return nil, utils.ReasonableError{
			Reason:  err,
			Message: "Error in quering command '" + command + "'",
		}
	}
	return rows, err
}
