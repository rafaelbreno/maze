package maze

// NewMaze generates a maze.
func NewMaze() Maze {
	m := Maze{}
	return m
}

// If I wanted to implement an auto-generating Maze, I would try to
// implement in a way that I don't have the whole maze generated, but
// instead, the maze is constantly increasing, keeping up with the
// player.
//
// With that in mind, I would have a "Room Generator":
// roomGenerator - this function will generate a room,
// it's possible "doors" depends on a fixed chance of generating:
// 	- An Exit
// 	- Door to a Monster
// 	- Door to another Room.
//
// So, with that, you will have a finite Maze, that generates
// depending on the player movements.
//
// But, If I had to implement a Maze generator, that would return me
// a Maze, I could use the same implementation above, recursively, until
// I'd generated an "exit".
//
//
// Another possibility, is to generate a "noise map", and on top of that
// generate a maze, each maze would be a little more random and unique
// compared to other, the only question here is:
// 	- "How would a program generate a Maze on top of a Noise map?"
//
// With a 2D Noise map, that is basically a Matrix, we can iterate columns
// and lines, depending on the value of those items, a set of [-1;+1], we
// can take an action, for example:
// 	- 0 => exit
// 	- [-1; 0) => Monster
// 	- (0;+1] => Another Room.
//
// That's a few approaches that I can think of right now, and all of them have a
// chance of not having an exit, so, the Generator could have a limit of runs,
// and if reached that limit without and "exit" it would generate one by force,
// or even, with each "run", it could increase the chance of generating an "exit".
