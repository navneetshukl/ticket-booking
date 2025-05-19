package ticket

import (
	"ticket-booking/pkg/helpers"
)

// GenerateTicketID generates a ticket ID
func GenerateTicketID() (string, error) {
	id, err := helpers.GenerateID(10)
	if err != nil {
		return "", err
	}
	return "TKT-" + id, nil
}
