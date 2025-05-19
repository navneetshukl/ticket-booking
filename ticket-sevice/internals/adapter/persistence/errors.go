package db

import "errors"

var (
	ErrSeatIsBooked error = errors.New("seat is booked")
)