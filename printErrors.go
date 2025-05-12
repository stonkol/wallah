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

// func printErrorMessage(errorMessage string) {

// 	errorTitle := lipgloss.NewStyle().
// 		Foreground(lipgloss.Color("#f7fcfe")).
// 		Background(lipgloss.Color("#db0f27")).
// 		Padding(0, 1).
// 		Render("ERROR")

// 	errorStyle := lipgloss.NewStyle().
// 		Foreground(lipgloss.Color("#db0f27"))

// 	fullMessage := fmt.Sprintf("%s %s", errorTitle, errorStyle.Render(errorMessage))

// 	fmt.Println(fullMessage)
// }

// func printTwoErrorMessages(errorMessage string, error error) {

// 	errorTitle := lipgloss.NewStyle().
// 		Foreground(lipgloss.Color("#f7fcfe")).
// 		Background(lipgloss.Color("#db0f27")).
// 		Padding(0, 1).
// 		Render("ERROR")

// 	errorStyle := lipgloss.NewStyle().
// 		Foreground(lipgloss.Color("#db0f27"))

// 	fullMessage := fmt.Sprintf("%s %s %s", errorTitle, errorStyle.Render(errorMessage), errorStyle.Render(error))

// 	fmt.Println(fullMessage)
// }
