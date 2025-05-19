package ticket

import "time"

type Ticket struct {
	TicketID   string     `json:"ticket_id"`
	UserID     string     `json:"user_id"`
	EventID    string     `json:"event_id"`
	SeatNumber int        `json:"seat_number"`
	BookedDate *time.Time `json:"booked_date"`
	BookedTime *time.Time `json:"booked_time"`
}

type TicketReq struct {
	UserID     string `json:"user_id"`
	TicketID   string `json:"ticket_id"`
	EventID    string `json:"event_id"`
	SeatNumber int    `json:"seat_number"`
}

type TicketUseCase interface{}

