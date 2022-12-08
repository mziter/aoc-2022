package fsnode

import (
	"fmt"
	"strings"
)

const (
	FILE_COLOR      = "\033[34m"
	DIR_COLOR       = "\033[32m"
	HIGHLIGHT_COLOR = "\033[35m"
	COLOR_RESET     = "\033[0m"
)

type Node struct {
	Name     string
	Size     int
	IsDir    bool
	Parent   *Node
	Children []*Node
}

func NewFilesystem() *Node {
	return &Node{
		Name:     "/",
		Size:     0,
		IsDir:    true,
		Children: make([]*Node, 0),
		Parent:   nil,
	}
}

func (n *Node) addChild(Name string, isDir bool, size int) {
	child := &Node{
		Name:     Name,
		Size:     size,
		IsDir:    isDir,
		Children: make([]*Node, 0),
		Parent:   n,
	}
	n.Children = append(n.Children, child)
}

func (n *Node) AddChildDir(Name string) {
	n.addChild(Name, true, 0)
}

func (n *Node) AddChildFile(Name string, size int) {
	n.addChild(Name, false, size)
}

func (n *Node) GetParent() *Node {
	if n.Parent == nil {
		panic("cannot move outside of outermost node")
	}
	return n.Parent
}

func (n *Node) GetChild(Name string) *Node {
	for _, c := range n.Children {
		if c.Name == Name {
			return c
		}
	}
	panic(fmt.Sprintf("could not find child with Name=%s within %s", Name, n.Name))
}

func (n *Node) CalculateSize() {
	if !n.IsDir {
		return
	} else {
		size := 0
		for _, c := range n.Children {
			c.CalculateSize()
			size += c.Size
		}
		n.Size = size
	}
}
func PrettyPrintCD(fs *Node, cd *Node) {
	fs.prettyPrintHighlight(0, cd)
}

func (n *Node) PrettyPrint() {
	n.prettyPrint(0)
}

func (n *Node) prettyPrint(depth int) {
	indent := strings.Repeat("  ", depth)
	fmt.Printf("%s%s - %s%s\n", getColor(n, nil), indent, n.String(), COLOR_RESET)
	for _, c := range n.Children {
		c.prettyPrint(depth + 1)
	}
}

func (n *Node) prettyPrintHighlight(depth int, highlight *Node) {
	indent := strings.Repeat("  ", depth)
	fmt.Printf("%s%s - %s%s\n", getColor(n, highlight), indent, n.String(), COLOR_RESET)
	for _, c := range n.Children {
		c.prettyPrintHighlight(depth+1, highlight)
	}
}

func getColor(n *Node, highlight *Node) string {
	if !n.IsDir {
		return FILE_COLOR
	}
	if highlight != nil && n == highlight {
		return HIGHLIGHT_COLOR
	}
	return DIR_COLOR
}

func (n *Node) String() string {
	if n.IsDir {
		return fmt.Sprintf("%s (dir, size=%d)", n.Name, n.Size)
	}
	return fmt.Sprintf("%s (file, size=%d)", n.Name, n.Size)
}
