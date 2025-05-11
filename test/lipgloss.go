package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func testLipGloss() {
	// lipgloss.Color("5")       // magenta
	// lipgloss.Color("201")     // hot pink
	// lipgloss.Color("#04B575") // a green

	// Define a style with foreground and background colors
	redBg := lipgloss.NewStyle().
		Background(lipgloss.Color("#800000")).
		Foreground(lipgloss.Color("#ffffff")).
		Padding(0, 1) // top/bottom, left/right padding

	// Render styled text with padding spaces
	styledText := redBg.Render("red-b")

	fmt.Printf("Your wallah is %s 🤟\n", styledText)

	// Lipgloss supports borders and rounded corners
	bordered := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#ff0000")).
		Padding(0, 1).
		Render("red-b")
	fmt.Println(bordered)

	var hello = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		PaddingTop(2).
		PaddingLeft(4).
		Width(22)
	fmt.Println(hello.Render("Hello, kitty"))

	fmt.Println("\n")

	var colorName = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		PaddingLeft(6).
		PaddingRight(6).
		PaddingTop(1).
		PaddingBottom(2).
		Width(25)
	fmt.Println(colorName.Render("Hello, Wallah"))

	style := lipgloss.NewStyle().Foreground(lipgloss.Color("219"))
	copiedStyle := style           // this is a true copy
	wildStyle := style.Blink(true) // this is also true copy, with blink added

	fmt.Println(copiedStyle.Render("copied style"))
	fmt.Println(wildStyle.Render("wild style"))

}
