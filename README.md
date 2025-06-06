<div align="center">

<img src="assets/wallah-logo.png" width="50%">

</br>

![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/stonkol/wallah?sort=semver)
![Go Version](https://img.shields.io/github/go-mod/go-version/stonkol/wall-cli)
![Build Status](https://github.com/stonkol/wallah/actions/workflows/go.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/stonkol/wall-cli)](https://goreportcard.com/report/github.com/stonkol/wall-cli)

</div>

wallah-cli is a CLI which you can change your macOS wallpaper, in which the menu bar part is black, in order to hide that ugly notch. Furthermore, it is less distracting if you have the menu bar auto-hide enabled.

It is a minimal alternative to the app [Top Notch](https://topnotch.app/), and it is inspired by [removethenotch](removethenotch.com).

### Installation

You need to have Go installed. This command will compile and install the CLI tool in your system PATH.
```sh
go install github.com/stonkol/wallah@latest
```

If it is installed correctly, this will set your wallpaper to blue:
```sh
wallah blue
```

<div align="center">
</br>
<p>Colors available</p>
<img src="assets/color-list.jpg" width="66%">

</div>

## Index

1. [Build yourself](#1-build-yourself)
1. [Colors included](#2-colors-included)
1. [How it works](#4-how-it-works)
1. [Roadmap](#5-roadmap)
1. [Having issues?](#5-having-issues)
1. [Contributing](#6-contributing)

---

## 1. Build yourself

### 1.1 Build the Binary

Compile all the go files:
```sh
go build
```

> [!caution] Caution
> Don't use `wall` or other name that is used for a standard Unix/Linux command.

### 1.2 Move the Binary

Move the binary to a directory in your PATH, to run the app globally:
```sh
sudo mv wallah /usr/local/bin/
sudo chmod +x /usr/local/bin/wallah # make the file executable
```

1. macOS may require explicit permission for apps (including Terminal or the compiled CLI binary) to change the wallpaper.
1. macOS may ask to accept permissions the first time your run the application.

### 1.3 Try it out

```sh
wallah -h        # show help
wallah -l        # show list of colors available
wallah -r        # set a random color
wallah -v        # show version
wallah blue      # set blue (#0000EE)
wallah blue-b    # set bright blue (#5C5CFF)
wallah pikachu   # set pikachu color (#F8A21C)
```

### 1.4 Visuals

Set a wallpaper: `wallah akira`

<div align="center">
<img src="assets/clit-visual.jpg" width="50%">
</div>

Show list: `wallah -l`

## 2. Colors included

Example of one of the wallpapers (bright blue):

<div align="center">
<img src="wallpapers/mbp-14/blue-b.png" width="50%">
</div>

### 2.1 ANSI 16

The basic 8 colors, by default, is the dark mode version. Add `b-` prefix for bright (light) versions (`b-blue`).

`black`, `blue`, `cyan`, `green`, `magenta`, `red`, `yellow` and `white`.

<div align="center">
<img src="assets/ansi-16.png" width="50%">
</div>

### 2.2 CHAR 16

Based on 8 famous characters, each with their own light and dark mode. Add `b-` prefix for bright (light) versions (`b-pikachu`).

`akira`, `blender`, `grimace`, `kirby`, `pikachu`, `stitchy`, `teddy` and `yoshi`.

<div align="center">
<img src="assets/char-16.png" width="50%">
</div>

## 3. Roadmap

### General

- [x] --help flag
- [x] --list flag
- [x] have a working binary working alone
- [x] Include the 16 [ANSI colors](https://en.wikipedia.org/wiki/List_of_software_palettes)
- [x] Make a logo to put on the top of the README on GitHub
- [x] set random color for the wallah `-r` `--random`
- [ ] change to previous color wallah `-p` `--previous`
- [ ] set a wallpaper to all the desktops `-a` `--all`
- [ ] have a gif of using wallah using a tape of [vhs](https://github.com/charmbracelet/vhs) in readme.
- [ ] Let the user be able to have settings
- [ ] autocompletion
- [ ] if input `wallah mama`, easter egg appears

### Styling the cli

- [x] Colors of 8 characters
- [x] Have colored text on the CLI
- [x] have multiple columns to show the list
- [x] have a styled list

### Light / Dark mode

- [x] Let users set the wallpaper for dark mode and light mode, which will change according to the system settings.
- [ ] Include a `-d` `--dark` and `-l` `--light` flags for toggling mode
- [ ] Let user enable/disable auto-light/dark mode.
- [ ] Change toggle dark mode of the wall (between [color] and [color]-b) according to system settings.

### Future Features

#### Image

- [ ] Let users change the wallpaper to an image
    - [ ] apply de-notch to the wallpaper

#### Init select prompt

The first time the users runs wallah will prompt them to select a device. Then choose if the wallpaper will have a [black/white/choose a color] top menu bar or without.

#### Dynamic Wallpaper Creator
based on users input color

##### Colors

- [ ] ANSI 16/256 support
- [ ] True Color (HEX) support
- [ ] Adaptive Colors (specify color options for light and dark mode)
- [ ] Apply color themes - [ghostty](/Applications/Ghostty.app/Contents/Resources/ghostty/themes/GitLab-Dark) have a big collection
    - e.g: Catppuccin for the ANSI 16 colors
    - `wallah red -t catppuccin` it will apply the red of Catppuccin
    - also can set the default theme: `wallah --set-theme catppuccin`
- [ ] Gradient support (user input 2 to 3 colors and the direction)
- [ ] Adaptive Gradient support

> [!info] Adaptive Gradient
> same as gradient but will change depending on the sun, like the native wallpaper on macOS

##### Device support
- [ ] no-notch devices
- [ ] 14" and 16" MacBook Pro
- [ ] 13" and 15" MacBook Air

- [ ] Linux support

#### Animations

- [ ] When changing the wallpaper

## 4. How it works

- Using `osascript` command to change the color of the wallpaper on macOS.
- It runs AppleScript from Go using the `os/exec` package.
- All the default wallpapers are located in the `wallpapers` directory.
- There is a `design-files` directory with the design files (`.psd` and `.afdesign`) used to create the wallpapers. Which you can use to create your own wallpapers, either with solid colors, gradients, or an image.

## 5. Having issues

#### 5.1 Only changed on one desktop?

On newer macOS versions (Ventura, Sonoma), wallpaper management changed:
The wallpaper might only change on the current desktop/space.
You might need to set wallpaper for each desktop separately.

#### 5.2 Running the AppleScript manually

Test the AppleScript works with your wallpaper manually in the Terminal:

```sh
osascript -e 'tell application "System Events" to set picture of every desktop to POSIX file "/absolute/path/to/image.jpg"'
```

#### 5.3 Run the command as the logged-in user

If you run your Go program as root or via sudo, the AppleScript may run as root and fail to change the wallpaper for the logged-in user.
Make sure you run the CLI as the current logged-in user.


## 6. Contributing

Welcome to make a pull request or add an issue, but before, please read the README.md before contributing. 😀
