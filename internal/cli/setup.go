package cli

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/stonkol/wallah/internal/settings"
	"github.com/stonkol/wallah/wallpapers"
)

func promptSettings(in io.Reader, out io.Writer) (settings.Config, error) {
	reader := bufio.NewReader(in)
	fmt.Fprintln(out, "Welcome to Wallah.\n")

	device, err := promptChoice(in, out, "Select your Mac:", []string{
		"MacBook Pro 14\"",
		"MacBook Pro 16\"",
		"MacBook Air 13\"",
		"MacBook Air 15\"",
	})
	if err != nil {
		return settings.Config{}, err
	}
	deviceIDs := map[string]string{
		"MacBook Pro 14\"": "mbp-14",
		"MacBook Pro 16\"": "mbp-16",
		"MacBook Air 13\"": "mba-13",
		"MacBook Air 15\"": "mba-15",
	}

	menuChoice, err := promptChoice(in, out, "Choose the menu bar color:", []string{
		"Black", "White", "Choose a color",
	})
	if err != nil {
		return settings.Config{}, err
	}
	menuBarColor := "#000000"
	switch menuChoice {
	case "White":
		menuBarColor = "#ffffff"
	case "Choose a color":
		for {
			fmt.Fprint(out, "Enter a color name (for example blue) or a hex color (#RRGGBB): ")
			value, err := readLine(reader)
			if err != nil {
				return settings.Config{}, err
			}
			menuBarColor, err = normalizeMenuColor(value)
			if err == nil {
				break
			}
			fmt.Fprintf(out, "%v\n", err)
		}
	}

	cornerChoice, err := promptChoice(in, out, "Choose the screen corner style:", []string{
		"macOS 26 (rounder)", "macOS 15",
	})
	if err != nil {
		return settings.Config{}, err
	}
	cornerStyle := "macos26"
	if cornerChoice == "macOS 15" {
		cornerStyle = "macos15"
	}

	return settings.Config{
		Device:       deviceIDs[device],
		MenuBarColor: menuBarColor,
		CornerStyle:  cornerStyle,
	}, nil
}

func promptChoice(in io.Reader, out io.Writer, title string, options []string) (string, error) {
	executable, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("find Wallah executable for setup picker: %w", err)
	}
	args := []string{"__wallah-choose", title}
	args = append(args, options...)
	cmd := exec.Command(executable, args...)
	cmd.Stdin = in
	cmd.Stderr = out
	selection, err := cmd.Output()
	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) && exitErr.ExitCode() == 1 {
			return "", fmt.Errorf("setup selection canceled")
		}
		return "", fmt.Errorf("run bundled Gum picker: %w", err)
	}
	choice := strings.TrimSpace(string(selection))
	for _, option := range options {
		if choice == option {
			return choice, nil
		}
	}
	return "", fmt.Errorf("gum returned an unknown choice %q", choice)
}

func readLine(reader *bufio.Reader) (string, error) {
	line, err := reader.ReadString('\n')
	if err != nil && len(line) == 0 {
		return "", err
	}
	return strings.TrimSpace(line), nil
}

func normalizeMenuColor(value string) (string, error) {
	value = strings.ToLower(strings.TrimSpace(value))
	if color, ok := wallpapers.Color(value); ok {
		return color.Background, nil
	}
	hex := strings.TrimPrefix(value, "#")
	if len(hex) == 3 {
		hex = string([]byte{hex[0], hex[0], hex[1], hex[1], hex[2], hex[2]})
	}
	if len(hex) != 6 {
		return "", fmt.Errorf("use a known color name or a 3 or 6 digit hex color")
	}
	if _, err := strconv.ParseUint(hex, 16, 24); err != nil {
		return "", fmt.Errorf("invalid hex color %q", value)
	}
	return "#" + hex, nil
}
