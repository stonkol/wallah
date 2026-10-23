// This file contains the function to change the wallpaper from a embedded image located in the binaries. It run an AppleScript command using the osascript cli to set the wallpaper.

package wallpapers

import (
	"embed"
	"fmt"
	"os"
	"os/exec"
)

//go:embed mbp-14/*
var wallpaperFS embed.FS

var wallpaperFiles = map[string]string{
	"black": "mbp-14/black.png", "blue": "mbp-14/blue.png", "cyan": "mbp-14/cyan.png",
	"green": "mbp-14/green.png", "magenta": "mbp-14/magenta.png", "red": "mbp-14/red.png",
	"white": "mbp-14/white.png", "yellow": "mbp-14/yellow.png",
	"black-b": "mbp-14/black-b.png", "blue-b": "mbp-14/blue-b.png", "cyan-b": "mbp-14/cyan-b.png",
	"green-b": "mbp-14/green-b.png", "magenta-b": "mbp-14/magenta-b.png", "red-b": "mbp-14/red-b.png",
	"white-b": "mbp-14/white-b.png", "yellow-b": "mbp-14/yellow-b.png",
	"orange": "mbp-14/orange.png", "orange-b": "mbp-14/orange-b.png",
	"akira": "mbp-14/akira.png", "blender": "mbp-14/blender.png", "grimace": "mbp-14/grimace.png",
	"kirby": "mbp-14/kirby.png", "pikachu": "mbp-14/pikachu.png", "stitchy": "mbp-14/stitchy.png",
	"teddy": "mbp-14/teddy.png", "yoshi": "mbp-14/yoshi.png",
	"akira-b": "mbp-14/akira-b.png", "blender-b": "mbp-14/blender-b.png", "grimace-b": "mbp-14/grimace-b.png",
	"kirby-b": "mbp-14/kirby-b.png", "pikachu-b": "mbp-14/pikachu-b.png", "stitchy-b": "mbp-14/stitchy-b.png",
	"teddy-b": "mbp-14/teddy-b.png", "yoshi-b": "mbp-14/yoshi-b.png",
}

// Change sets the macOS wallpaper to one of the bundled color images.
func Change(colorName string) error {

	// Look up the image file path corresponding
	// to the given color in the wallpapers map
	imgPath, ok := wallpaperFiles[colorName]

	// Check if the color exists in the map
	if !ok {
		return fmt.Errorf("The color '%s' is not available 😅", colorName)
	}

	// Read the image file bytes from the embedded FS
	data, err := wallpaperFS.ReadFile(imgPath)
	if err != nil {
		return fmt.Errorf("Failed to read the embedded wallpaper file: %w 🫨", err)
	}

	// Create a temp file to write the wallpaper image
	tmpFile, err := os.CreateTemp("", "wallpaper-*.png")
	if err != nil {
		return fmt.Errorf("Failed to create temp file: %w 🫨", err)
	}
	// Closes the open file descriptor associated with tmpFile.
	// release the file handle and flush any buffered writes.
	defer tmpFile.Close()

	// Write the embedded image data to the temp file
	if _, err := tmpFile.Write(data); err != nil {
		return fmt.Errorf("Failed to write data to temp file: %w 🤧", err)
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
