package settings

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Config struct {
	Display                 string `json:"display"`
	MenuBarColor            string `json:"menu_bar_color"`
	MenuBarMatchesWallpaper bool   `json:"transparent_menu_bar,omitempty"`
	CornerStyle             string `json:"corner_style"`
}

var displays = map[string]struct{}{
	"mbp-14":    {},
	"mbp-16":    {},
	"mba-13":    {},
	"mba-15":    {},
	"imac-24":   {},
	"studio-27": {},
}

func (c *Config) UnmarshalJSON(data []byte) error {
	type configAlias Config
	var decoded configAlias
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	if decoded.Display == "" {
		var legacy struct {
			Device string `json:"device"`
		}
		if err := json.Unmarshal(data, &legacy); err != nil {
			return err
		}
		decoded.Display = legacy.Device
	}
	*c = Config(decoded)
	return nil
}

func Path() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("find user config directory: %w", err)
	}
	return filepath.Join(dir, "wallah", "settings.json"), nil
}

func Load() (Config, bool, error) {
	path, err := Path()
	if err != nil {
		return Config{}, false, err
	}
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return Config{}, false, nil
	}
	if err != nil {
		return Config{}, false, fmt.Errorf("read settings: %w", err)
	}
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return Config{}, false, fmt.Errorf("decode settings: %w", err)
	}
	if err := config.Validate(); err != nil {
		return Config{}, false, fmt.Errorf("invalid saved settings: %w", err)
	}
	return config, true, nil
}

func (c Config) Validate() error {
	if _, ok := displays[c.Display]; !ok {
		return fmt.Errorf("unknown display %q", c.Display)
	}
	hexColor := strings.TrimPrefix(c.MenuBarColor, "#")
	if len(hexColor) != 6 {
		return fmt.Errorf("menu bar color must be a 6 digit hex value")
	}
	if _, err := strconv.ParseUint(hexColor, 16, 24); err != nil {
		return fmt.Errorf("invalid menu bar color %q", c.MenuBarColor)
	}
	if c.CornerStyle != "macos26" && c.CornerStyle != "macos15" && c.CornerStyle != "macosx" {
		return fmt.Errorf("unknown corner style %q", c.CornerStyle)
	}
	return nil
}

func Save(config Config) error {
	if err := config.Validate(); err != nil {
		return err
	}
	path, err := Path()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return fmt.Errorf("create settings directory: %w", err)
	}
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("encode settings: %w", err)
	}
	tmpFile, err := os.CreateTemp(filepath.Dir(path), ".settings-*.json")
	if err != nil {
		return fmt.Errorf("create temporary settings file: %w", err)
	}
	tmpName := tmpFile.Name()
	defer os.Remove(tmpName)
	if _, err := tmpFile.Write(data); err != nil {
		tmpFile.Close()
		return fmt.Errorf("write settings: %w", err)
	}
	if err := tmpFile.Close(); err != nil {
		return fmt.Errorf("close settings: %w", err)
	}
	if err := os.Rename(tmpName, path); err != nil {
		return fmt.Errorf("save settings: %w", err)
	}
	return nil
}
