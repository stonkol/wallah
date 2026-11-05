package cli

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode/utf8"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
	"github.com/stonkol/wallah/internal/settings"
	"github.com/stonkol/wallah/wallpapers"
)

var errSetupCanceled = errors.New("settings canceled")

type settingsMenuModel struct {
	config       settings.Config
	cursor       int
	section      int
	option       int
	focusOptions bool
	editingColor bool
	colorInput   string
	message      string
	canceled     bool
	width        int
}

var displayOptions = []string{
	"MacBook Pro 14\"",
	"MacBook Pro 16\"",
	"MacBook Air 13\"",
	"MacBook Air 15\"",
	"iMac 24\" 4.5K",
	"Studio Display 27\" 5K",
}

var menuBarOptions = []string{"Black", "White", "Transparent", "Custom"}
var cornerOptions = []string{"macOS 26 (round)", "macOS 15 (less round)", "macOS X (squared)"}

func promptSettings(in io.Reader, out io.Writer) (settings.Config, error) {
	config, configured, err := settings.Load()
	if err != nil {
		return settings.Config{}, err
	}
	if !configured {
		config = settings.Config{
			Display:      "mbp-14",
			MenuBarColor: "#000000",
			CornerStyle:  "macos26",
		}
	}

	model := settingsMenuModel{config: config, section: 0, width: 80}
	program := tea.NewProgram(model, tea.WithInput(in), tea.WithOutput(out))
	finalModel, err := program.Run()
	if errors.Is(err, tea.ErrInterrupted) {
		return settings.Config{}, errSetupCanceled
	}
	if err != nil {
		return settings.Config{}, fmt.Errorf("run settings menu: %w", err)
	}
	result, ok := finalModel.(settingsMenuModel)
	if !ok {
		return settings.Config{}, fmt.Errorf("settings menu returned an unexpected model")
	}
	if result.canceled {
		return settings.Config{}, errSetupCanceled
	}
	return result.config, nil
}

func (m settingsMenuModel) Init() tea.Cmd { return nil }

func (m settingsMenuModel) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	if size, ok := message.(tea.WindowSizeMsg); ok {
		m.width = size.Width
		return m, nil
	}
	key, ok := message.(tea.KeyPressMsg)
	if !ok {
		return m, nil
	}
	pressed := key.String()
	if pressed == "ctrl+c" {
		m.canceled = true
		return m, tea.Quit
	}

	if m.editingColor {
		switch pressed {
		case "enter":
			color, err := normalizeMenuColor(m.colorInput)
			if err != nil {
				m.message = "Invalid color. Try again."
				return m, nil
			}
			m.config.MenuBarColor = color
			m.config.MenuBarMatchesWallpaper = false
			m.editingColor = false
			m.focusOptions = false
			m.message = ""
		case "esc", "left":
			m.editingColor = false
			m.focusOptions = false
			m.message = ""
		case "backspace":
			if len(m.colorInput) > 0 {
				_, size := utf8.DecodeLastRuneInString(m.colorInput)
				m.colorInput = m.colorInput[:len(m.colorInput)-size]
			}
		default:
			if key.Mod == 0 {
				m.colorInput += key.Text
				m.message = ""
			}
		}
		return m, nil
	}

	if pressed == "q" {
		return m, tea.Quit
	}

	if m.focusOptions && m.section >= 0 {
		options := m.options()
		switch pressed {
		case "up", "k":
			if m.option > 0 {
				m.option--
			} else {
				m.option = len(options) - 1
			}
		case "down", "j":
			if m.option < len(options)-1 {
				m.option++
			} else {
				m.option = 0
			}
		case "esc", "left":
			m.focusOptions = false
		case "enter", "right":
			m.selectOption(options[m.option])
		}
		return m, nil
	}

	switch pressed {
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		} else {
			m.cursor = 3
		}
		m.updateSection()
	case "down", "j":
		if m.cursor < 3 {
			m.cursor++
		} else {
			m.cursor = 0
		}
		m.updateSection()
	case "enter", "right", "space":
		if m.cursor == 3 {
			return m, tea.Quit
		}
		m.section = m.cursor
		m.option = m.selectedOptionIndex()
		m.focusOptions = true
	case "esc":
		m.canceled = true
		return m, tea.Quit
	}
	return m, nil
}

func (m *settingsMenuModel) updateSection() {
	if m.cursor >= 3 {
		m.section = -1
		return
	}
	m.section = m.cursor
	m.option = m.selectedOptionIndex()
}

