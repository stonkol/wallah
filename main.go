package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	// Define the root command for the CLI app using Cobra
	var rootCmd = &cobra.Command{
		Use:   "wall",                                                               // The command name users will type (e.g., `wall`)
		Short: "Wall is a CLI tool for setting wallpapers that hides the top notch", // A short description
		Long:  "Wall is a CLI tool for setting MacBooks' wallpapers. It set a wallpaper with a solid color, rounded borders ",
		Args:  cobra.ExactArgs(1), // Enforce that exactly one argument must be provided (like `blue`)

		// The function to run when the command is executed
		Run: func(cmd *cobra.Command, args []string) {
			color := args[0] // retrieve the first argument (the color name)

			// Call the function to change the wallpaper using the provided color
			err := changeWallpaper(color)
			if err != nil {
				fmt.Println(" 🥲 Error changing wallpaper:", err)
				os.Exit(1)
			}

			// If successful, print
			fmt.Println("  Your wall is now", color, "🤟")
		},
	}
	// Execute the root command, which parses arguments and runs the Run function
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error executing command:", err)
		os.Exit(1)
	}
}
