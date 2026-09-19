# GoLang Learning 🚀

A hands-on repository for learning **Go (Golang)** from the fundamentals to building real-world applications.

This repository contains my notes, code examples, experiments, problem-solving practice, and projects while learning Go.

The primary goal is to **learn by writing and executing code**, rather than only studying theory.

---

## 🎯 Goals

* Learn Go from fundamentals to advanced concepts
* Understand Go's syntax, conventions, and philosophy
* Practice concepts through small executable programs
* Build progressively complex projects
* Learn Go's approach to concurrency and networking
* Develop good Go programming practices
* Prepare for real-world backend/software development using Go

---

## 🛠️ Setup

### Prerequisites

* [Go](https://go.dev/) installed
* Visual Studio Code
* Go extension for VS Code
* Git
* GitHub account

Check the Go installation:

```bash
go version
```

Check Git:

```bash
git --version
```

---

## ▶️ Running Go Code

Run a Go file directly:

```bash
go run main.go
```

Build a program:

```bash
go build
```

Format code:

```bash
go fmt ./...
```

Run all tests:

```bash
go test ./...
```

---

## 📚 Learning Roadmap

### 01 — Basics

* Hello World
* Variables
* Constants
* Data types
* Type conversion
* Input / Output
* Operators
* Comments
* Go syntax and conventions

### 02 — Control Flow

* `if / else`
* `switch`
* `for`
* `break`
* `continue`
* Nested loops

### 03 — Functions

* Function declaration
* Parameters
* Return values
* Multiple return values
* Named return values
* Variadic functions
* Anonymous functions
* Closures

### 04 — Arrays & Slices

* Arrays
* Slices
* `append`
* `copy`
* Slice capacity and length
* Multidimensional slices

### 05 — Maps & Strings

* Maps
* Map operations
* Strings
* Runes
* String manipulation
* Unicode

### 06 — Structs, Methods & Pointers

* Structs
* Struct fields
* Methods
* Pointer basics
* Value vs pointer receivers
* Embedded structs

### 07 — Interfaces

* Interfaces
* Implicit interface implementation
* Empty interface
* Type assertions
* Type switches
* Polymorphism in Go

### 08 — Error Handling

* `error`
* Creating errors
* Error wrapping
* Custom errors
* `panic`
* `recover`

### 09 — Packages & Modules

* Packages
* Imports
* Exported vs unexported identifiers
* Go modules
* Dependency management
* Creating custom packages

### 10 — Concurrency

* Goroutines
* Channels
* Buffered channels
* `select`
* `sync.WaitGroup`
* Mutexes
* Race conditions
* Concurrent programming patterns

### 11 — File Handling

* Reading files
* Writing files
* Directories
* CSV
* JSON
* Serialization / deserialization

### 12 — Testing

* Unit testing
* Table-driven tests
* Test coverage
* Benchmarks
* Example tests

### 13 — HTTP & APIs

* HTTP servers
* HTTP clients
* Request / response handling
* REST APIs
* JSON APIs
* Middleware
* Routing

### 14 — Databases

* SQL with Go
* Database connections
* CRUD operations
* Transactions
* Database-backed applications

### 15 — Projects

Projects that combine multiple Go concepts into practical applications.

Potential projects:

* CLI applications
* Todo API
* URL shortener
* REST API
* Concurrent web scraper
* File processing utility
* Authentication service
* Backend service

---

## 📁 Repository Structure

```text
golang-learning/
│
├── 01-basics/
├── 02-control-flow/
├── 03-functions/
├── 04-arrays-slices/
├── 05-maps-strings/
├── 06-structs-methods/
├── 07-interfaces/
├── 08-error-handling/
├── 09-packages/
├── 10-concurrency/
├── 11-file-handling/
├── 12-testing/
├── 13-http/
├── 14-database/
├── 15-projects/
│
├── notes/
│   ├── concepts.md
│   ├── interview-questions.md
│   └── useful-commands.md
│
├── README.md
└── go.mod
```

Each topic contains small, independent programs that can be compiled and executed while learning.

---

## 🧪 Learning Approach

For each concept, I aim to follow this process:

```text
Learn concept
     ↓
Write a small example
     ↓
Run the program
     ↓
Experiment / modify it
     ↓
Break it intentionally
     ↓
Understand the error
     ↓
Solve a small problem
     ↓
Move to the next concept
```

The repository is intentionally focused on **hands-on learning and experimentation**.

---

## 📝 Notes

Important concepts, observations, commands, and interview-related material are maintained separately in the `notes/` directory.

### Notes

* `concepts.md` — Important Go concepts and explanations
* `interview-questions.md` — Go interview questions and answers
* `useful-commands.md` — Frequently used Go and Git commands

---

## 🔄 Git Workflow

After completing a learning session:

```bash
git status
git add .
git commit -m "Learned Go slices and maps"
git push
```

Example commits:

```text
Learned Go variables and data types
Added control flow examples
Practiced functions
Added slice and map examples
Learned structs and methods
Practiced goroutines and channels
Added HTTP server example
```

---

## 📈 Progress

* [ ] Go basics
* [ ] Control flow
* [ ] Functions
* [ ] Arrays & slices
* [ ] Maps & strings
* [ ] Structs & methods
* [ ] Pointers
* [ ] Interfaces
* [ ] Error handling
* [ ] Packages & modules
* [ ] Concurrency
* [ ] File handling
* [ ] Testing
* [ ] HTTP
* [ ] Databases
* [ ] Real-world projects

---

## 🔗 Resources

* [Official Go Website](https://go.dev/)
* [A Tour of Go](https://go.dev/tour/)
* [Go Documentation](https://go.dev/doc/)
* [Go by Example](https://gobyexample.com/)

---

## 📌 Repository Philosophy

> **Don't just read Go. Write Go.**

Every concept in this repository should ideally have executable code demonstrating how it works.

The repository will evolve alongside my understanding of Go, from simple `Hello World` programs to complete applications.
