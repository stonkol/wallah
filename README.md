# rmv-mbp-notch

Wallpaper maker for removing the top notch and rounding the lower corners.

A none-resource-consuming, app-free, zero-dependency alternative to app [Top Notch](https://topnotch.app/). Inspired by [removethenotch](removethenotch.com)

You also can adjust your own colors or put your own images on the

## wall-cli

A CLI app for changing the color of your wallpaper.

### The colors

Example of one of the wallpapers (bright blue):

<img src="wallpapers/mbp-14/b-blue.png" max-width="350px">

#### ANSI 16

The basic 8:
`black`, `red`, `green`, `yellow`, `blue`, `magenta`, `cyan`, `white`.

And their bright versions:
`b-black`, `b-red`, `b-green`, `b-yellow`, `b-blue`, `b-magenta`, `b-cyan`, `b-white`.

<img src="wallpapers/mbp-14/ansi.png" max-width="400px">

2. Colors of 8 Characters: `pikachu`, `charmander`, `stitchy`, `yoshi`, `blender`, `kirby`, `teddy`, `wario`.

<img src="wallpapers/mbp-14/pokemon.png" max-width="250px">

### Usage

```sh
wall blue      # set blue (#0000EE)
wall b-blue   # set bright blue (#5C5CFF)
```

1. macOS requires explicit permission for apps (including Terminal or the compiled CLI binary) to change the wallpaper

1. macOS may ask to accept permissions the first time your run the application.


### Having problems

#### Only changed on one deskptop?

On newer macOS versions (Ventura, Sonoma), wallpaper management changed:
The wallpaper might only change on the current desktop/space.
You might need to set wallpaper for each desktop separately.

#### Running the AppleScript manually

Test the AppleScript works with your wallpaper manually in the Terminal:
```sh
osascript -e 'tell application "System Events" to set picture of every desktop to POSIX file "/absolute/path/to/image.jpg"'
```

#### Run the command as the logged-in user
If you run your Go program as root or via sudo, the AppleScript may run as root and fail to change the wallpaper for the logged-in user.
Make sure you run the CLI as the current logged-in user.

### How it works

- Using `osascript` command to change the color of the wallpaper on macOS.
- It runs AppleScript from Go using `os/exec` package.

### Build and Run

Build it (in your project directory run):
```sh
go build -o wall # compiles the Go code into an executable named wall
```

Run the cli and input a `color`:
```sh
./wall blue     # set blue (#0000EE)
./wall b-blue   # set bright blue (#5C5CFF)
```

If you get a permission error on macOS or Linux, make the file executable:
```sh
chmod +x ./wall
```

To run the app globally, you can move the built executable to a directory in your PATH, like `/usr/local/bin`.

## TODO

- Make a script to rmv the notch of a wallpaper.
    - [x] improve flags
    - [ ] Have colored text on the CLI

- [ ] cli tool to update a unnotch wallpaper with the color you want

- [ ] dark mode script
    - a script to change the color of the wall according to system settings.
    - let the user setting if

- Add other versions
    - [ ] Include the 16 [ANSI colors](https://en.wikipedia.org/wiki/List_of_software_palettes)
    - [ ] 16" MBP version.
    - [ ] 13" and 15" MacBook Air
    - [ ] Colors of 8 Characters

## Acknowledgements

The Pikachu GIF used above, provided by the [PokeAPI](https://github.com/PokeAPI/sprites) sprites repository are the property of The Pokémon Company. All rights to the Pokémon characters and related artwork are owned by The Pokémon Company.
