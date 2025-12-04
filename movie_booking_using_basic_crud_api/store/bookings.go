package store

import (
	"movie_booking_using_basic_crud_api/models"
	"time"
)

var bookings []models.Booking
var idCounter = 5001



func SaveBooking(b models.Booking) models.Booking {
	b.BookingID = idCounter
	b.BookedAt = time.Now()
	b.IsActive = true

	idCounter++
	bookings = append(bookings, b)
	return b
}
