package repository

import (
	"time"

	"github.com/wycemiro/booking-site/internal/models"
)

type DatabaseRepo interface {
	InsertReservation(res models.Reservation) (int, error)
	InstertRoomRestriction(r models.RoomRestriction) error
	SearchAvailabilityByDates(start, end time.Time, roomID int) (bool, error)
}
