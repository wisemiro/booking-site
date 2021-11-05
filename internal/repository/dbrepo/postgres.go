package dbrepo

import (
	"context"
	"time"

	"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/wycemiro/booking-site/internal/models"
)

func (m *postgresDBRepo) InsertReservation(res models.Reservation) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	var newID int

	stmt := `insert into reservations (first_name, last_name, email, phone, start_date,
			 end_date, room_id, created_at, updated_at) 
			 values ($1,$2,$3,$4,$5,$6,$7,$8,$9)
			 return id
			 `
	err := m.DB.QueryRowContext(
		ctx,
		stmt,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		res.StartDate,
		res.EndDate,
		res.RoomID,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

func (m *postgresDBRepo) InstertRoomRestriction(r models.RoomRestriction){}