// This file contains two lists:
//
// 1. Ordered list of the colors
// 1. The colors available and their mapped styles (fg and bg).
// 2. The colors available the location for each of them.

package cli

// a slice of the colors in order,
// instead of output the map itself (which is random)
var wallpaperOrder = []string{
	"black", "black-b",
	"teddy", "teddy-b",
	"yellow", "yellow-b",
	"pikachu", "pikachu-b",
	"cyan", "cyan-b",
	"blender", "blender-b",
	"blue", "blue-b",
	"stitchy", "stitchy-b",
	"white", "white-b",
	"kirby", "kirby-b",
	"red", "red-b",
	"akira", "akira-b",
	"green", "green-b",
	"yoshi", "yoshi-b",
	"magenta", "magenta-b",
	"grimace", "grimace-b",
}

var colorStyles = map[string][2]string{

	// lowercase for hex recommended for compatibility
	// 0 is black, f is white
	// ansi 16
	"black":     {"#ffffff", "#000000"},
	"blue":      {"#ffffff", "#0000ff"},
	"cyan":      {"#4f4f4f", "#00cdcd"},
	"green":     {"#ffffff", "#00dd00"},
	"magenta":   {"#ffffff", "#cd00cd"},
	"red":       {"#ffffff", "#cd0000"},
	"yellow":    {"#4f4f4f", "#cdcd00"},
	"white":     {"#4f4f4f", "#e5e5e5"},
	"black-b":   {"#ffffff", "#7f7f7f"},
	"blue-b":    {"#ffffff", "#5c5cff"},
	"cyan-b":    {"#4f4f4f", "#00ffff"},
	"green-b":   {"#ffffff", "#00ff00"},
	"magenta-b": {"#ffffff", "#ff00ff"},
	"red-b":     {"#ffffff", "#ff0000"},
	"yellow-b":  {"#4f4f4f", "#ffff00"},
	"white-b":   {"#4f4f4f", "#ffffff"},

	// bonus
	"orange":   {"#fff", "#ff5516"},
	"orange-b": {"#fff", "#ff7e23"},

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
	"blender-b": {"#fff", "#c2d6e2"},
	"grimace-b": {"#fff", "#8e1e8c"},
	"kirby-b":   {"#fff", "#f49792"},
	"pikachu-b": {"#fff", "#f4db06"},
	"stitchy-b": {"#fff", "#69b8dc"},
	"teddy-b":   {"#fff", "#d09249"},
	"yoshi-b":   {"#fff", "#69c12b"},
}
