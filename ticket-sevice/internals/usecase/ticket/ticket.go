package ticket

import (
	"context"
	"errors"
	db "ticket-booking/internals/adapter/persistence"
	tktRepo "ticket-booking/internals/adapter/ports"
	"ticket-booking/internals/core/ticket"
	"ticket-booking/pkg/logger"
)

type TicketUseCaseImpl struct {
	ticketRepo tktRepo.TicketRepository
}

func NewTicketUseCase(repo tktRepo.TicketRepository) ticket.TicketUseCase {
	return &TicketUseCaseImpl{
		ticketRepo: repo,
	}
}

// BookSingleTicket function will book ticket for single user
func (t *TicketUseCaseImpl) BookSingleTicket(ctx context.Context, req *ticket.TicketReq) error {

	// check from redis if this seat number is not booked(do it later)
	ticketID, err := ticket.GenerateTicketID()
	if err != nil {
		logger.LogStatus("ERROR", "error in generating ticket id")
		return ticket.ErrGeneratingTicketID
	}
	req.TicketID = ticketID

	err = t.ticketRepo.BookSingleTicket(ctx, req)
	if err != nil {
		logger.LogStatus("ERROR", err.Error())
		if errors.Is(err, db.ErrSeatIsBooked) {
			return ticket.ErrSeatIsBooked
		}
		return ticket.ErrInsertingTicketDetail
	}

	return nil
}
