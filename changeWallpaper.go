// This file contains the function to change the wallpaper from a embedded image located in the binaries. It run an AppleScript command using the osascript cli to set the wallpaper.

package main

import (
	"embed"
	"fmt"
	"os"
	"os/exec"
)

//go:embed wallpapers/mbp-14/*
var wallpaperFS embed.FS

func changeWallpaper(colorName string) error {

	// Look up the image file path corresponding
	// to the given color in the wallpapers map
	imgPath, ok := Wallpapers[colorName]

	// Check if the color exists in the map
	if !ok {
		return fmt.Errorf("\nThe color '%s' is not available 😅", colorName)
	}

	// Read the image file bytes from the embedded FS
	data, err := wallpaperFS.ReadFile(imgPath)
	if err != nil {
		return fmt.Errorf("\nFailed to read the embedded wallpaper file: %w 🫨", err)
	}

	// Create a temp file to write the wallpaper image
	tmpFile, err := os.CreateTemp("", "wallpaper-*.png")
	if err != nil {
		return fmt.Errorf("\nFailed to create temp file: %w 🫨", err)
	}
	// Closes the open file descriptor associated with tmpFile.
	// release the file handle and flush any buffered writes.
	defer tmpFile.Close()

	// Write the embedded image data to the temp file
	if _, err := tmpFile.Write(data); err != nil {
		return fmt.Errorf("\nFailed to write data to temp file: %w 🤧", err)
	}

	// Use the temp file path for the AppleScript command
	absPath := tmpFile.Name()

	// Prepare an AppleScript command as a string
	// that tells macOS to set the wallpaper
	// POSIX file "%s" converts the Go absolute path
	// into a format AppleScript understands
	script := fmt.Sprintf(`tell application "System Events" to set picture of every desktop to POSIX file "%s"`, absPath)

	// run AppleScript command using the osascript cli
	cmd := exec.Command("osascript", "-e", script)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("\nFailed to execute AppleScript: %w 🤧", err)
	}

	return nil
}
