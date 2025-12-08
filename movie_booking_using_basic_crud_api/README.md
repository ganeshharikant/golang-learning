🎬 Movie Booking API (Go)

A simple CRUD API built using Go, net/http, and in-memory storage (slice).

✅ API Endpoints
1️⃣ Create Booking

POST /api/bookings

2️⃣ Get All Active Bookings

GET /api/bookings

3️⃣ Get Booking by ID

GET /api/bookings/{id}

4️⃣ Update Booking

PUT /api/bookings/{id}

PATCH /api/bookings/{id}

5️⃣ Delete Booking (Soft Delete)

DELETE /api/bookings/{id}

🧪 Unit Tests

All handlers have basic unit test coverage using httptest.

Run tests:

go test ./...

🚀 How to Run the Project

Clone the repo

Run server:

go run main.go


Server starts at:

http://localhost:8080