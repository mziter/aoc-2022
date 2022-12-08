package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/mziter/aoc-2022/pkg/fsnode"
)

const (
	FILESYSTEM_CAPACITY = 70000000
	REQUIRED_SPACE      = 30000000
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
}

func ParseInput(content string) *fsnode.Node {
	fs := fsnode.NewFilesystem()
	root := fs
	r := strings.NewReader(content)
	s := bufio.NewScanner(r)

	for s.Scan() {
		line := s.Bytes()
		if line[0] == '$' {
		START:
			switch {
			case bytes.Equal(line[2:4], []byte("ls")):
				line = processLS(fs, s)
				goto START
			case bytes.Equal(line[2:6], []byte("cd /")):
				fs = root
			case bytes.Equal(line[2:7], []byte("cd ..")):
				fs = fs.GetParent()
			case bytes.Equal(line[2:4], []byte("cd")):
				fs = fs.GetChild(string(line[5:]))
			}
		}
	}
	return root
}

// TODO: Create a scanner that doesn't allocate and return new strings!
func partOne(content string) string {
	fs := ParseInput(content)
	fs.CalculateSize()
	//PrettyPrintCD(root, fs)
	sum := getDirSumsAtMost(fs, 100000)
	return strconv.Itoa(sum)
}

func partTwo(content string) string {
	fs := ParseInput(content)
	fs.CalculateSize()
	//PrettyPrintCD(root, fs)
	needed := REQUIRED_SPACE - (FILESYSTEM_CAPACITY - fs.Size)
	min := getDirMinOver(fs, needed)
	return strconv.Itoa(min)
}

func getDirMinOver(n *fsnode.Node, threshold int) int {
	min := math.MaxInt
	m := getMin(n, threshold)
	if m < min {
		min = m
	}
	for _, c := range n.Children {
		m := getDirMinOver(c, threshold)
		if m < min {
			min = m
		}
	}
	return min
}

func getMin(n *fsnode.Node, threshold int) int {
	if !n.IsDir || n.Size < threshold {
		return math.MaxInt
	}
	return n.Size
}

func getDirSumsAtMost(n *fsnode.Node, maximum int) int {
	size := 0
	size += getSizeIfDirAtMost(n, maximum)
	for _, c := range n.Children {
		size += getDirSumsAtMost(c, maximum)
	}
	return size
}

func getSizeIfDirAtMost(n *fsnode.Node, size int) int {
	if !n.IsDir {
		return 0
	}
	if n.Size <= size {
		return n.Size
	}
	return 0
}

func processLS(fs *fsnode.Node, s *bufio.Scanner) []byte {
	var line []byte
	for s.Scan() {
		line = s.Bytes()
		if line[0] == '$' {
			return line
		}
		if bytes.Equal(line[0:3], []byte("dir")) {
			Name := line[4:]
			fs.AddChildDir(string(Name))
			continue
		}
		tokens := bytes.Split(line, []byte{' '})
		Name := string(tokens[1])
		size, err := strconv.Atoi(string(tokens[0]))
		if err != nil {
			panic(err)
		}
		fs.AddChildFile(Name, size)
	}
	return line
}
