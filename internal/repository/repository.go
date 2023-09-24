package repository

import "github.com/caovanhoang63/bookings/internal/models"

// DatabaseRepo is the interface for the database repository
type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(models.Reservation) (int, error)
	InsertRoomRestriction(res models.RoomRestrictions) error
}
