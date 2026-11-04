# Testing Notes

## Try it while dev

```sh
go run . -l      # to list colors
go run . blue    # set a blue wallpaper
```

### Preview wallahs

It generates four sample PNGs with distinct colors for the MacBook Pro 14″ and 16″ and MacBook Air 13″ and 15″.
```sh
wallah preview
```

## How to update 

### Update from local source
To try the changes you have in your working copy, run this from the project directory:

```sh
go install .

# Then start the installed app
wallah list
wallah version
```

### Update from Github 
That installs your local code to `$(go env GOPATH)/bin`. Make sure that directory is on your **PATH**.

To install from GitHub instead, the changes need to be committed and pushed first. Then run:

```sh
go install github.com/stonkol/wallah@main
```

Use `@latest` if you want the latest tagged release instead of the latest commit on **main**.

## Autocompletion

In your current Zsh session, run:

```sh
mkdir -p ~/.zsh/completions
wallah completion zsh > ~/.zsh/completions/_wallah
fpath=(~/.zsh/completions $fpath)
autoload -Uz compinit && compinit
```

Now type `wallah blu` and press **Tab**. It should complete `blue` or offer `blue-b`.

To keep completion working in new terminal sessions, add these lines to `~/.zshrc`:

```sh
fpath=(~/.zsh/completions $fpath)
autoload -Uz compinit
compinit
```

Then source it: `source ~/.zshrc`

## Roadmap

### General

- [x] --help flag
- [x] --list flag
- [x] have a working binary working alone
- [x] Include the 16 [ANSI colors](https://en.wikipedia.org/wiki/List_of_software_palettes)
- [x] Make a logo to put on the top of the README on GitHub
- [x] set random color for the wallah `-r` `--random`
- [x] change to previous color wallah `-p` `--previous`
- [x] set a wallpaper to all the desktops `-a` `--all`
- [x] autocompletion

#### Init select prompt

- [ ] init prompt
    - [ ] The first time the users runs wallah (when the user type `wallah`) will prompt them to select a device.
        1. MacBook Pro 14"
        2. Macbook Pro 16"
        3. MacBook Air 13"
        4. MacBook Air 15"
    - [ ] Then choose if the wallpaper will have a [black/white/choose a color] top menu bar.
    - [ ] then chosse the corner border 
        - [ ] macOS 26 (rounder)
        - [ ] macOS 15
    - Previously all the wallpapers where already made by a graphic editor, but now i want this Go app to make the wallpaper image in real time when the user input the color, so it will clip mask the wallpaper to a be according to the user selected MacBook. So it will have a menu bar part in according to waht the user have selected and the other part of the wallpaper as what the user have selected. And then make corner borders like in macOS 26 which is more rounder or before (less rounder).
    - after user have selected these three parameters, tell the user that now can select colors from the list above, showing the printList. these three settings above will be saved, and the user will only need to type a color next time it uses wallah, but if uses in the future type wallah -s or wallah --settings it will appear again the init prompt

- [ ] update printwallah's wallahMenuBar for macos26 version which it will be a notch [color selected][notch][●][notch][color selected] instead of the [black][●][black]


### Styling the cli

- [x] Colors of 8 characters
- [x] Have colored text on the CLI
- [x] have multiple columns to show the list
- [x] have a styled list
- [ ] have a gif of using wallah using a tape of [vhs](https://github.com/charmbracelet/vhs) in readme.

### Light / Dark mode

- [x] Let users set the wallpaper for dark mode and light mode, which will change according to the system settings.
- [ ] Include a `-d` `--dark` and `-l` `--light` flags for toggling mode
- [ ] Let user enable/disable auto-light/dark mode.
- [ ] Change toggle dark mode of the wall (between [color] and [color]-b) according to system settings.

### Future Features

#### Image

- [ ] Let users change the wallpaper to an image
    - [ ] apply de-notch to the wallpaper

#### Dynamic Wallpaper Creator
- [ ] based on users input color

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
- [ ] 14" and 16" MacBook Pro
- [ ] 13" and 15" MacBook Air
- [ ] no-notch devices
- [ ] Linux support

### More
- [ ] if input `wallah mama`, easter egg appears
