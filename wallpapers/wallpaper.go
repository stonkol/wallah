package wallpapers

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type Palette struct {
	Foreground string
	Background string
}

var palette = map[string]Palette{
	"black": {"#ffffff", "#000000"}, "blue": {"#ffffff", "#0000ff"},
	"cyan": {"#4f4f4f", "#00cdcd"}, "green": {"#ffffff", "#00dd00"},
	"magenta": {"#ffffff", "#cd00cd"}, "red": {"#ffffff", "#cd0000"},
	"yellow": {"#4f4f4f", "#cdcd00"}, "white": {"#4f4f4f", "#e5e5e5"},
	"black-l": {"#ffffff", "#7f7f7f"}, "blue-l": {"#ffffff", "#5c5cff"},
	"cyan-l": {"#4f4f4f", "#00ffff"}, "green-l": {"#ffffff", "#00ff00"},
	"magenta-l": {"#ffffff", "#ff00ff"}, "red-l": {"#ffffff", "#ff0000"},
	"yellow-l": {"#4f4f4f", "#ffff00"}, "white-l": {"#4f4f4f", "#ffffff"},
	"orange": {"#fff", "#ff5516"}, "orange-l": {"#fff", "#ff7e23"},
	"akira": {"#fff", "#ab2f1a"}, "blender": {"#fff", "#a2bcd0"},
	"grimace": {"#fff", "#6a1c63"}, "kirby": {"#fff", "#fc5c8e"},
	"pikachu": {"#fff", "#f8a21c"}, "stitchy": {"#fff", "#527bac"},
	"teddy": {"#fff", "#925923"}, "yoshi": {"#fff", "#43a934"},
	"akira-l": {"#fff", "#df2626"}, "blender-l": {"#fff", "#c2d6e2"},
	"grimace-l": {"#fff", "#8e1e8c"}, "kirby-l": {"#fff", "#f49792"},
	"pikachu-l": {"#fff", "#f4db06"}, "stitchy-l": {"#fff", "#69b8dc"},
	"teddy-l": {"#fff", "#d09249"}, "yoshi-l": {"#fff", "#69c12b"},
}

type DisplaySettings struct {
	Display                 string
	MenuBarColor            string
	CornerStyle             string
	MenuBarMatchesWallpaper bool
}

