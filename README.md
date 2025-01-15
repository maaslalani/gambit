<img src="./assets/gambit.png?raw=true" alt="Gambit: Play chess in your terminal" />

# Gambit

Chess board in your terminal.

### Installation

```sh
go install github.com/maaslalani/gambit@latest
```

or run from source

```sh
git clone https://github.com/maaslalani/gambit
cd gambit/
go run ./...
```

#### Docker

Gambit is available as a Docker image at [ghcr.io/maaslalani/gambit](https://ghcr.io/maaslalani/gambit).

### Play

You can play a game by running:

```
gambit
```

You can press <kbd>ctrl+f</kbd> to flip the board to give a better perspective
for the second player.

<br/>
<p align="center">
  <img width="90%" src="./assets/chess.gif?raw=true" alt="Terminal chess" />
</p>

### Networked Play

You can play chess with `gambit` over `ssh`.

```
ssh [<name>@]<host> -p <port> -t <room> [<password>]
```

You can host your own `gambit` SSH server with:

```
gambit serve
```

### Move

There are two ways to move in `gambit`:

* Type out the square the piece you want to move is on, then type out the square to which you want to move the piece.
* With the mouse, click on the target piece and target square.
