package db

import (
	"context"
	"database/sql"
	"ticket-booking/internals/core/ticket"
)

type ticketRepoImpl struct {
	db *sql.DB
}

func NewTicketRepoImpl(db *sql.DB) *ticketRepoImpl {
	return &ticketRepoImpl{
		db: db,
	}
}

func (r *ticketRepoImpl) BookSingleTicket(ctx context.Context, req *ticket.TicketReq) error {

	// begin the transaction

	tx, err := r.db.Begin()
	if err != nil {
		tx.Rollback()
		return err

	}

	// check if this seat is booked

	var count int
	query := `SELECT count(seat_number) from user_tickets where seat_number=$1`
	err = tx.QueryRowContext(ctx, query, req.SeatNumber).Scan(&count)
	if err != nil {
		tx.Rollback()
		return err
	}
	if count == 1 {
		tx.Rollback()
		return ErrSeatIsBooked
	}
	query = `INSERT into user_tickets (ticket_id,event_id,user_id,seat_number) VALUES($1,$2,$3,$4)`
	_, err = tx.ExecContext(ctx, query, req.TicketID, req.EventID, req.UserID, req.SeatNumber)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
