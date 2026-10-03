# Expense Tracker

A simple but structured command-line expense tracker written in Go.

The project is intentionally small enough to understand while learning Go, but it uses a real application structure with packages, persistent storage, validation, commands, and unit tests.

## Features

- Add expenses
- List all expenses
- Filter expenses by category
- Delete an expense by ID
- Show total spending
- Show spending grouped by category
- Filter summaries by month
- Persistent JSON storage
- Input validation
- Unit tests
- Atomic-ish file replacement when saving data

## Project Structure

```text
expense-tracker/
├── go.mod
├── main.go
├── README.md
├── data/
│   └── expenses.json          # Created automatically at runtime
└── internal/
    ├── expense/
    │   ├── expense.go
    │   └── expense_test.go
    └── storage/
        ├── json_store.go
        └── json_store_test.go
```

## Requirements

- Go 1.22 or newer

Check your installation:

```bash
go version
```

## Run the Project

From the `expense-tracker` directory:

```bash
go run .
```

This prints the available commands.

## Commands

### Add an expense

```bash
go run . add -amount 250 -category Food -description "Lunch"
```

You can provide a specific date:

```bash
go run . add -amount 1200 -category Travel -description "Metro pass" -date 2026-10-02
```

If no date is supplied, today's date is used.

### List expenses

```bash
go run . list
```

Example:

```text
ID   DATE         CATEGORY         DESCRIPTION                         AMOUNT
------------------------------------------------------------------------------
2    2026-10-02   Travel           Metro pass                           1200.00
1    2026-10-01   Food             Lunch                                 250.00
```

### Filter by category

```bash
go run . list -category Food
```

Category filtering is case-insensitive.

### Delete an expense

```bash
go run . delete -id 1
```

### Show summary

```bash
go run . summary
```

Example:

```text
Number of expenses: 3
Total spent:        1650.00

By category:
  Food             450.00
  Travel           1200.00
```

### Summary for one category

```bash
go run . summary -category Food
```

### Summary for a month

Use `YYYY-MM`:

```bash
go run . summary -month 2026-10
```

You can combine filters:

```bash
go run . summary -category Food -month 2026-10
```

## Data Storage

Expenses are stored locally in:

```text
data/expenses.json
```

The `data` directory and JSON file are created automatically the first time the application runs.

Example data:

```json
[
  {
    "id": 1,
    "amount": 250,
    "category": "Food",
    "description": "Lunch",
    "date": "2026-10-01T00:00:00Z"
  }
]
```

The data file is intentionally not committed to Git. If you want to keep personal expense data locally, add `data/expenses.json` to `.gitignore`.

## Run Tests

Run all tests:

```bash
go test ./...
```

Run tests with verbose output:

```bash
go test -v ./...
```

## What This Project Teaches

This project is useful as an early Go project because it introduces several concepts that will appear in larger applications:

### Go fundamentals

- Variables
- Functions
- Structs
- Slices
- Maps
- Packages
- Error handling
- Pointers/references through normal Go usage

### Standard library

- `flag`
- `fmt`
- `os`
- `encoding/json`
- `path/filepath`
- `time`
- `strings`

### Software engineering

- Separating business logic from storage
- Input validation
- Unit testing
- Persistent data
- Command-line interfaces
- Package organization
- Atomic-style file replacement

## Design Notes

The project intentionally does not use a database or third-party CLI framework yet.

That is a deliberate learning decision.

For a first Go application, using the standard library makes it easier to understand what Go itself provides.

A future production-oriented version could replace JSON storage with PostgreSQL and add:

- REST API
- Authentication
- User accounts
- Database migrations
- Docker
- Redis caching
- Monthly budgets
- Recurring expenses
- CSV import/export
- Web dashboard

Those additions would turn this from a learning project into a more realistic backend application.

## Suggested Learning Exercises

After understanding the existing implementation, try adding these yourself:

1. Add an `edit` command.
2. Add a `-limit` option to `list`.
3. Add a `-from` and `-to` date filter.
4. Add monthly spending statistics.
5. Add CSV export.
6. Add CSV import.
7. Replace JSON storage with SQLite or PostgreSQL.
8. Add an HTTP REST API.
9. Add authentication.
10. Add automated API tests.

Do not implement all of these immediately. Use them as progressively harder exercises.

## License

This project is intended as a learning project and portfolio foundation.
