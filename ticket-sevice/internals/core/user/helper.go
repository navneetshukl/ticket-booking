package user

import "ticket-booking/pkg/helpers"

// GenerateUserID generates a user ID
func GenerateUserID() (string, error) {
	id, err := helpers.GenerateID(8)
	if err != nil {
		return "", err
	}

	return "USR-" + id, nil
}
