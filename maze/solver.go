package maze

import (
	"strings"
)

// GetExitPath receives a mao[string]any and prints the result
// 	- If there's an exit, it returns the formatted order to reach the exit:
// 		- e.g: left -> forward -> upstairs
// 	- If there IS NOT an exit, it returns: "sorry"
func GetExitPath(m Maze) string {
	path := findShortestPath(m, true)
	if len(path) == 0 {
		return "sorry"
	}

	return strings.Join(path, " -> ")
}

func findShortestPath(m Maze, start bool) []string {
	paths := [][]string{}
	for k, v := range m {
		switch i := v.(type) {
		case string:
			switch {
			case i == "exit" && !start:
				path := []string{k}
				return path
			case i == "exit" && start:
				paths = append(paths, []string{k})
			case i != "exit" && !start:
				return []string{}
			}
		case map[string]any:
			returnedPath := findShortestPath(i, false)
			lenReturned := len(returnedPath)
			switch {
			case !start && lenReturned == 0:
				return []string{}
			case !start && lenReturned > 0:
				return append([]string{k}, returnedPath...)
			case start && lenReturned > 0:
				path := append([]string{k}, returnedPath...)
				paths = append(paths, path)
				continue
			case start && lenReturned == 0:
				continue
			}
		}
	}

	var shortestPath []string

	for pos, path := range paths {
		if pos == 0 {
			shortestPath = path
			continue
		}
		if len(shortestPath) > len(path) {
			shortestPath = path
		}
	}

	return shortestPath
}
