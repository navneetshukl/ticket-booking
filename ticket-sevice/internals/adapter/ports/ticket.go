package ports

import (
	"context"
	"ticket-booking/internals/core/ticket"
)

type TicketRepository interface {
	BookSingleTicket(ctx context.Context, req *ticket.TicketReq) error
}
