package models

import "time"

type Booking struct {
    BookingID   int       `json:"booking_id"`
    FilmTitle   string    `json:"film_title"`
    Screen      string    `json:"screen"`
    SeatNumber  string    `json:"seat_number"`
    Customer    string    `json:"customer"`
    IsActive    bool      `json:"is_active"`
    BookedAt    time.Time `json:"booked_at"`
}
