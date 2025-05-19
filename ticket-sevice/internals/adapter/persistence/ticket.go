package db

import (
	"context"
	"database/sql"
	"ticket-booking/internals/adapter/ports"
	"ticket-booking/internals/core/ticket"
)

type ticketRepoImpl struct {
	db *sql.DB
}

func NewTicketRepoImpl(db *sql.DB) ports.TicketRepository {
	return &ticketRepoImpl{
		db: db,
	}
}

func (r *ticketRepoImpl) BookSingleTicket(ctx context.Context, req *ticket.TicketReq) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		}
	}()

	// Check if seat is already booked
	var count int
	query := `SELECT COUNT(seat_number) FROM user_tickets WHERE seat_number = $1 AND event_id = $2`
	err = tx.QueryRowContext(ctx, query, req.SeatNumber, req.EventID).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return ErrSeatIsBooked
	}

	// Book the ticket
	query = `INSERT INTO user_tickets (ticket_id, event_id, user_id, seat_number) VALUES ($1, $2, $3, $4)`
	_, err = tx.ExecContext(ctx, query, req.TicketID, req.EventID, req.UserID, req.SeatNumber)
	if err != nil {
		return err
	}

	err = tx.Commit()
	return err
}

