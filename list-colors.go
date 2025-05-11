// print.go
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

// printColorsInColumns prints the color names styled in columns
func printColorsInColumns(colorNames []string, colors map[string][2]string) {
	// Get terminal width
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		width = 50 // fallback width
	}

	// Style each color label with lipgloss
	var styledLabels []string
	maxLabelWidth := 0
	for _, name := range colorNames {
		fg, bg := "#fff", "#000" // default colors
		if c, ok := colors[name]; ok {
			fg, bg = c[0], c[1]
		}
		style := lipgloss.NewStyle().
			Bold(true).
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
