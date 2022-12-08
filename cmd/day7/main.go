package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	FILESYSTEM_CAPACITY = 70000000
	REQUIRED_SPACE      = 30000000

	FILE_COLOR      = "\033[34m"
	DIR_COLOR       = "\033[32m"
	HIGHLIGHT_COLOR = "\033[35m"
	COLOR_RESET     = "\033[0m"
)

type FSNode struct {
	name     string
	size     int
	isDir    bool
	parent   *FSNode
	children []*FSNode
}

func NewFilesystem() *FSNode {
	return &FSNode{
		name:     "/",
		size:     0,
		isDir:    true,
		children: make([]*FSNode, 0),
		parent:   nil,
	}
}

func (n *FSNode) addChild(name string, isDir bool, size int) {
	child := &FSNode{
		name:     name,
		size:     size,
		isDir:    isDir,
		children: make([]*FSNode, 0),
		parent:   n,
	}
	n.children = append(n.children, child)
}

func (n *FSNode) AddChildDir(name string) {
	n.addChild(name, true, 0)
}

func (n *FSNode) AddChildFile(name string, size int) {
	n.addChild(name, false, size)
}

func (n *FSNode) GetParent() *FSNode {
	if n.parent == nil {
		panic("cannot move outside of outermost node")
	}
	return n.parent
}

func (n *FSNode) GetChild(name string) *FSNode {
	for _, c := range n.children {
		if c.name == name {
			return c
		}
	}
	panic(fmt.Sprintf("could not find child with name=%s within %s", name, n.name))
}

func (n *FSNode) CalculateSize() {
	if !n.isDir {
		return
	} else {
		size := 0
		for _, c := range n.children {
			c.CalculateSize()
			size += c.size
		}
		n.size = size
	}
}
func PrettyPrintCD(fs *FSNode, cd *FSNode) {
	fs.prettyPrintHighlight(0, cd)
}

func (n *FSNode) PrettyPrint() {
	n.prettyPrint(0)
}

func (n *FSNode) prettyPrint(depth int) {
	indent := strings.Repeat("  ", depth)
	fmt.Printf("%s%s - %s%s\n", getColor(n, nil), indent, n.String(), COLOR_RESET)
	for _, c := range n.children {
		c.prettyPrint(depth + 1)
	}
}

func (n *FSNode) prettyPrintHighlight(depth int, highlight *FSNode) {
	indent := strings.Repeat("  ", depth)
	fmt.Printf("%s%s - %s%s\n", getColor(n, highlight), indent, n.String(), COLOR_RESET)
	for _, c := range n.children {
		c.prettyPrintHighlight(depth+1, highlight)
	}
}

func getColor(n *FSNode, highlight *FSNode) string {
	if !n.isDir {
		return FILE_COLOR
	}
	if highlight != nil && n == highlight {
		return HIGHLIGHT_COLOR
	}
	return DIR_COLOR
}

func (n *FSNode) String() string {
	if n.isDir {
		return fmt.Sprintf("%s (dir, size=%d)", n.name, n.size)
	}
	return fmt.Sprintf("%s (file, size=%d)", n.name, n.size)
}

//go:embed input.txt
var input string

func main() {
	fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
}

func ParseInput(content string) *FSNode {
	fs := NewFilesystem()
	root := fs
	r := strings.NewReader(content)
	s := bufio.NewScanner(r)

	for s.Scan() {
		line := s.Text()
		if line[0] == '$' {
		START:
			switch {
			case line[2:4] == "ls":
				line = processLS(fs, s)
				goto START
			case line[2:6] == "cd /":
				fs = root
			case line[2:7] == "cd ..":
				fs = fs.GetParent()
			case line[2:4] == "cd":
				fs = fs.GetChild(line[5:])
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
	needed := REQUIRED_SPACE - (FILESYSTEM_CAPACITY - fs.size)
	min := getDirMinOver(fs, needed)
	return strconv.Itoa(min)
}

func getDirMinOver(n *FSNode, threshold int) int {
	min := math.MaxInt
	m := getMin(n, threshold)
	if m < min {
		min = m
	}
	for _, c := range n.children {
		m := getDirMinOver(c, threshold)
		if m < min {
			min = m
		}
	}
	return min
}

func getMin(n *FSNode, threshold int) int {
	if !n.isDir || n.size < threshold {
		return math.MaxInt
	}
	return n.size
}

func getDirSumsAtMost(n *FSNode, maximum int) int {
	size := 0
	size += getSizeIfDirAtMost(n, maximum)
	for _, c := range n.children {
		size += getDirSumsAtMost(c, maximum)
	}
	return size
}

func getSizeIfDirAtMost(n *FSNode, size int) int {
	if !n.isDir {
		return 0
	}
	if n.size <= size {
		return n.size
	}
	return 0
}

func processLS(fs *FSNode, s *bufio.Scanner) string {
	var line string
	for s.Scan() {
		line = s.Text()
		if line[0] == '$' {
			return line
		}
		if line[0:3] == "dir" {
			name := line[4:]
			fs.AddChildDir(name)
			continue
		}
		tokens := strings.Split(line, " ")
		name := tokens[1]
		size, err := strconv.Atoi(tokens[0])
		if err != nil {
			panic(err)
		}
		fs.AddChildFile(name, size)
	}
	return line
}
