package storage

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"expense-tracker/internal/expense"
)

type JSONStore struct {
	Path string
}

func NewJSONStore(path string) *JSONStore {
	return &JSONStore{Path: path}
}

func (s *JSONStore) Ensure() error {
	dir := filepath.Dir(s.Path)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	if _, err := os.Stat(s.Path); errors.Is(err, os.ErrNotExist) {
		return s.Save([]expense.Expense{})
	}

	return nil
}

func (s *JSONStore) Load() ([]expense.Expense, error) {
	data, err := os.ReadFile(s.Path)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return []expense.Expense{}, nil
	}

	var expenses []expense.Expense
	if err := json.Unmarshal(data, &expenses); err != nil {
		return nil, err
	}

	expense.SortByDateDescending(expenses)
	return expenses, nil
}

func (s *JSONStore) Save(expenses []expense.Expense) error {
	if err := os.MkdirAll(filepath.Dir(s.Path), 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(expenses, "", "  ")
	if err != nil {
		return err
	}

	tempPath := s.Path + ".tmp"
	if err := os.WriteFile(tempPath, data, 0644); err != nil {
		return err
	}

	if err := os.Rename(tempPath, s.Path); err != nil {
		_ = os.Remove(tempPath)
		return err
	}

	return nil
}
