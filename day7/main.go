package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	var cmd Command
	var dir *Directory
	var start *Directory

	dir = &Directory{name: "/", directories: make(map[string]*Directory)}

	start = dir
	input := readFile()

	for _, l := range input {
		// It's a command: parse it
		if l[0] == '$' {
			cmd = commandParser(l)

			if cmd.name == "cd" {
				if cmd.arg == ".." {
					dir = dir.previous
				} else {
					dir = dir.directories[cmd.arg]
				}
			}
		} else {
			splitted := strings.Split(l, " ")
			// New directory found: add it if it doesn't exist
			if strings.HasPrefix(splitted[0], "dir") {
				if _, ok := dir.directories[splitted[1]]; !ok {
					dir.directories[splitted[1]] = &Directory{name: splitted[1], previous: dir, directories: make(map[string]*Directory)}
				}
			} else {
				size, _ := strconv.Atoi(splitted[0])

				dir.files = append(dir.files, File{name: splitted[1], size: size})
			}
		}
	}

	computeSizeRecursively(start)

	fmt.Printf("Root size: %d\nAnswer: %d\n", start.size, calculateAnswer(start))

	const maxSpace = 70000000
	const neededSpace = 30000000
	fmt.Printf("Smallest directory to delete: %d\n", smallestDirectoryToDelete(start, neededSpace-(maxSpace-start.size)))
}

func smallestDirectoryToDelete(dir *Directory, needed int) int {
	dirSizes := concatenateDirSizes(dir)
	smallest := dirSizes[0]

	for _, s := range dirSizes {
		if s < smallest && s > needed {
			smallest = s
		}
	}

	return smallest
}

func concatenateDirSizes(dir *Directory) []int {
	var sizes []int
	for _, d := range dir.directories {
		sizes = append(sizes, d.size)
		sizes = append(sizes, concatenateDirSizes(d)...)
	}

	return sizes
}

func computeSizeRecursively(dir *Directory) {
	for _, d := range dir.directories {
		computeSizeRecursively(d)
	}

	for _, f := range dir.files {
		dir.size += f.size
	}

	if dir.previous != nil {
		dir.previous.size += dir.size
	}
}

func calculateAnswer(dir *Directory) int {
	var total int
	for _, d := range dir.directories {
		if d.size <= 100000 {
			total += d.size
		}
		total += calculateAnswer(d)
	}

	return total
}

func commandParser(l string) Command {
	var cmd Command
	splitted := strings.Split(l[2:], " ")
	cmd.name = splitted[0]
	if len(splitted) > 1 {
		cmd.arg = splitted[1]
	} else {
		cmd.arg = ""
	}

	return cmd
}

func readFile() []string {
	f, _ := os.Open("input.txt")
	bytes, _ := io.ReadAll(f)

	return strings.Split(string(bytes), "\n")[1:]
}
