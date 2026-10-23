package cli

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/stonkol/wallah/wallpapers"
)

const version = "0.2.7 🫯"

const customHelpTemplate = `{{with (or .Long .Short)}}{{. | trimTrailingWhitespaces}}

{{end}}Usage:
  {{.UseLine}}

Flags:
{{.Flags.FlagUsages | trimTrailingWhitespaces}}
`

// NewCommand builds the root CLI command.
func NewCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "wallah",
		Short: "A fast and simple CLI to change your macOS wallpaper and elegantly hide that ugly notch.",
		Long:  "\nHi, I'm Wallah. I will help you change your wallpaper to your desired color and get rid of that notch.\n\nWith Wallah, you can easily apply wallpapers featuring solid colors and rounded borders that blend seamlessly around the notch area.\n\n----- HOW TO USE -------\n\nExample:\n  wallah [color]    Set a wallpaper color\n  wallah random     Randomly sets a wallpaper\n  wallah -l or list List available colors\n  wallah -v or version Show the version",
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
					return fmt.Errorf("no arguments allowed when using a flag")
				}
				return nil
			}
			if len(args) > 1 {
				return fmt.Errorf("requires one argument")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			listFlag, err := cmd.Flags().GetBool("list")
			if err != nil {
				return fmt.Errorf("problem reading list flag: %w", err)
			}
			versionFlag, err := cmd.Flags().GetBool("version")
			if err != nil {
				return fmt.Errorf("problem reading version flag: %w", err)
			}
			randomFlag, err := cmd.Flags().GetBool("random")
			if err != nil {
				return fmt.Errorf("problem reading random flag: %w", err)
			}

			if versionFlag || (len(args) == 1 && args[0] == "version") {
				style := lipgloss.NewStyle().Foreground(lipgloss.Color("#ff7e23"))
				fmt.Println(style.Render("version " + version))
				return nil
			}
			if listFlag || (len(args) == 1 && args[0] == "list") {
				printList()
				return nil
			}

			colorName := ""
			if randomFlag || (len(args) == 1 && args[0] == "random") {
				rng := rand.New(rand.NewSource(time.Now().UnixNano()))
				colorName = wallpaperOrder[rng.Intn(len(wallpaperOrder))]
			} else if len(args) == 1 {
				colorName = args[0]
			} else {
				if err := cmd.Help(); err != nil {
					return err
				}
				return fmt.Errorf("you must specify one color")
			}

			if err := wallpapers.Change(colorName); err != nil {
				return fmt.Errorf("error changing the wallah to '%s': %w", colorName, err)
			}
			printWall(colorName)
			return nil
		},
	}

	rootCmd.Flags().BoolP("list", "l", false, "list available colors")
	rootCmd.Flags().BoolP("version", "v", false, "current version")
	rootCmd.Flags().BoolP("random", "r", false, "set a random wallpaper color")
	rootCmd.SetHelpTemplate(customHelpTemplate)
	rootCmd.SilenceErrors = true
	rootCmd.SilenceUsage = true
	return rootCmd
}

// Execute runs the CLI using the current process arguments.
func Execute() error {
	return NewCommand().Execute()
}
