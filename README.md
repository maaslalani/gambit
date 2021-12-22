# Gambit

Chess board in your terminal.

<br/>
<p align="center">
  <img width="90%" src="./chess.gif?raw=true" alt="Terminal chess" />
</p>
<br/>

### Warning

`gambit` does not have a large number of features at the
moment. I plan on adding a chess engine, mouse support,
timers, networked play, game replays, etc...

### Move

You may use the mouse to move your pieces. You may also
type out the square the piece you want to move is on
followed by the square to which you want to move the piece.

Like a real chess board, `gambit` supports illegal moves.

### Players

`gambit` supports local and networked play. You can play a
local game by running `gambit` and moving the pieces. You
can flip the board by pressing <kbd>ctrl+f</kbd> to allow
the second player to go.

For networked play (not available yet), both players can
run `gambit unique-room-id`, this will connect both players
to a shared room in which both can take turns making moves.
