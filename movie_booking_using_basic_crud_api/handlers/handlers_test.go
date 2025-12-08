package handlers

import (
	"bytes"
	"encoding/json"
	"movie_booking_using_basic_crud_api/models"
	"movie_booking_using_basic_crud_api/store"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestCreateBooking(t *testing.T) {

	body := map[string]string{
		"film_title":  "Leo",
		"screen":      "Screen-2",
		"seat_number": "A10",
		"customer":    "Ganesh",
	}

	b, _ := json.Marshal(body)

	req := httptest.NewRequest("POST", "/api/bookings", bytes.NewBuffer(b))

	w := httptest.NewRecorder()

	CreateBooking(w, req)

	if w.Result().StatusCode != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Result().StatusCode)
	}
}

func TestGetAllBookings(t *testing.T) {

	store.SaveBooking(models.Booking{
		FilmTitle:  "Test",
		Screen:     "Scr-1",
		SeatNumber: "A1",
		Customer:   "User",
	})
	req := httptest.NewRequest("GET", "/api/bookings", nil)
	w := httptest.NewRecorder()

	GetAllBookings(w, req)

	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Result().StatusCode)
	}
}

func TestGetBookingByID(t *testing.T) {

	store.SaveBooking(models.Booking{

		FilmTitle:  "Test",
		Screen:     "Scr-1",
		SeatNumber: "A2",
		Customer:   "User",
		IsActive:   true,
	})

	req := httptest.NewRequest("GET", "/api/bookings/1", nil)

	req.SetPathValue("booking_id", "1")
	w := httptest.NewRecorder()

	GetBookingByID(w, req)

	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Result().StatusCode)
	}
}
func TestUpdateBooking_PUT(t *testing.T) {

	saved := store.SaveBooking(models.Booking{
		FilmTitle:  "Old Title",
		Screen:     "Old Screen",
		SeatNumber: "B1",
		Customer:   "Old User",
		IsActive:   true,
	})

	body := map[string]string{
		"film_title":  "Updated",
		"screen":      "Scr-5",
		"seat_number": "C2",
		"customer":    "Rocky",
	}

	b, _ := json.Marshal(body)
	idStr := strconv.Itoa(saved.BookingID)

	req := httptest.NewRequest("PUT", "/api/bookings/"+idStr, bytes.NewBuffer(b))
	req.SetPathValue("booking_id", idStr)

	w := httptest.NewRecorder()

	UpdateBooking(w, req)

	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Result().StatusCode)
	}
}

func TestUpdateBooking_PATCH(t *testing.T) {

	saved := store.SaveBooking(models.Booking{
		FilmTitle:  "Old Title",
		Screen:     "Old Screen",
		SeatNumber: "B1",
		Customer:   "Old User",
		IsActive:   true,
	})

	body := map[string]string{
		"customer": "Updated-Only-Name",
	}

	b, _ := json.Marshal(body)
	idStr := strconv.Itoa(saved.BookingID)

	req := httptest.NewRequest("PATCH", "/api/bookings/1", bytes.NewBuffer(b))
	req.SetPathValue("booking_id", idStr)

	w := httptest.NewRecorder()

	UpdateBooking(w, req)

	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Result().StatusCode)
	}
}

func TestDeleteBooking(t *testing.T) {

	req := httptest.NewRequest("DELETE", "/api/bookings/1", nil)
	req.SetPathValue("booking_id", "1")

	w := httptest.NewRecorder()

	DeleteBooking(w, req)

	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Result().StatusCode)
	}
}