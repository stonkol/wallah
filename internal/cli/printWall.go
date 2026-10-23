// This file contains the logic for printing the wallah, which is a preview of the color of the wallpaper in a more graphic way

package cli

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func printWall(colorName string) {
	colors, ok := colorStyles[colorName]
	if !ok {
		// fallback: print plain text if color not found
		fmt.Printf("Your wallah is %s 🤟\n", colorName)
		return
	}

	fg := lipgloss.Color(colors[0])
	bg := lipgloss.Color(colors[1])
	black := lipgloss.Color("#000")

	wallahTopText := lipgloss.NewStyle().
		Foreground(bg).
		MarginLeft(6).
		Width(31).
		Align(lipgloss.Center).
		Render("your wallah is")

	// holes := []rune{'○', '。', '●', '◯', '◎', '◉', '◌', '⚫', '⚪', '🕳'}
	// make another version for macOS 26 style
	wallahMenuBar := lipgloss.NewStyle().
		Background(black).
		Foreground(bg).
		MarginLeft(6).
		Width(31).
		Align(lipgloss.Center).
		Render("●")

	// print a preview of the color of the wallpaper
	wallah := lipgloss.NewStyle().
		Bold(true).
		Background(bg).
		Foreground(fg).
		PaddingTop(3).
		PaddingBottom(3).
		MarginLeft(6).
		Width(31).
		Align(lipgloss.Center).
		Render(colorName)

	fmt.Printf("\n%s\n%s", wallahTopText, wallahMenuBar)
	fmt.Printf("\n%s\n\n\n", wallah)
}
