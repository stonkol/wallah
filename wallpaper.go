package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

func changeWallpaper(color string) error {
	wallpapers := map[string]string{
		"b-black":   "/wallpapers/mbp-14/b-black.png",
		"b-blue":    "/wallpapers/mbp-14/b-blue.png",
		"b-cyan":    "/wallpapers/mbp-14/b-cyan.png",
		"b-green":   "/wallpapers/mbp-14/b-green.png",
		"b-magenta": "/wallpapers/mbp-14/b-magenta.png",
		"b-red":     "/wallpapers/mbp-14/b-red.png",
		"b-white":   "/wallpapers/mbp-14/b-white.png",

		"black":   "/wallpapers/mbp-14/black.png",
		"blue":    "/wallpapers/mbp-14/blue.png",
		"cyan":    "/wallpapers/mbp-14/cyan.png",
		"green":   "/wallpapers/mbp-14/green.png",
		"magenta": "/wallpapers/mbp-14/magenta.png",
		"red":     "/wallpapers/mbp-14/red.png",
		"white":   "/wallpapers/mbp-14/white.png",
	}

	// Look up the image file path corresponding
	// to the given color in the wallpapers map
	imgPath, ok := wallpapers[color]
	// Check if the color exists in the map
	if !ok {
		return fmt.Errorf("color not found: %s 😅", color)
	}

	// Convert the image path to an absolute path
	absPath, err := filepath.Abs(imgPath)
	if err != nil {
		// If there is an error getting the absolute path
		return err
		// return fmt.Errorf("failed to get absolute path: %w", err)
	}

	// Prepare an AppleScript command as a string
	// that tells macOS to set the wallpaper
	// POSIX file "%s" converts the Go absolute path
	// into a format AppleScript understands
	script := fmt.Sprintf(`tell application "System Events" to set picture of every desktop to POSIX file "%s"`, absPath)
	// run AppleScript command using the osascript cli
	cmd := exec.Command("osascript", "-e", script)

	return cmd.Run()
}
