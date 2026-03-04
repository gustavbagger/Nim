# Nim
This is a CLI game based on the mathematical combinatorial game **Nim**. The game can be played against the computer or another human locally. The two players take turns removing (or "nimming") objects from distinct heaps or piles. On each turn, a player must remove at least one object, and may remove any number of objects provided they all come from the same heap or pile. The goal of the game is to take the last object.
## Quick Start
**Running on Windows**
1) Download NimWindows.exe
2) Run NimWindows.exe as a normal programme

**install and run using Go toolchain**
```
go install github.com/gustavbagger/Nim
Nim
```
## Usage
The game should prompt you through the setup process. 
- *Who is playing:* enter one or two names depending on if you want to play against the computer or against another human
- *Enter desired game setup:* enter the initial heap sizes, the string 'd' will give you the standard setup of "7 5 3 1" corresponding to four piles with a total of 16 tiles.
- *Who starts?* enter a player name from the list of players. The computer is (unsurprisingly) signified by entering the string "computer"
- *Player's turn:* On each turn, enter the heap you want to remove from and the amount of tiles you want to remove. Example: "1 2" would remove 2 tiles from row 1.
- At the end of the game you will be prompted to keep playing or exit. Enter either "y" for yes or "n" for no.
## Contributing

### Clone the repo

```bash
git clone https://github.com/gustavbagger/Nim
cd Nim
```

### Build the compiled binary

```bash
go build
```
or 
```bash
GOOS=windows GOARCH=amd64 go build -o <fileName.exe>
```
