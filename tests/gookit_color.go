package main

import (
	"fmt"

	"github.com/gookit/color"
)

func main() {

	// Named color with inline tags
	red := "This is <fg=white;bg=red> `red` </> explicitly."
	fmt.Println(color.Render(red))

	// 256-color with inline tags
	two := "This is <fg=f;bg=46> green (46) </> in 256 colors."
	fmt.Println(color.Render(two))

	// Hex colors with inline tags (truecolor)
	hex := "This is <fg=#0000aa;bg=#00ff00> blue hex </> ."
	s := "This is <fg=#000000;bg=#00ff00> green hex </> ."
	fmt.Println(color.Render(hex))
	fmt.Println(color.Render(s))

	// Hex colors with API calls (alternative to inline tags)
	c := color.New(color.FgHex("#0000aa"), color.BgHex("#00ff00"))
	c.Println("blue hex with API")

	// Fallback to 256-color codes if inline hex tags don't work reliably
	hexFallback := "This is <fg=4;bg=46> blue hex fallback </>."
	fmt.Println(color.Render(hexFallback))

	// test this on the terminal
	// printf "\x1b[38;2;255;100;0mTRUECOLOR\x1b[0m\n"
}
