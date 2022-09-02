package maze

import (
	"encoding/json"
	"fmt"
)

// Maze is the data struct representing
// a given maze.
type Maze map[string]any

// PrettyPrint - prints a indented JSON
// representation of Maze.
func (m *Maze) PrettyPrint() {
	b, err := json.MarshalIndent(*m, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
