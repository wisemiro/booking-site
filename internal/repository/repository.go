package repository

import "github.com/wycemiro/booking-site/internal/models"

type DatabaseRepo interface {
	InsertReservation(res models.Reservation) (int, error)
}