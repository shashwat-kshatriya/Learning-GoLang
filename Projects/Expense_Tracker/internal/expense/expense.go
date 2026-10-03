package expense

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type Expense struct {
	ID          int       `json:"id"`
	Amount      float64   `json:"amount"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

type Summary struct {
	Count      int
	Total      float64
	ByCategory map[string]float64
}

func New(amount float64, category, description string, date time.Time) (Expense, error) {
	if amount <= 0 {
		return Expense{}, fmt.Errorf("amount must be greater than zero")
	}

	category = strings.TrimSpace(category)
	if category == "" {
		return Expense{}, fmt.Errorf("category cannot be empty")
	}

	description = strings.TrimSpace(description)
	if description == "" {
		return Expense{}, fmt.Errorf("description cannot be empty")
	}

	return Expense{
		Amount:      amount,
		Category:    category,
		Description: description,
		Date:        date,
	}, nil
}

func NextID(expenses []Expense) int {
	maxID := 0
	for _, item := range expenses {
		if item.ID > maxID {
			maxID = item.ID
		}
	}
	return maxID + 1
}

func FilterByCategory(expenses []Expense, category string) []Expense {
	category = strings.TrimSpace(category)
	result := make([]Expense, 0)

	for _, item := range expenses {
		if strings.EqualFold(item.Category, category) {
			result = append(result, item)
		}
	}

	return result
}

func FilterByMonth(expenses []Expense, month string) []Expense {
	result := make([]Expense, 0)

	for _, item := range expenses {
		if item.Date.Format("2006-01") == month {
			result = append(result, item)
		}
	}

	return result
}

func DeleteByID(expenses []Expense, id int) ([]Expense, bool) {
	for i, item := range expenses {
		if item.ID == id {
			result := append([]Expense{}, expenses[:i]...)
			result = append(result, expenses[i+1:]...)
			return result, true
		}
	}

	return expenses, false
}

func BuildSummary(expenses []Expense) Summary {
	summary := Summary{
		Count:      len(expenses),
		ByCategory: make(map[string]float64),
	}

	for _, item := range expenses {
		summary.Total += item.Amount
		summary.ByCategory[item.Category] += item.Amount
	}

	return summary
}

func SortByDateDescending(expenses []Expense) {
	sort.Slice(expenses, func(i, j int) bool {
		return expenses[i].Date.After(expenses[j].Date)
	})
}
