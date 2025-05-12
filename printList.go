// This file contains the function to showcase the list of colors available in a colored and organized way, by using the -l or --list flags

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

// printColorsInColumns prints the color names styled in columns
func printList() {
	// Get terminal width
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		width = 50 // fallback width
	}

	var listTitle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#a2bcd0")).
		// Background(lipgloss.Color("#242521")).
		PaddingTop(1).
		PaddingBottom(1).
		Align(lipgloss.Center).
		Width(60)

	// fmt.Println("")
	fmt.Println(listTitle.Render("--- colors available ---"))
	// fmt.Println("")

	// Style each color label with lipgloss
	var styledLabels []string
	maxLabelWidth := 0
	for _, name := range wallpaperOrder {
		fg, bg := "#fff", "#000" // default colors
		if c, ok := colorStyles[name]; ok {
			fg, bg = c[0], c[1]
		}
		style := lipgloss.NewStyle().
			// Bold(true).
			Foreground(lipgloss.Color(fg)).
			Background(lipgloss.Color(bg)).
			Align(lipgloss.Center).
			Width(15)

		label := style.Render(name)
		styledLabels = append(styledLabels, label)

		plain := lipgloss.NewStyle().Render(name)
		if len(plain) > maxLabelWidth {
			maxLabelWidth = len(plain)
		}
	}

	colWidth := maxLabelWidth + 6 // spacing between columns
	numCols := width / colWidth
	if numCols == 0 {
		numCols = 1
	}

	// Clamp to max 4 columns
	if numCols > 4 {
		numCols = 4
	}

	for i, label := range styledLabels {
		fmt.Print(label)

		if (i+1)%numCols == 0 {
			fmt.Println()
		} else {
			spaces := colWidth - lipgloss.Width(label)
			fmt.Print(strings.Repeat(" ", spaces))
		}
	}
	fmt.Println()
}