func (m *settingsMenuModel) selectOption(choice string) {
	switch m.section {
	case 0:
		displayIDs := map[string]string{
			"MacBook Pro 14\"":       "mbp-14",
			"MacBook Pro 16\"":       "mbp-16",
			"MacBook Air 13\"":       "mba-13",
			"MacBook Air 15\"":       "mba-15",
			"iMac 24\" 4.5K":         "imac-24",
			"Studio Display 27\" 5K": "studio-27",
		}
		m.config.Display = displayIDs[choice]
	case 1:
		switch choice {
		case "Black":
			m.config.MenuBarColor = "#000000"
			m.config.MenuBarMatchesWallpaper = false
		case "White":
			m.config.MenuBarColor = "#ffffff"
			m.config.MenuBarMatchesWallpaper = false
		case "Transparent":
			m.config.MenuBarMatchesWallpaper = true
		case "Custom":
			m.editingColor = true
			m.colorInput = "#"
			m.message = "Enter a color name or hex value, then press Enter."
			m.focusOptions = false
			return
		}
	case 2:
		switch {
		case strings.HasPrefix(choice, "macOS 15"):
			m.config.CornerStyle = "macos15"
		case strings.HasPrefix(choice, "macOS X"):
			m.config.CornerStyle = "macosx"
		default:
			m.config.CornerStyle = "macos26"
		}
	}
	m.focusOptions = false
}

func (m settingsMenuModel) options() []string {
	switch m.section {
	case 0:
		return displayOptions
	case 1:
		return menuBarOptions
	case 2:
		return cornerOptions
	default:
		return nil
	}
}

func (m settingsMenuModel) View() tea.View {
	const leftWidth, rightWidth, columnMargin = 27, 27, 2
	lines := make([]string, 0, 12)
	leftItems := []string{"Display", "Menubar", "Corner", "", "Save and exit"}
	rightLines := m.rightPaneLines()
	const contentRows = 6 // Keep the box height steady across all display choices.
	lineCount := max(contentRows, len(leftItems))
	for i := 0; i < lineCount; i++ {
		leftText := ""
		if i < len(leftItems) {
			leftText = leftItems[i]
			cursorIndex := i
			if i == 3 {
				cursorIndex = -1
			} else if i == 4 {
				cursorIndex = 3
			}
			if cursorIndex == m.cursor {
				if !m.focusOptions && !m.editingColor {
					leftText = "› " + leftText
				}
				leftText = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#ff7e23")).Render(leftText)
			} else {
				leftText = "  " + leftText
			}
		}
		rightText := ""
		if i < len(rightLines) {
			rightText = rightLines[i]
		}
		leftColumn := lipgloss.NewStyle().Width(leftWidth).Align(lipgloss.Right).Render(leftText + "  ")
		rightColumn := lipgloss.NewStyle().Width(rightWidth).Align(lipgloss.Left).Render(rightText)
		lines = append(lines, strings.Repeat(" ", columnMargin)+leftColumn+" │ "+rightColumn+strings.Repeat(" ", columnMargin))
	}
	lines = append(lines, "")
	footer := "↑/↓ navigate • → select • Q save and exit"
	if m.focusOptions || m.editingColor {
		footer = "↑/↓ navigate • ← back • Enter apply • Ctrl+C cancel"
	}
	if m.editingColor {
		footer = "Type color after # • Enter apply • Esc back"
	}
	box := renderSettingsBox(lines, m.width)
	footerLine := lipgloss.PlaceHorizontal(m.width, lipgloss.Center, footer)
	return tea.NewView(box + footerLine)
}

func (m settingsMenuModel) rightPaneLines() []string {
	if m.editingColor {
		lines := []string{"Custom menu bar color", "", "  " + m.colorInput + "_"}
		if m.message != "" {
			lines = append(lines, "", m.message)
		}
		return lines
	}
	if m.section < 0 {
		return []string{"", "", ""}
	}
	lines := make([]string, 0, len(m.options()))
	for optionIndex, option := range m.options() {
		marker := "  "
		if m.focusOptions && optionIndex == m.option {
			marker = "› "
		}
		optionLine := marker + option
		if m.optionSelected(option) {
			optionLine += " ✓"
		}
		if m.focusOptions && optionIndex == m.option {
			optionLine = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff7e23")).Bold(true).Render(optionLine)
		}
		lines = append(lines, optionLine)
	}
	return lines
}

