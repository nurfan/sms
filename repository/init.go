package repository

import (
	"github.com/jmoiron/sqlx"
)

// RepositoryPsql ...
type RepositoryPsql struct {
	Conn *sqlx.DB
}

// NewRepositoryPsql ...
func NewRepositoryPsql(Conn *sqlx.DB) *RepositoryPsql {
	return &RepositoryPsql{
		Conn: Conn,
	}
}
