# 🚀 Go Learning Journey

## 📅 Day 1 – Topics Covered

Today, I learned the fundamental concepts of Go programming.  
Below are the folders I completed on Day 1:

| Folder Name | Topic |
|-------------|-------|
| 1_hello_world | Hello World Program |
| 2_simple_values | Simple Values & Basic Data Types |
| 03_variables | Variables and Data Types |
| 04_const | Constants |
| 05_for_loop | For Loops |
| 06_ifelse | If Else Conditions |
| 07_switch | Switch Statement |
| 8_arrays | Arrays in Go |
| 9_slice | Slices in Go |
| 10_maps | Maps in Go |

---

### 📝 Summary of Day 1 Learning

- Learned how Go workspace and folder structure works  
- Wrote my first Go program (`Hello World`)  
- Understood variables, constants, and data types  
- Practiced loops, conditions, and switch statements  
- Explored arrays, slices, and maps  
- Learned difference between array and slice  
- Pushed my Go code to GitHub for the first time 🎉

---


🚀 Go Learning Journey

📅 Day 2 – Topics Covered

Today, I strengthened my understanding of Go Functions and built a small app.
Below are the topics and concepts I completed on Day 2:

Folder/Concept	Topic
11_closures	Closures (function inside function)
12_higher_order	Higher-Order Functions (functions as values, parameters, return types)
13_variadic_functions	Variadic Functions (func sum(nums ...int))

📝 Summary of Day 2 Learning

Understood Closures with simple examples

Practiced Higher-Order functions

Learned and used Variadic function syntax

Built a Calculator CLI app using functions





📅 Day 3 – Topics Covered

Today, I learned and implemented deeper backend-level Go concepts focused on memory reference, struct design, composition, and abstraction.

| Folder/Concept                  | Topic / Understanding                                                                                                                     |
| ------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------- |
| **14_pointers**                 | Understood memory addresses using `&` and value access/modification via dereferencing `*` (reference-based updates).                      |
| **15_structs**                  | Learned to model real entities using structs (`type User struct {}`) and access/modify fields using dot notation (`u.name`).              |
| **16_embedded_structs**         | Learned struct composition by embedding one struct inside another for reuse (`type A struct { B }`).                                      |
| **17_interfaces**               | Learned method contracts using `type Animal interface { Speak() }` and struct-based interface implementation for scalable backend design. |
| **18_struct_interface_binding** | Practiced assigning a struct instance to an interface type and calling methods polymorphically (`var a Animal = Dog{}`).                  |

📝 Summary of Day 3 Learning

Pointers help pass data by reference instead of copy, useful inside functions and struct updates

Structs are used for clean data modeling in backend systems

Embedded structs provide structure reuse without duplication (Go’s composition model)

Interfaces ensure scalable, testable, abstraction-friendly code

Practiced interface implementation inside structs

No direct slice == comparison → used slices.Equal() ✅

Map grows by inserting keys, slice grows by append() ✅

Ran Go code using terminal inside VS Code confidently



## 📅 Day 4 – Topics Covered

Today, I built a **modular CLI backend simulation project (Smart Inventory System)** and applied multiple Go fundamentals together like a real backend flow (without database).

### ✅ Concepts Practiced Today

| Concept | Usage in practice |
|--------|------------------|
| Slice | Created and modified order history (`[]string`) |
| Map | Built fast lookup storage (`map[string]*inventory`) |
| Pointer | Updated stock using reference (`*inventory`) |
| Embedded Struct | Inventory reused product fields via composition |
| Interface | Implemented stock update contract (`StockUpdater`) |
| Loop/Range | Iterated over map & slice for prints and analytics |
| Switch | Built CLI menu & category routing logic |
| Functions (multi-return) | Returned total stock + total products count |
| If-Else | Added validation for name, price, and stock updates |

---

### 🧪 Project Backend Flow Tested

- Added multiple products into inventory map  
- Updated product stock using pointer/interface method  
- Placed repeated orders stored in a slice  
- Printed full inventory using loop/range  
- Counted order frequency using a map + loop  
- Routed products using switch inside function  
- Printed analytics and stats clearly in CLI  

---

### 📝 Summary of Day 4 Learning

- Learned how to **break backend logic into small modular functions**
- Practiced **pointer-based updates inside embedded structs ✅**
- Stored values in **map using pointers for real updates ✅**
- Appended order history using slice (`append()`) ✅
- Used `switch` for CLI menu control and category routing ✅
- Created and handled functions returning 2 values confidently ✅
- Ran and tested my Go project successfully using terminal + VS Code
- Pushed Day 4 updates to GitHub with meaningful commit messages 🎉

🚀 Feeling more confident building backend simulations with Go!


📅 Day 5 – Topics Covered

Today, I learned and practiced concurrent and scalable Go concepts including goroutines, synchronization using WaitGroup, generics, and enum-like patterns using constants/iota.



| Concept                    | Usage / What I learned                                                                                          |
| -------------------------- | --------------------------------------------------------------------------------------------------------------- |
| **Goroutines**             | Learned lightweight concurrency using `go func()` for parallel execution and improved performance.              |
| **WaitGroup**              | Understood synchronization to wait for multiple goroutines using `sync.WaitGroup`, `Add()`, `Done()`, `Wait()`. |
| **Enums**                  | Implemented enum-like behavior using `iota` with constants (`type Status int`).                                 |
| **Generics**               | Learned to create reusable type-safe functions using generics (`func PrintAny[T any](v T)`).                    |
| **WaitGroup + Slice Flow** | Used `WaitGroup` inside loops to wait for parallel tasks dynamically.                                           |


Development Flow Simulated

Ran functions concurrently using goroutines ✅

Waited for multiple goroutine tasks inside main flow using WaitGroup ✅

Built reusable function logic using Generics ✅

Created backend state categories using Enums (iota) ✅

Used wg.Wait() to control parallel execution finish before next step ✅


 
Movie Booking CRUD API Project

Today, I built a complete CRUD-based Movie Booking Backend API using Go.
This project helped me understand real API development, routing, JSON handling, and backend logic flow.

| Feature               | Description                                          |
| --------------------- | ---------------------------------------------------- |
| **Create Movie**      | Add new movie with name, rating, and available seats |
| **Read All Movies**   | Fetch list of all movies                             |
| **Read One Movie**    | Fetch a single movie by ID                           |
| **Update Movie**      | Modify movie details (rating, seats, name)           |
| **Delete Movie**      | Remove movie by ID                                   |
| **Book a Ticket**     | Reduce seats only if seats are available             |
| **In-Memory DB**      | Used slice/map as temporary storage                  |
| **Validation**        | Checked for missing fields & wrong IDs               |
| **JSON Input/Output** | Full JSON body handling for requests                 |
Tested API Endpoints
1️⃣ Create Movie

POST /movies

2️⃣ Get All Movies

GET /movies

3️⃣ Get Movie by ID

GET /movies/{id}

4️⃣ Update Movie

PUT /movies/{id}

5️⃣ Delete Movie

DELETE /movies/{id}

6️⃣ Book Movie Ticket

POST /movies/{id}/book

✔ Tested everything using Postman
✔ Validated JSON requests/responses
✔ Confirmed CRUD flow end-to-end

📝 Summary of Day 10 Learning

Built a functional CRUD backend in Go

Understood REST routing and HTTP methods

Learned JSON marshaling/unmarshaling

Modeled data using structs

Used maps and slices to store runtime data

Improved error handling using proper status codes

Learned how booking affects seat count

Pushed project to GitHub with clean folder structure


