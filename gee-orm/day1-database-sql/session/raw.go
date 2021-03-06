package session

import (
	"database/sql"
	"geeorm/log"
)

// Session keep a pointer to sql.DB and provides all execution of all
// kind of database operations.
type Session struct {
	db      *sql.DB
	sql     string
	sqlVars []interface{}
}

// New creates a instance of Session
func New(db *sql.DB) *Session {
	return &Session{db: db}
}

// Clear initialize the state of a session
func (s *Session) Clear() {
	s.sqlVars = nil
	s.sql = ""
}

// Exec raw sql with sqlVars
func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	log.Info(s.sql, s.sqlVars)
	if result, err = s.db.Exec(s.sql, s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}

// QueryRow gets a record from db
func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Info(s.sql, s.sqlVars)
	return s.db.QueryRow(s.sql, s.sqlVars...)
}

// QueryRows gets a list of records from db
func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	log.Info(s.sql, s.sqlVars)
	if rows, err = s.db.Query(s.sql, s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}

// Raw appends sql and sqlVars
func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql += sql
	s.sqlVars = append(s.sqlVars, values...)
	return s
}
