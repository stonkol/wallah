package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	errorTitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#f7fcfe")).
			Background(lipgloss.Color("#db0f27")).
			Padding(0, 1)

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#db0f27"))
)

func printErrorMessage(errorMessage string) {
	fmt.Printf("%s", errorTitleStyle.Render("ERROR"))
	fmt.Printf(" %s ", errorStyle.Render(errorMessage))
}

func printErr(err error) {
	// convert error to string
	errString := err.Error()

	fmt.Printf("%s\n", errorStyle.Render(errString))
}
