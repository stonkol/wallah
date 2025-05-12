// the flags logic and cobra usage are located here

package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var version = "0.2.5 🍭"

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
		Short: "A fast and simple CLI to change your macOS wallpaper and elegantly hide that ugly notch.", // A short description

		Long: "\nHi I'm Wallah, I will help you to change your wallpaper to your desire color and get ride of that notch.\n\nWith wallah you can easily apply wallpapers featuring solid colors and rounded borders that blend seamlessly around the notch area.\n\nExample:\n  wallah [color]\t Set a wallpaper color\n  wallah random\t Randomly sets a wallpaper",

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
			randomFlag, err := cmd.Flags().GetBool("random")
			if err != nil {
				return err
			}

			if listFlag || versionFlag || randomFlag {
				if len(args) != 0 {
					printErrorMessage("No arguments allowed when using a flag.")
					return fmt.Errorf("no arguments allowed when using a flag.")
				}
				return nil
			}

			if len(args) != 1 {
				printErrorMessage("Requires exactly one argument")
				return fmt.Errorf("requires exactly one argument")
			}

			// Allow "random" as a positional argument
			if args[0] == "random" {
				return nil
			}

			return nil
		},

		// The function to run when the command is executed, logic only
		Run: func(cmd *cobra.Command, args []string) {

			// Check if --list flag is set
			listFlag, err := cmd.Flags().GetBool("list")
			if err != nil {
				printErrorMessage("problem for reading flags:")
				printErr(err)
				os.Exit(1)
			}

			// Check if --version flag is set
			versionFlag, err := cmd.Flags().GetBool("version")
			if err != nil {
				printErrorMessage("problem for reading flags:")
				printErr(err)
				os.Exit(1)
			}

			// Check if --random flag is set
			randomFlag, err := cmd.Flags().GetBool("random")
			if err != nil {
				printErrorMessage("problem for reading flags:")
				printErr(err)
				os.Exit(1)
			}

			if versionFlag {
				versionStyle := lipgloss.NewStyle().
					Foreground(lipgloss.Color("#ff7e23")) //orange

				versionLine := "version " + version
				fmt.Println(versionStyle.Render(versionLine))
				return
			}

			if listFlag {
				printList()
				return
			}

			if randomFlag || (len(args) > 0 && args[0] == "random") {
				rand.Seed(time.Now().UnixNano())
				randomColor := wallpaperOrder[rand.Intn(len(wallpaperOrder))]

				err := changeWallpaper(randomColor)
				if err != nil {
					fullMessage := fmt.Sprintf("Error changing the wallah to '%s'.", randomColor)
					printErrorMessage(fullMessage)
					printErr(err)
					os.Exit(1)
				}

				printWall(randomColor)
				return
			}

			//////////////// INPUT CHECK //////////////
			// Expect exactly one positional argument (one color)
			if len(args) != 1 {
				printErrorMessage("You must specify exactly one color.")
				cmd.Help()
				os.Exit(1)
			}
			// retrieve the first argument (the color name)
			colorName := args[0]

			// Call the func to change the wallpaper using the provided color
			err = changeWallpaper(colorName)
			if err != nil {
				fullMessage := fmt.Sprintf("Error changing the wallah to '%s'.", colorName)
				printErrorMessage(fullMessage)
				printErr(err)
				os.Exit(1)
			}

			// If successful, print
			printWall(colorName)
		},
	}

	// Add flags
	rootCmd.Flags().BoolP("list", "l", false, "list available colors")

	rootCmd.Flags().BoolP("version", "v", false, "current version")

	rootCmd.Flags().BoolP("random", "r", false, "set a random wallpaper color")

	// Set custom help template
	rootCmd.SetHelpTemplate(customHelpTemplate)

	// Execute the root command, which parses arguments and runs the Run function
	if err := rootCmd.Execute(); err != nil {
		printErrorMessage("Error executing command: ")
		printErr(err)
		os.Exit(1)
	}
}
