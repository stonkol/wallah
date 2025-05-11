package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var version = "0.2.2"

var colorStyles = map[string][2]string{

	// lowercase for hex recommended for compatibility
	// 0 is black, f is white
	// ansi 16
	"black":     {"#fff", "#000000"},
	"blue":      {"#fff", "#0000ff"},
	"cyan":      {"#000", "#00cdcd"},
	"green":     {"#fff", "#00dd00"},
	"magenta":   {"#fff", "#cd00cd"},
	"red":       {"#fff", "#cd0000"},
	"yellow":    {"#000", "#cdcd00"},
	"white":     {"#000", "#ffffff"},
	"black-b":   {"#fff", "#7f7f7f"},
	"blue-b":    {"#fff", "#5c5cff"},
	"cyan-b":    {"#000", "#00ffff"},
	"green-b":   {"#000", "#00ff00"},
	"magenta-b": {"#fff", "#ff00ff"},
	"red-b":     {"#fff", "#ff0000"},
	"yellow-b":  {"#000", "#ffff00"},
	"white-b":   {"#000", "#e5e5e5"},

	// bonus
	"orange":   {"#000", "#ff5516"},
	"orange-b": {"#000", "#ff7e23"},

	// char 16
	"akira":     {"#fff", "#ab2f1a"},
	"blender":   {"#fff", "#a2bcd0"},
	"grimace":   {"#fff", "#6a1c63"},
	"kirby":     {"#fff", "#fc5c8e"},
	"pikachu":   {"#fff", "#f8a21c"},
	"stitchy":   {"#fff", "#527bac"},
	"teddy":     {"#fff", "#925923"},
	"yoshi":     {"#fff", "#43a934"},
	"akira-b":   {"#fff", "#df2626"},
	"blender-b": {"#000", "#c2d6e2"},
	"grimace-b": {"#fff", "#8e1e8c"},
	"kirby-b":   {"#fff", "#f49792"},
	"pikachu-b": {"#000", "#f4db06"},
	"stitchy-b": {"#fff", "#69b8dc"},
	"teddy-b":   {"#fff", "#d09249"},
	"yoshi-b":   {"#fff", "#69c12b"},
}

func printColoredLabel(colorName string) {
	colors, ok := colorStyles[colorName]
	if !ok {
		// fallback: print plain text if color not found
		fmt.Printf("Your wallah is %s 🤟\n", colorName)
		return
	}

	fg := lipgloss.Color(colors[0])
	bg := lipgloss.Color(colors[1])

	style := lipgloss.NewStyle().
		Bold(true).
		Background(bg).
		Foreground(fg).
		PaddingTop(2).
		PaddingBottom(2).
		Width(27).
		Align(lipgloss.Center).
		Render(colorName)

	// ⬇️🔽
	fmt.Printf("\n      Your wallah is\n")
	fmt.Printf("%s\n\n", style)
}

func main() {
	// custom help template, replace cobra's help
	const customHelpTemplate = `{{with (or .Long .Short)}}{{. | trimTrailingWhitespaces}}

{{end}}Usage:
  {{.UseLine}}

Flags:
{{.Flags.FlagUsages | trimTrailingWhitespaces}}
`

	// Define the root command for the CLI app using Cobra
	var rootCmd = &cobra.Command{
		// users will type `wall` on the Command Line
		Use:   "wallah",
		Short: "A fast and simple CLI to change your macOS wallpaper and elegantly hide that ugly notch 🌀", // A short description

		Long: "\nHi I'm Wallah, I will help you to change your wallpaper to your desire color and get ride of that notch.\n\nWith wallah you can easily apply wallpapers featuring solid colors and rounded borders that blend seamlessly around the notch area.\n\nExample:\n  wallah [color]\n",

		// validation only in args
		Args: func(cmd *cobra.Command, args []string) error {
			listFlag, err := cmd.Flags().GetBool("list")
			if err != nil {
				return err
			}
			versionFlag, err := cmd.Flags().GetBool("version")
			if err != nil {
				return err
			}

			if listFlag || versionFlag {
				if len(args) != 0 {
					return fmt.Errorf("no arguments allowed when using --list or --version")
				}
				return nil
			}

			if len(args) != 1 {
				return fmt.Errorf("requires exactly one argument")
			}
			return nil
		},

		// The function to run when the command is executed, logic only
		Run: func(cmd *cobra.Command, args []string) {

			/////////////// list colorsFLAG ///////////////////
			// a slice of the colors in order,
			// instead of output the map itself (which is random)
			var wallpaperOrder = []string{
				"black",
				"blue",
				"red",
				"green",
				"black-b",
				"blue-b",
				"red-b",
				"green-b",

				"yellow",
				"cyan",
				"magenta",
				"white",
				"yellow-b",
				"cyan-b",
				"magenta-b",
				"white-b",

				"akira",
				"blender",
				"grimace",
				"kirby",
				"akira-b",
				"blender-b",
				"grimace-b",
				"kirby-b",

				"pikachu",
				"stitchy",
				"teddy",
				"yoshi",
				"pikachu-b",
				"stitchy-b",
				"teddy-b",
				"yoshi-b",
			}

			// Check if --list flag is set
			listFlag, err := cmd.Flags().GetBool("list")
			if err != nil {
				fmt.Println("Error reading flags:", err)
				os.Exit(1)
			}

			// Check if --version flag is set
			versionFlag, err := cmd.Flags().GetBool("version")
			if err != nil {
				fmt.Println("Error reading flags:", err)
				os.Exit(1)
			}

			if versionFlag {
				fmt.Println("version", version, "🌺")
				return
			}
			if listFlag {
				printColorsInColumns(wallpaperOrder, colorStyles)
				return
			}

			//////////////// INPUT CHECK //////////////
			// Normal behavior: expect exactly one positional argument (color)
			if len(args) != 1 {
				fmt.Println("You must specify exactly one color. 🥲")
				cmd.Help()
				os.Exit(1)
			}
			// retrieve the first argument (the color name)
			colorName := args[0]

			// Call the func to change the wallpaper using the provided color
			err = changeWallpaper(colorName)
			if err != nil {
				fmt.Printf("Error changing to the color '%s' 🥲 %s\n", colorName, err)
				os.Exit(1)
			}

			// If successful, print
			printColoredLabel(colorName)
		},
	}

	// Add flags
	rootCmd.Flags().BoolP("list", "l", false, "list available colors")

	rootCmd.Flags().BoolP("version", "v", false, "current version:")

	// Set custom help template
	rootCmd.SetHelpTemplate(customHelpTemplate)

	// Execute the root command, which parses arguments and runs the Run function
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error executing command:", err)
		os.Exit(1)
	}
}
