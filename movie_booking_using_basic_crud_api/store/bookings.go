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
func DeactivateBooking(id int) bool {
	for i, v := range bookings {
		if v.BookingID == id {
			if !v.IsActive {
				return false 
			}
			bookings[i].IsActive = false
			return true
		}
	}
	return false
}

func GetAllActive() []models.Booking {
	active := []models.Booking{}

	for _, b := range bookings {
		if b.IsActive {
			active = append(active, b)
		}
	}
	return active
}


func GetOneActive(id int) (models.Booking, bool) {
	for _, b := range bookings {
		if b.BookingID == id && b.IsActive {
			return b, true
		}
	}
	return models.Booking{}, false
}


func UpdateBooking(id int, data models.Booking, isPatch bool)(models.Booking, bool) {
 for i, b := range bookings {
      if b.BookingID == id && b.IsActive {
          if isPatch{
               if data.FilmTitle != "" {
					bookings[i].FilmTitle = data.FilmTitle
				}
				if data.Screen != "" {
					bookings[i].Screen = data.Screen
				}
				if data.SeatNumber != "" {
					bookings[i].SeatNumber = data.SeatNumber
				}
				if data.Customer != "" {
					bookings[i].Customer = data.Customer
				}

			}else {
				
				bookings[i].FilmTitle = data.FilmTitle
				bookings[i].Screen = data.Screen
				bookings[i].SeatNumber = data.SeatNumber
				bookings[i].Customer = data.Customer
			}

			return bookings[i], true
		  }

	  }
	return models.Booking{}, false

}

