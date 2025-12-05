package handlers

import (
	"encoding/json"
	"movie_booking_using_basic_crud_api/models"
	"movie_booking_using_basic_crud_api/store"
	"net/http"
	"strconv"
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

func DeleteBooking(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("booking_id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid booking ID", 400)
		return
	}

	ok := store.DeactivateBooking(id)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Booking not found",
		})
		return 
	}else{

w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Booking deleted successfully",
	})

	}

	
}

func GetAllBookings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bookings := store.GetAllActive()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bookings)
}

