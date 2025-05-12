// This file contains two lists:
//
// 1. The colors available and their mapped styles (fg and bg).
// 2. The colors available the location for each of them.

package main

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
	"white":     {"#000", "#e5e5e5"},
	"black-b":   {"#fff", "#7f7f7f"},
	"blue-b":    {"#fff", "#5c5cff"},
	"cyan-b":    {"#000", "#00ffff"},
	"green-b":   {"#000", "#00ff00"},
	"magenta-b": {"#fff", "#ff00ff"},
	"red-b":     {"#fff", "#ff0000"},
	"yellow-b":  {"#000", "#ffff00"},
	"white-b":   {"#000", "#ffffff"},

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
