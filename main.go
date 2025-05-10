package main

import (
	"fmt"
	"os"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var version = "0.2.0"

var colorStyles = map[string]struct{ Fg, Bg string }{

	// ansi 16
	"black":     {"#FFFFFF", "#000000"},
	"blue":      {"#FFFFFF", "#0000FF"},
	"cyan":      {"#000000", "#00FFFF"},
	"green":     {"#000000", "#00FF00"},
	"magenta":   {"#FFFFFF", "#FF00FF"},
	"red":       {"#FFFFFF", "#FF0000"},
	"yellow":    {"#000000", "#FFFF00"},
	"white":     {"#000000", "#FFFFFF"},
	"black-b":   {"#FFFFFF", "#7F7F7F"},
	"blue-b":    {"#FFFFFF", "#00007F"},
	"cyan-b":    {"#000000", "#007F7F"},
	"green-b":   {"#000000", "#007F00"},
	"magenta-b": {"#FFFFFF", "#7F007F"},
	"red-b":     {"#FFFFFF", "#7F0000"},
	"yellow-b":  {"#000000", "#7F7F00"},
	"white-b":   {"#000000", "#C0C0C0"},

	// bonus
	"orange":   {"#000000", "#FF5516"},
	"orange-b": {"#000000", "#FF7E23"},

	// char 16
	"akira":     {"#000000", "#AB2F1A"},
	"blender":   {"#000000", "#A2BCD0"},
	"grimace":   {"#000000", "#6A1C63"},
	"kirby":     {"#000000", "#FC5C8E"},
	"pikachu":   {"#000000", "#F8A21C"},
	"stitchy":   {"#000000", "#527BAC"},
	"teddy":     {"#000000", "#925923"},
	"yoshi":     {"#000000", "#43A934"},
	"akira-b":   {"#ffffff", "#DF2626"},
	"blender-b": {"#ffffff", "#C2D6E2"},
	"grimace-b": {"#ffffff", "#8E1E8C"},
	"kirby-b":   {"#ffffff", "#F49792"},
	"pikachu-b": {"#ffffff", "#F4DB06"},
	"stitchy-b": {"#ffffff", "#69B8DC"},
	"teddy-b":   {"#ffffff", "#D09249"},
	"yoshi-b":   {"#ffffff", "#69C12B"},
}

func printColoredLabel(colorName string) {
	style, ok := colorStyles[colorName]
	if !ok {
		// fallback: print plain text if color not found
		fmt.Printf("Your wallah is %s 🤟\n", colorName)
		return
	}

	// Compose the output line with inline tags
	// Only the colorName word is styled with fg/bg and padded spaces
	output := "Your wallah is <fg=" + style.Fg + ";bg=" + style.Bg + "> " + colorName + " </> 🤟"
	color.Println(output)
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
				"cyan",
				"green",
				"magenta",
				"red",
				"white",
				"black-b",
				"blue-b",
				"cyan-b",
				"green-b",
				"magenta-b",
				"red-b",
				"white-b",

				"akira",
				"blender",
				"grimace",
				"kirby",
				"pikachu",
				"stitchy",
				"teddy",
				"yoshi",

				"akira-b",
				"blender-b",
				"grimace-b",
				"kirby-b",
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
				// Iterate over the slice to print colors in order
				for _, colorName := range wallpaperOrder {
					fmt.Println(" -", colorName)
				}
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
			// fmt.Println("  Your wallah is now", color, "🤟")
			// fmt.Print("Your wallah is ")
			// printColoredLabel("red-b")
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