// Names returns the available wallpaper colors in sorted order.
func Names() []string {
	names := make([]string, 0, len(palette))
	for name := range palette {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

// Color returns the foreground and wallpaper colors for a named color.
func Color(name string) (Palette, bool) {
	c, ok := palette[name]
	if ok {
		return c, true
	}
	parsed, err := parseHexColor(name)
	if err != nil {
		return Palette{}, false
	}
	background := fmt.Sprintf("#%02x%02x%02x", parsed.R, parsed.G, parsed.B)
	luminance := 0.299*float64(parsed.R) + 0.587*float64(parsed.G) + 0.114*float64(parsed.B)
	foreground := "#ffffff"
	if luminance > 150 {
		foreground = "#000000"
	}
	return Palette{Foreground: foreground, Background: background}, true
}

// Change renders a wallpaper for the selected display and applies it in macOS.
func Change(colorName string, display DisplaySettings, allDesktops bool) error {
	colors, ok := Color(colorName)
	if !ok {
		return fmt.Errorf("invalid color %q: use a named color or a 3 or 6 digit hex value", colorName)
	}
	if _, named := palette[colorName]; !named {
		colorName = colors.Background
	}
	wallpaperColor, err := parseHexColor(colors.Background)
	if err != nil {
		return fmt.Errorf("invalid wallpaper color for %q: %w", colorName, err)
	}
	menuColor, err := parseHexColor(display.MenuBarColor)
	if err != nil {
		return fmt.Errorf("invalid menu bar color: %w", err)
	}
	img, err := render(wallpaperColor, menuColor, display)
	if err != nil {
		return err
	}

	wallpaperPath, err := saveGeneratedImage(colorName, display, img)
	if err != nil {
		return err
	}

	desktopTarget := "current desktop"
	if allDesktops {
		desktopTarget = "every desktop"
	}
	appleScriptPath := strings.ReplaceAll(wallpaperPath, `\`, `\\`)
	appleScriptPath = strings.ReplaceAll(appleScriptPath, `"`, `\"`)
	script := fmt.Sprintf(`tell application "System Events" to set picture of %s to POSIX file "%s"`, desktopTarget, appleScriptPath)
	if err := exec.Command("osascript", "-e", script).Run(); err != nil {
		return fmt.Errorf("execute AppleScript: %w", err)
	}
	return nil
}

// GeneratePreviews renders one wallpaper for each supported display without
// applying any of them to the desktop.
func GeneratePreviews() ([]string, error) {
	previews := []struct {
		display string
		color   string
	}{
		{display: "mba-15", color: "orange"},
		{display: "mba-13", color: "green"},
		{display: "mbp-14", color: "red"},
		{display: "mbp-16", color: "blue"},
		{display: "imac-24", color: "yellow"},
		{display: "studio-27", color: "magenta"},
	}

	cornerStyles := []string{"macos26", "macos15", "macosx"}
	paths := make([]string, 0, len(previews)*len(cornerStyles)*2)
	for _, preview := range previews {
		colors := palette[preview.color]
		base, err := parseHexColor(colors.Background)
		if err != nil {
			return nil, fmt.Errorf("parse preview color %q: %w", preview.color, err)
		}
		menu, err := parseHexColor("#000000")
		if err != nil {
			return nil, fmt.Errorf("parse preview menu bar color: %w", err)
		}
		for _, cornerStyle := range cornerStyles {
			for _, menuBarMatchesWallpaper := range []bool{false, true} {
				display := DisplaySettings{
					Display:                 preview.display,
					MenuBarColor:            "#000000",
					CornerStyle:             cornerStyle,
					MenuBarMatchesWallpaper: menuBarMatchesWallpaper,
				}
				img, err := render(base, menu, display)
				if err != nil {
					return nil, fmt.Errorf("render %s %s preview: %w", preview.display, cornerStyle, err)
				}
				path, err := savePreviewImage(preview.display, cornerStyle, menuBarMatchesWallpaper, img)
				if err != nil {
					return nil, fmt.Errorf("save %s %s preview: %w", preview.display, cornerStyle, err)
				}
				paths = append(paths, path)
			}
		}
	}
	return paths, nil
}

func saveGeneratedImage(colorName string, display DisplaySettings, img image.Image) (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("find user config directory: %w", err)
	}
	wallpaperDir := filepath.Join(configDir, "wallah", "wallpapers")
	return saveImage(wallpaperDir, colorName, display, img)
}

func savePreviewImage(display, cornerStyle string, menuBarMatchesWallpaper bool, img image.Image) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("find home directory for wallpaper previews: %w", err)
	}
	previewDir := filepath.Join(homeDir, "Downloads", "previews")
	if err := os.MkdirAll(previewDir, 0o700); err != nil {
		return "", fmt.Errorf("create wallpaper preview directory: %w", err)
	}
	filename := cornerStyle + "-" + display + ".png"
	if menuBarMatchesWallpaper {
		filename = cornerStyle + "-transparent-" + display + ".png"
	}
	path := filepath.Join(previewDir, filename)
	if err := savePNG(path, previewDir, img); err != nil {
		return "", err
	}
	return path, nil
}

func saveImage(wallpaperDir, colorName string, display DisplaySettings, img image.Image) (string, error) {
	if err := os.MkdirAll(wallpaperDir, 0o700); err != nil {
		return "", fmt.Errorf("create wallpaper directory: %w", err)
	}
	radius := "r0"
	if display.CornerStyle == "macos15" {
		radius = "r20"
	} else if display.CornerStyle == "macos26" {
		radius = "r52"
	}
	filename := fmt.Sprintf("%s-%s-%s-%s-%s.png", colorName, display.Display, strings.TrimPrefix(display.MenuBarColor, "#"), display.CornerStyle, radius)
	if display.MenuBarMatchesWallpaper {
		filename = strings.TrimSuffix(filename, ".png") + "-transparent.png"
	}
	path := filepath.Join(wallpaperDir, filename)
	if err := savePNG(path, wallpaperDir, img); err != nil {
		return "", err
	}
	return path, nil
}

func savePNG(path, directory string, img image.Image) error {
	tmpFile, err := os.CreateTemp(directory, ".wallpaper-*.png")
	if err != nil {
		return fmt.Errorf("create temporary wallpaper: %w", err)
	}
	tmpPath := tmpFile.Name()
	defer os.Remove(tmpPath)
	if err := png.Encode(tmpFile, img); err != nil {
		tmpFile.Close()
		return fmt.Errorf("encode wallpaper image: %w", err)
	}
	if err := tmpFile.Close(); err != nil {
		return fmt.Errorf("close temporary wallpaper: %w", err)
	}
	if err := os.Rename(tmpPath, path); err != nil {
		return fmt.Errorf("save generated wallpaper: %w", err)
	}
	return nil
}

