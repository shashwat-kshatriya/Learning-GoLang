package expense

import (
	"testing"
	"time"
)

func TestNewRejectsInvalidExpense(t *testing.T) {
	date := time.Date(2026, 10, 1, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name        string
		amount      float64
		category    string
		description string
	}{
		{"zero amount", 0, "Food", "Lunch"},
		{"negative amount", -10, "Food", "Lunch"},
		{"empty category", 100, "", "Lunch"},
		{"empty description", 100, "Food", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := New(tt.amount, tt.category, tt.description, date); err == nil {
				t.Fatal("expected validation error")
			}
		})
	}
}

func TestNextID(t *testing.T) {
	expenses := []Expense{
		{ID: 1},
		{ID: 7},
		{ID: 3},
	}

	if got := NextID(expenses); got != 8 {
		t.Fatalf("expected 8, got %d", got)
	}
}

func TestDeleteByID(t *testing.T) {
	expenses := []Expense{
		{ID: 1, Amount: 100},
		{ID: 2, Amount: 200},
	}

	updated, deleted := DeleteByID(expenses, 1)

	if !deleted {
		t.Fatal("expected expense to be deleted")
	}
	if len(updated) != 1 {
		t.Fatalf("expected 1 expense, got %d", len(updated))
	}
	if updated[0].ID != 2 {
		t.Fatalf("expected remaining expense ID 2, got %d", updated[0].ID)
	}
}

func TestBuildSummary(t *testing.T) {
	expenses := []Expense{
		{Amount: 100, Category: "Food"},
		{Amount: 250, Category: "Travel"},
		{Amount: 50, Category: "Food"},
	}

	summary := BuildSummary(expenses)

	if summary.Count != 3 {
		t.Fatalf("expected count 3, got %d", summary.Count)
	}

	if summary.Total != 400 {
		t.Fatalf("expected total 400, got %.2f", summary.Total)
	}

	if summary.ByCategory["Food"] != 150 {
		t.Fatalf("expected Food total 150, got %.2f", summary.ByCategory["Food"])
	}

	if summary.ByCategory["Travel"] != 250 {
		t.Fatalf("expected Travel total 250, got %.2f", summary.ByCategory["Travel"])
	}
}