func renderSettingsBox(lines []string, terminalWidth int) string {
	const innerWidth = 61
	title := " Wallah Settings "
	remaining := innerWidth - lipgloss.Width(title)
	left := remaining / 2
	right := remaining - left
	border := lipgloss.NewStyle().Foreground(lipgloss.Color("#ff7e23"))
	top := "╭" + strings.Repeat("─", left) + border.Render(title) + strings.Repeat("─", right) + "╮"
	boxLines := []string{top, "│" + strings.Repeat(" ", innerWidth) + "│"}
	for _, line := range lines {
		padded := lipgloss.NewStyle().Width(innerWidth).Render(line)
		boxLines = append(boxLines, "│"+padded+"│")
	}
	boxLines = append(boxLines, "╰"+strings.Repeat("─", innerWidth)+"╯")
	box := strings.Join(boxLines, "\n")
	if terminalWidth > lipgloss.Width(box) {
		box = lipgloss.PlaceHorizontal(terminalWidth, lipgloss.Center, box)
	}
	// Leave the terminal cursor below the bottom border so follow-up output
	// (such as the saved-settings notice) cannot overwrite it.
	return box + "\n"
}

func (m settingsMenuModel) selectedOptionIndex() int {
	for i, option := range m.options() {
		if m.optionSelected(option) {
			return i
		}
	}
	return 0
}

func (m settingsMenuModel) optionSelected(option string) bool {
	switch m.section {
	case 0:
		return map[string]string{
			"MacBook Pro 14\"":       "mbp-14",
			"MacBook Pro 16\"":       "mbp-16",
			"MacBook Air 13\"":       "mba-13",
			"MacBook Air 15\"":       "mba-15",
			"iMac 24\" 4.5K":         "imac-24",
			"Studio Display 27\" 5K": "studio-27",
		}[option] == m.config.Display
	case 1:
		if option == "Transparent" {
			return m.config.MenuBarMatchesWallpaper
		}
		if m.config.MenuBarMatchesWallpaper {
			return false
		}
		switch option {
		case "Black":
			return strings.EqualFold(m.config.MenuBarColor, "#000000")
		case "White":
			return strings.EqualFold(m.config.MenuBarColor, "#ffffff")
		case "Custom":
			return !strings.EqualFold(m.config.MenuBarColor, "#000000") && !strings.EqualFold(m.config.MenuBarColor, "#ffffff")
		}
	case 2:
		return strings.HasPrefix(option, "macOS 15") && m.config.CornerStyle == "macos15" || strings.HasPrefix(option, "macOS 26") && m.config.CornerStyle == "macos26" || strings.HasPrefix(option, "macOS X") && m.config.CornerStyle == "macosx"
	}
	return false
}

func (m settingsMenuModel) displayLabel() string {
	labels := map[string]string{
		"mbp-14":    "MacBook Pro 14\"",
		"mbp-16":    "MacBook Pro 16\"",
		"mba-13":    "MacBook Air 13\"",
		"mba-15":    "MacBook Air 15\"",
		"imac-24":   "iMac 24\" 4.5K",
		"studio-27": "Studio Display 27\" 5K",
	}
	if label := labels[m.config.Display]; label != "" {
		return label
	}
	return "Select a display"
}

func (m settingsMenuModel) menuBarLabel() string {
	if m.config.MenuBarMatchesWallpaper {
		return "Transparent (wallpaper color)"
	}
	switch strings.ToLower(m.config.MenuBarColor) {
	case "#000000":
		return "Black"
	case "#ffffff":
		return "White"
	default:
		return m.config.MenuBarColor
	}
}

func (m settingsMenuModel) cornerLabel() string {
	if m.config.CornerStyle == "macosx" {
		return "macOS X (squared)"
	}
	if m.config.CornerStyle == "macos15" {
		return "macOS 15 (less round)"
	}
	return "macOS 26 (round)"
}

func normalizeMenuColor(value string) (string, error) {
	value = strings.ToLower(strings.TrimSpace(value))
	if color, ok := wallpapers.Color(value); ok {
		return color.Background, nil
	}
	hex := strings.TrimPrefix(value, "#")
	if len(hex) == 3 {
		hex = string([]byte{hex[0], hex[0], hex[1], hex[1], hex[2], hex[2]})
	}
	if len(hex) != 6 {
		return "", fmt.Errorf("use a known color name or a 3 or 6 digit hex color")
	}
	if _, err := strconv.ParseUint(hex, 16, 24); err != nil {
		return "", fmt.Errorf("invalid hex color %q", value)
	}
	return "#" + hex, nil
}
