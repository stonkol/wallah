package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/stonkol/wallah/wallpapers"
)

type colorHistory struct {
	Current  string `json:"current"`
	Previous string `json:"previous"`
}

func historyPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("find user config directory: %w", err)
	}
	return filepath.Join(configDir, "wallah", "color-history.json"), nil
}

func loadColorHistory() (colorHistory, error) {
	path, err := historyPath()
	if err != nil {
		return colorHistory{}, err
	}
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return colorHistory{}, nil
	}
	if err != nil {
		return colorHistory{}, fmt.Errorf("read color history: %w", err)
	}

	var history colorHistory
	if err := json.Unmarshal(data, &history); err != nil {
		return colorHistory{}, fmt.Errorf("decode color history: %w", err)
	}
	history.Current = migrateColorSuffix(history.Current)
	history.Previous = migrateColorSuffix(history.Previous)
	return history, nil
}

func migrateColorSuffix(name string) string {
	if !strings.HasSuffix(name, "-b") {
		return name
	}
	updated := strings.TrimSuffix(name, "-b") + "-l"
	if _, ok := wallpapers.Color(updated); ok {
		return updated
	}
	return name
}

func recordColor(color string) error {
	history, err := loadColorHistory()
	if err != nil {
		return err
	}
	if history.Current != color {
		history.Previous = history.Current
		history.Current = color
	}

	path, err := historyPath()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return fmt.Errorf("create config directory: %w", err)
	}

	data, err := json.Marshal(history)
	if err != nil {
		return fmt.Errorf("encode color history: %w", err)
	}
	tmpFile, err := os.CreateTemp(filepath.Dir(path), ".color-history-*.json")
	if err != nil {
		return fmt.Errorf("create temporary history file: %w", err)
	}
	tmpName := tmpFile.Name()
	defer os.Remove(tmpName)

	if _, err := tmpFile.Write(data); err != nil {
		tmpFile.Close()
		return fmt.Errorf("write color history: %w", err)
	}
	if err := tmpFile.Close(); err != nil {
		return fmt.Errorf("close color history: %w", err)
	}
	if err := os.Rename(tmpName, path); err != nil {
		return fmt.Errorf("save color history: %w", err)
	}
	return nil
}
