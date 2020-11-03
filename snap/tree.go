package snap

type Tree struct {
	file     *File
	level    int
	children []*Tree
}

func (t *Tree) generateRoot(file *File) {
	t = new(Tree)
	t.level = 0
	t.children = nil
	t.file = file
}

func (t *Tree) addChild() *Tree {
	if t.children == nil {
		t.children = make([]*Tree, 0, 0)
	}

	child := &Tree{
		file:     nil,
		level:    t.level + 1,
		children: nil,
	}

	t.children = append(t.children, child)
	return child
}

func (t *Tree) getLevel() int {
	return t.level
}

func (t *Tree) getChildrenList() []*Tree {
	return t.children
}

func (t *Tree) getChildByIndex(i int) *Tree {
	if i >= len(t.children) {
		return nil
	} else {
		return t.children[i]
	}
}

func (t *Tree) getData() *File {
	return t.file
}
