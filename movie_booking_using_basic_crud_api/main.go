package main

import (
	"movie_booking_using_basic_crud_api/handlers"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/bookings", handlers.CreateBooking)
		mux.HandleFunc("DELETE /api/bookings/{booking_id}", handlers.DeleteBooking)
		mux.HandleFunc("GET /api/bookings", handlers.GetAllBookings)
     mux.HandleFunc("GET /api/bookings/{booking_id}",handlers.GetBookingByID)
     mux.HandleFunc("PUT /api/bookings/{booking_id}",handlers.UpdateBooking)
     mux.HandleFunc("PATCH /api/bookings/{booking_id}",handlers.UpdateBooking)


	println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