func render(base, menu color.RGBA, display DisplaySettings) (*image.RGBA, error) {
	width, height, err := dimensions(display.Display)
	if err != nil {
		return nil, err
	}
	if display.CornerStyle != "macos26" && display.CornerStyle != "macos15" && display.CornerStyle != "macosx" {
		return nil, fmt.Errorf("unknown corner style %q", display.CornerStyle)
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	if display.MenuBarMatchesWallpaper {
		menu = base
	}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.SetRGBA(x, y, menu)
		}
	}

	// Round the wallpaper area below the menu bar. The menu color remains
	// visible outside its top corners, matching the background behind the notch.
	left, right := 0, width
	radius := 0.0
	if display.CornerStyle == "macos15" {
		radius = 20.0
	} else if display.CornerStyle == "macos26" {
		radius = 52.0
	}
	// On a 14-inch MacBook Pro, the 66 px menu bar leaves 1898 px below it.
	menuHeight := max(1, (height*66+982)/1964)
	top, bottom := menuHeight, height
	for y := top; y < bottom; y++ {
		for x := left; x < right; x++ {
			coverage := 0.0
			if inCornerRegion(x, y, left, top, right, bottom, radius) {
				coverage = cornerCoverage(x, y, left, top, right, bottom, radius)
			} else {
				coverage = 1.0
			}
			if coverage <= 0 {
				continue
			}
			img.SetRGBA(x, y, blend(base, menu, coverage))
		}
	}
	return img, nil
}

func blend(foreground, background color.RGBA, coverage float64) color.RGBA {
	return color.RGBA{
		R: uint8(float64(foreground.R)*coverage + float64(background.R)*(1-coverage)),
		G: uint8(float64(foreground.G)*coverage + float64(background.G)*(1-coverage)),
		B: uint8(float64(foreground.B)*coverage + float64(background.B)*(1-coverage)),
		A: uint8(float64(foreground.A)*coverage + float64(background.A)*(1-coverage)),
	}
}

func dimensions(display string) (int, int, error) {
	switch display {
	case "mbp-14":
		return 3024, 1964, nil
	case "mbp-16":
		return 3456, 2234, nil
	case "mba-13":
		return 2560, 1664, nil
	case "mba-15":
		return 2880, 1864, nil
	case "imac-24":
		return 4480, 2520, nil
	case "studio-27":
		return 5120, 2880, nil
	default:
		return 0, 0, fmt.Errorf("unknown display %q", display)
	}
}

func inCornerRegion(x, y, left, top, right, bottom int, radius float64) bool {
	margin := int(math.Ceil(radius))
	return x < left+margin && y < top+margin ||
		x >= right-margin && y < top+margin ||
		x < left+margin && y >= bottom-margin ||
		x >= right-margin && y >= bottom-margin
}

func cornerCoverage(x, y, left, top, right, bottom int, radius float64) float64 {
	const samples = 4
	inside := 0
	for sy := 0; sy < samples; sy++ {
		for sx := 0; sx < samples; sx++ {
			px := float64(x) + (float64(sx)+0.5)/samples
			py := float64(y) + (float64(sy)+0.5)/samples
			if insideRoundedRect(px, py, float64(left), float64(top), float64(right), float64(bottom), radius) {
				inside++
			}
		}
	}
	return float64(inside) / (samples * samples)
}

func insideRoundedRect(x, y, left, top, right, bottom, radius float64) bool {
	if x >= left+radius && x < right-radius || y >= top+radius && y < bottom-radius {
		return true
	}
	cx, cy := left+radius, top+radius
	if x >= right-radius {
		cx = right - radius
	}
	if y >= bottom-radius {
		cy = bottom - radius
	}
	dx, dy := x-cx, y-cy
	return dx*dx+dy*dy <= radius*radius
}

func parseHexColor(value string) (color.RGBA, error) {
	if len(value) > 0 && value[0] == '#' {
		value = value[1:]
	}
	if len(value) == 3 {
		value = string([]byte{value[0], value[0], value[1], value[1], value[2], value[2]})
	}
	if len(value) != 6 {
		return color.RGBA{}, fmt.Errorf("expected a 3 or 6 digit hex color")
	}
	parsed, err := strconv.ParseUint(value, 16, 24)
	if err != nil {
		return color.RGBA{}, fmt.Errorf("invalid hex color %q", value)
	}
	return color.RGBA{R: uint8(parsed >> 16), G: uint8(parsed >> 8), B: uint8(parsed), A: 255}, nil
}
