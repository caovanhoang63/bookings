package repository

import (
	"github.com/caovanhoang63/bookings/internal/models"
	"time"
)

// DatabaseRepo is the interface for the database repository
type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(models.Reservation) (int, error)
	InsertRoomRestriction(res models.RoomRestrictions) error
	SearchForAvailabilityRoomByRoomID(start, end time.Time, roomID int) (bool, error)
	SearchAvailabilityForAllRoom(start, end time.Time) ([]models.Room, error)
}
