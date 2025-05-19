package ticket

import (
	"context"
	"ticket-booking/internals/core/ticket"
)

type TicketUseCaseImpl struct {
}

func NewTicketUseCase() ticket.TicketUseCase {
	return &TicketUseCaseImpl{}
}

func(t *TicketUseCaseImpl)BookSingleTicket(ctx context.Context,req *ticket.TicketReq)error{
return nil
}