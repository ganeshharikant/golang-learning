package handlers

import (
	"encoding/json"
	"movie_booking_using_basic_crud_api/models"
	"movie_booking_using_basic_crud_api/store"
	"net/http"
)


func CreateBooking(w http.ResponseWriter, r *http.Request) {
	var body models.Booking

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", 400)
		return
	}

	result := store.SaveBooking(body)

	w.WriteHeader(201)
	json.NewEncoder(w).Encode(result)

}


