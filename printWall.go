package main

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

	wallTop := lipgloss.NewStyle().
		Background(fg).
		Foreground(bg).
		// MarginLeft(6).
		Width(37).
		Align(lipgloss.Center).
		Render("Your wallah is")

	wall := lipgloss.NewStyle().
		Bold(true).
		Background(bg).
		Foreground(fg).
		PaddingTop(2).
		PaddingBottom(2).
		MarginLeft(6).
		Width(25).
		Align(lipgloss.Center).
		Render(colorName)

	fmt.Printf("\n%s\n\n%s\n\n", wallTop, wall)
}
