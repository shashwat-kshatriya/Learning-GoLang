package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"expense-tracker/internal/expense"
	"expense-tracker/internal/storage"
)

const dateFormat = "2006-01-02"

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	store := storage.NewJSONStore("data/expenses.json")
	if err := store.Ensure(); err != nil {
		exitWithError(err)
	}

	switch os.Args[1] {
	case "add":
		runAdd(store, os.Args[2:])
	case "list":
		runList(store, os.Args[2:])
	case "delete":
		runDelete(store, os.Args[2:])
	case "summary":
		runSummary(store, os.Args[2:])
	case "help", "-h", "--help":
		printUsage()
	default:
		fmt.Printf("Unknown command: %s\n\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}
}

func runAdd(store *storage.JSONStore, args []string) {
	fs := flag.NewFlagSet("add", flag.ExitOnError)
	amount := fs.String("amount", "", "expense amount, e.g. 250.50")
	category := fs.String("category", "", "expense category, e.g. Food")
	description := fs.String("description", "", "short description")
	date := fs.String("date", time.Now().Format(dateFormat), "date in YYYY-MM-DD format")
	fs.Usage = func() {
		fmt.Println("Usage: expense-tracker add -amount <amount> -category <category> -description <description> [-date YYYY-MM-DD]")
		fs.PrintDefaults()
	}
	fs.Parse(args)

	if *amount == "" || *category == "" || *description == "" {
		fs.Usage()
		os.Exit(1)
	}

	value, err := strconv.ParseFloat(*amount, 64)
	if err != nil {
		exitWithError(fmt.Errorf("invalid amount %q: %w", *amount, err))
	}

	parsedDate, err := time.Parse(dateFormat, *date)
	if err != nil {
		exitWithError(fmt.Errorf("invalid date %q; use YYYY-MM-DD", *date))
	}

	item, err := expense.New(value, *category, *description, parsedDate)
	if err != nil {
		exitWithError(err)
	}

	items, err := store.Load()
	if err != nil {
		exitWithError(err)
	}

	item.ID = expense.NextID(items)
	items = append(items, item)

	if err := store.Save(items); err != nil {
		exitWithError(err)
	}

	fmt.Printf("Expense added successfully. ID: %d\n", item.ID)
}

func runList(store *storage.JSONStore, args []string) {
	fs := flag.NewFlagSet("list", flag.ExitOnError)
	category := fs.String("category", "", "filter by category")
	fs.Usage = func() {
		fmt.Println("Usage: expense-tracker list [-category <category>]")
		fs.PrintDefaults()
	}
	fs.Parse(args)

	items, err := store.Load()
	if err != nil {
		exitWithError(err)
	}

	if *category != "" {
		items = expense.FilterByCategory(items, *category)
	}

	if len(items) == 0 {
		fmt.Println("No expenses found.")
		return
	}

	fmt.Printf("%-4s %-12s %-16s %-28s %12s\n", "ID", "DATE", "CATEGORY", "DESCRIPTION", "AMOUNT")
	fmt.Println(strings.Repeat("-", 78))

	for _, item := range items {
		fmt.Printf("%-4d %-12s %-16s %-28s %12.2f\n",
			item.ID,
			item.Date.Format(dateFormat),
			item.Category,
			truncate(item.Description, 28),
			item.Amount,
		)
	}
}

func runDelete(store *storage.JSONStore, args []string) {
	fs := flag.NewFlagSet("delete", flag.ExitOnError)
	id := fs.Int("id", 0, "ID of the expense to delete")
	fs.Usage = func() {
		fmt.Println("Usage: expense-tracker delete -id <expense-id>")
		fs.PrintDefaults()
	}
	fs.Parse(args)

	if *id <= 0 {
		fs.Usage()
		os.Exit(1)
	}

	items, err := store.Load()
	if err != nil {
		exitWithError(err)
	}

	updated, deleted := expense.DeleteByID(items, *id)
	if !deleted {
		exitWithError(fmt.Errorf("expense with ID %d not found", *id))
	}

	if err := store.Save(updated); err != nil {
		exitWithError(err)
	}

	fmt.Printf("Expense %d deleted successfully.\n", *id)
}

func runSummary(store *storage.JSONStore, args []string) {
	fs := flag.NewFlagSet("summary", flag.ExitOnError)
	category := fs.String("category", "", "show summary for one category")
	month := fs.String("month", "", "filter by month in YYYY-MM format")
	fs.Usage = func() {
		fmt.Println("Usage: expense-tracker summary [-category <category>] [-month YYYY-MM]")
		fs.PrintDefaults()
	}
	fs.Parse(args)

	items, err := store.Load()
	if err != nil {
		exitWithError(err)
	}

	if *category != "" {
		items = expense.FilterByCategory(items, *category)
	}

	if *month != "" {
		items = expense.FilterByMonth(items, *month)
	}

	summary := expense.BuildSummary(items)

	fmt.Printf("Number of expenses: %d\n", summary.Count)
	fmt.Printf("Total spent:        %.2f\n", summary.Total)

	if summary.Count == 0 {
		return
	}

	fmt.Println("\nBy category:")
	for category, total := range summary.ByCategory {
		fmt.Printf("  %-16s %.2f\n", category, total)
	}
}

func truncate(value string, max int) string {
	if len(value) <= max {
		return value
	}
	return value[:max-3] + "..."
}

func printUsage() {
	fmt.Println(`Expense Tracker

A simple command-line expense tracker written in Go.

Usage:
  expense-tracker <command> [options]

Commands:
  add       Add a new expense
  list      List expenses
  delete    Delete an expense by ID
  summary   Show spending summary
  help      Show this help

Examples:
  expense-tracker add -amount 250 -category Food -description "Lunch"
  expense-tracker add -amount 1200 -category Travel -description "Metro pass" -date 2026-10-02
  expense-tracker list
  expense-tracker list -category Food
  expense-tracker delete -id 1
  expense-tracker summary
  expense-tracker summary -category Food
  expense-tracker summary -month 2026-10`)
}

func exitWithError(err error) {
	fmt.Fprintln(os.Stderr, "Error:", err)
	os.Exit(1)
}
