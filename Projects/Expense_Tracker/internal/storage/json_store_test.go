package storage

import (
	"path/filepath"
	"testing"
	"time"

	"expense-tracker/internal/expense"
)

func TestJSONStoreSaveAndLoad(t *testing.T) {
	path := filepath.Join(t.TempDir(), "expenses.json")
	store := NewJSONStore(path)

	items := []expense.Expense{
		{
			ID:          1,
			Amount:      250.50,
			Category:    "Food",
			Description: "Lunch",
			Date:        time.Date(2026, 10, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	if err := store.Save(items); err != nil {
		t.Fatalf("save failed: %v", err)
	}

	got, err := store.Load()
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}

	if len(got) != 1 {
		t.Fatalf("expected 1 expense, got %d", len(got))
	}

	if got[0].Amount != 250.50 {
		t.Fatalf("expected amount 250.50, got %.2f", got[0].Amount)
	}
}
