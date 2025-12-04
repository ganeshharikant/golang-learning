package main

import (
	"movie_booking_using_basic_crud_api/handlers"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/bookings", handlers.CreateBooking)
	

	println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
