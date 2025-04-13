<a><img src="https://github.com/user-attachments/assets/1f233586-a2c3-4838-beeb-3fcbd793ba82" height="150" width="150"></a>

![Tests](https://github.com/Greeshmanth1909/Shadowfax/actions/workflows/ci.yml/badge.svg)
# Shadowfax
Shadowfax is a UCI chess engine written entirely from scratch in Golang. It will be using classic evlauation and move generation methods. I also plan on upgrading it to an NNUE in version 2. If that sounds interesting, consider giving this repo a star. Thanks!

You can play a game against shadowfax on [Lichess](https://lichess.org/@/UCI_Shadowfax)

If you're interested in contributing to the project, please go through [contributing.md](CONTRIBUTING.md)

# Build from source

## Requirements
- `go 1.22.2` or later

## Steps
1. Clone this repo with `git clone https://github.com/Greeshmanth1909/shadowfax`
2. Move to the root of the project with  `cd shadowfax`
3. Build the binary with `go build`
4. The shadowfax binary should appear in the root of the project
5. Link the binary to a chess gui

Here are the GUI's that have been tested so far
- [en-croissant](https://github.com/franciscoBSalgueiro/en-croissant) for mac
- [Arena Chess](http://www.playwitharena.de/) for windows

For more instructions on how to link the engine to a gui head [here](addGui.md)

# Credits
- bluefeversoftware for his VICE Engine
- The Chess Programming Wiki
