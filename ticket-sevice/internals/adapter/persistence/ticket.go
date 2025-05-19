package db

import "database/sql"

type ticketRepoImpl struct {
	db *sql.DB
}

func NewTicketRepoImpl(db *sql.DB)*ticketRepoImpl{
	return &ticketRepoImpl{
		db: db,
	}
}
