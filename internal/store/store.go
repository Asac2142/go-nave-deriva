// Package store handles the read & write operations to a JSON file.
package store

import (
	"encoding/json"
	"fmt"
	"os"
)

const storage = "db.json"

// Read attemps to read a JSON file in which it stores the "Nave a la deriva" status value.
func Read[T any]() (*T, error) {
	bytes, err := os.ReadFile(storage)
	if err != nil {
		return nil, fmt.Errorf("read storage: %w", err)
	}

	var data T
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, fmt.Errorf("unmarshal data: %w", err)
	}

	return &data, nil
}

// Write attemps to write into a JSON file a status value from the "Nave a la deriva" scenario.
func Write[T any](data T) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("marshal data: %w", err)
	}

	err = os.WriteFile(storage, bytes, 0644)
	if err != nil {
		return fmt.Errorf("write storage: %w", err)
	}

	return nil
}
