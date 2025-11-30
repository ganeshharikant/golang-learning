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



