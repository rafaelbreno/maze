package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/rafaelbreno/maze/maze"
)

const (
	sampleFolder = "samples/"
)

func main() {
	samples, err := ioutil.ReadDir(sampleFolder)
	if err != nil {
		panic(err)
	}

	for _, sample := range samples {
		b, err := os.ReadFile(fmt.Sprintf("%s%s", sampleFolder, sample.Name()))
		if err != nil {
			panic(err)
		}

		m := maze.Maze{}

		err = json.Unmarshal(b, &m)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Maze: %s\n", sample.Name())
		fmt.Printf("\t- %s \n", maze.GetExitPath(m))
	}
}
