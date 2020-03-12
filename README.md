# Minisweeper

This is a test application to play the classic game of [Minesweeper](https://en.wikipedia.org/wiki/Minesweeper_(video_game))

The server side is developed using Golang. The client has been created using React

The game can be played at [Play](https://minesweeperfulla.herokuapp.com/)

The user recognition is very basic, and it's based on the clientIP; so you will be able to resume an existent game when you are joining the Minesweeper application from the same IP at which you started the game.
However, the game just allows a few game instances, so maybe your game will be discarded if many players joined a new game after you stopped playing.

Also, there is no persistence yet, so your current game will be lost if the server resets.

When persistence is implemented, the "exported" format can be used to save the `client board` it in the database as a json.
The `board` can be marshaled to json as it is.

I run out of time to implement the "flag" action, but as it's not a fundamental action I decided to add that feature later.

The client is very dull, but I plan to improve it in time.
