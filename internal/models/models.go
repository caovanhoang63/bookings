package models

import (
	"time"
)

// Reservation is present reservation form
type Reservation struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

type Users struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	AccessLevel int
}

type Rooms struct {
	ID        int
	RoomName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Restrictions struct {
	ID              int
	RestrictionName string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Reservations struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
	RoomID    int
	StartDate time.Time
	EndDate   time.Time
	Room      Rooms
}

type RoomRestrictions struct {
	ID            int
	StartDate     time.Time
	EndDate       time.Time
	RoomID        int
	ReservationID int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	RestrictionID int
	Room          Rooms
	Restriction   Restrictions
	Reservation   Reservations
}
