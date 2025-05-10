package main

import (
	"embed"
	"fmt"
	"os"
	"os/exec"
)

//go:embed wallpapers/mbp-14/*
var wallpaperFS embed.FS

var Wallpapers = map[string]string{
	// ansi 16
	"black":   "wallpapers/mbp-14/black.png",
	"blue":    "wallpapers/mbp-14/blue.png",
	"cyan":    "wallpapers/mbp-14/cyan.png",
	"green":   "wallpapers/mbp-14/green.png",
	"magenta": "wallpapers/mbp-14/magenta.png",
	"red":     "wallpapers/mbp-14/red.png",
	"white":   "wallpapers/mbp-14/white.png",
	"yellow":  "wallpapers/mbp-14/yellow.png",

	"black-b":   "wallpapers/mbp-14/black-b.png",
	"blue-b":    "wallpapers/mbp-14/blue-b.png",
	"cyan-b":    "wallpapers/mbp-14/cyan-b.png",
	"green-b":   "wallpapers/mbp-14/green-b.png",
	"magenta-b": "wallpapers/mbp-14/magenta-b.png",
	"red-b":     "wallpapers/mbp-14/red-b.png",
	"yellow-b":  "wallpapers/mbp-14/yellow-b.png",
	"white-b":   "wallpapers/mbp-14/white-b.png",

	// BONUS
	"orange":   "wallpapers/mbp-14/orange.png",
	"orange-b": "wallpapers/mbp-14/orange-b.png",

	// char 16
	"akira":   "wallpapers/mbp-14/akira.png",
	"blender": "wallpapers/mbp-14/blender.png",
	"kirby":   "wallpapers/mbp-14/kirby.png",
	"pikachu": "wallpapers/mbp-14/pikachu.png",
	"stitchy": "wallpapers/mbp-14/stitchy.png",
	"teddy":   "wallpapers/mbp-14/teddy.png",
	"wario":   "wallpapers/mbp-14/wario.png",
	"yoshi":   "wallpapers/mbp-14/yoshi.png",

	"akira-b":   "wallpapers/mbp-14/akira-b.png",
	"blender-b": "wallpapers/mbp-14/blender-b.png",
	"kirby-b":   "wallpapers/mbp-14/kirby-b.png",
	"pikachu-b": "wallpapers/mbp-14/pikachu-b.png",
	"stitchy-b": "wallpapers/mbp-14/stitchy-b.png",
	"teddy-b":   "wallpapers/mbp-14/teddy-b.png",
	"wario-b":   "wallpapers/mbp-14/wario-b.png",
	"yoshi-b":   "wallpapers/mbp-14/yoshi-b.png",
}

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

	// err := cmd.Run()

	return nil
}
