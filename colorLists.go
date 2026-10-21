// This file contains two lists:
//
// 1. Ordered list of the colors
// 1. The colors available and their mapped styles (fg and bg).
// 2. The colors available the location for each of them.

package main

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

var Wallpapers = map[string]string{
	// ansi 16
	"black":   "wallpapers/mbp-14/black.png",
	"blue":    "wallpapers/mbp-14/blue.png",
	"cyan":    "wallpapers/mbp-14/cyan.png",
	"green":   "wallpapers/mbp-14/green.png",
	"magenta": "wallpapers/mbp-14/magenta.png",
	"red":     "wallpapers/mbp-14/red.png",
	"white":   "wallpapers/mbp-14/white.png",
	"yellow":  "wallpapers/mbp-14/yellow.png",

	"black-b":   "wallpapers/mbp-14/black-b.png",
	"blue-b":    "wallpapers/mbp-14/blue-b.png",
	"cyan-b":    "wallpapers/mbp-14/cyan-b.png",
	"green-b":   "wallpapers/mbp-14/green-b.png",
	"magenta-b": "wallpapers/mbp-14/magenta-b.png",
	"red-b":     "wallpapers/mbp-14/red-b.png",
	"yellow-b":  "wallpapers/mbp-14/yellow-b.png",
	"white-b":   "wallpapers/mbp-14/white-b.png",

	// BONUS
	"orange":   "wallpapers/mbp-14/orange.png",
	"orange-b": "wallpapers/mbp-14/orange-b.png",

	// char 16
	"akira":   "wallpapers/mbp-14/akira.png",
	"blender": "wallpapers/mbp-14/blender.png",
	"kirby":   "wallpapers/mbp-14/kirby.png",
	"pikachu": "wallpapers/mbp-14/pikachu.png",
	"stitchy": "wallpapers/mbp-14/stitchy.png",
	"teddy":   "wallpapers/mbp-14/teddy.png",
	"grimace": "wallpapers/mbp-14/grimace.png",
	"yoshi":   "wallpapers/mbp-14/yoshi.png",

	"akira-b":   "wallpapers/mbp-14/akira-b.png",
	"blender-b": "wallpapers/mbp-14/blender-b.png",
	"kirby-b":   "wallpapers/mbp-14/kirby-b.png",
	"pikachu-b": "wallpapers/mbp-14/pikachu-b.png",
	"stitchy-b": "wallpapers/mbp-14/stitchy-b.png",
	"teddy-b":   "wallpapers/mbp-14/teddy-b.png",
	"grimace-b": "wallpapers/mbp-14/grimace-b.png",
	"yoshi-b":   "wallpapers/mbp-14/yoshi-b.png",
}
