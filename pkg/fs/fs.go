package fs

type FileSystem struct {
	data   []fsData
	cwdIdx int
}

type fsData struct {
	name     string
	size     int
	isDir    bool
	parent   int
	children []int
}

func New(capacity int) *FileSystem {
	fs := &FileSystem{
		data: make([]fsData, 0, capacity),
	}
	root := fsData{
		name:  "/",
		isDir: true,
	}
	fs.data = append(fs.data, root)
	return fs
}

func (fs *FileSystem) AddDir(name string) int {
	dir := fsData{
		name:     name,
		isDir:    true,
		parent:   fs.cwdIdx,
		children: make([]int, 0),
	}
	return fs.attachChild(dir)
}

func (fs *FileSystem) AddFile(name string, size int) int {
	f := fsData{
		name:     name,
		isDir:    false,
		size:     size,
		parent:   fs.cwdIdx,
		children: make([]int, 0),
	}
	return fs.attachChild(f)
}

func (fs *FileSystem) attachChild(f fsData) int {
	idx := len(fs.data)
	fs.data = append(fs.data, f)
	fs.data[fs.cwdIdx].children = append(fs.data[fs.cwdIdx].children, idx)
	return idx
}
