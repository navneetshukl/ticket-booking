package ticket

import "errors"

var (
	ErrGeneratingTicketID error = errors.New("error in creating new ticket id")
	ErrSeatIsBooked error=errors.New("seat is already booked")
	ErrInsertingTicketDetail error=errors.New("error in booking seat")
)