package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/zellyn/adventofcode/util"
)

type file struct {
	name string
	size int
}

type dir struct {
	name    string
	parent  *dir
	subdirs map[string]*dir
	files   map[string]file
	size    int
}

func (d *dir) printTree() {
	d.printTreeIndent("")
}

func (d *dir) printTreeIndent(indent string) {
	fmt.Printf("%s- %s (dir)\n", indent, d.name)
	indent += "  "
	for _, f := range d.files {
		fmt.Printf("%s- %s (file, size=%d)\n", indent, f.name, f.size)
	}
	for _, s := range d.subdirs {
		s.printTreeIndent(indent)
	}
}

func (d *dir) mkdir(dirname string) *dir {
	if s, found := d.subdirs[dirname]; found {
		return s
	}

	s := &dir{
		name:    dirname,
		parent:  d,
		subdirs: make(map[string]*dir),
		files:   make(map[string]file),
	}

	d.subdirs[dirname] = s
	return s
}

func (d *dir) fullpath() string {
	if d.parent == nil {
		return "/"
	}
	return fmt.Sprintf("%s%s/", d.parent.fullpath(), d.name)
}

func (d *dir) cd(dirname string) *dir {
	if dirname == "/" {
		if d.parent == nil {
			return d
		}
		return d.parent.cd(dirname)
	}
	if dirname == ".." {
		if d.parent == nil {
			return d
		}
		return d.parent
	}
	if s, ok := d.subdirs[dirname]; ok {
		return s
	} else {
		panic(fmt.Sprintf("dir %s doesn't (yet) contain subdir %s", d.fullpath(), dirname))
	}
}

func (d *dir) populate(lines []string) {
	for _, line := range lines {
		if strings.HasPrefix(line, "dir ") {
			d.mkdir(line[4:])
		} else {
			parts := strings.Split(line, " ")
			if len(parts) != 2 {
				panic(fmt.Sprintf("weird file line: %q", line))
			}
			size, err := strconv.Atoi(parts[0])
			if err != nil {
				panic(fmt.Sprintf("weird file line %q: %v", line, err))
			}
			name := parts[1]
			d.files[name] = file{
				name: name,
				size: size,
			}
		}
	}
}

func (d *dir) postOrder(f func(*dir)) {
	for _, s := range d.subdirs {
		s.postOrder(f)
	}
	f(d)
}

func populate(inputs []string) *dir {
	root := &dir{
		name:    "/",
		subdirs: make(map[string]*dir),
		files:   make(map[string]file),
	}

	d := root

	pieces := util.SplitBefore[string](inputs, func(s string) bool { return strings.HasPrefix(s, "$") })

	for _, piece := range pieces {
		command := strings.Split(piece[0], " ")
		switch command[1] {
		case "cd":
			d = d.cd(command[2])
		case "ls":
			d.populate(piece[1:])
		default:
			panic(fmt.Sprintf("weird command: %q", piece[0]))
		}
	}

	return root
}

func (d *dir) calcTotalSizes() {
	d.postOrder(func(d *dir) {
		for _, f := range d.files {
			d.size += f.size
		}
		for _, s := range d.subdirs {
			d.size += s.size
		}
	})
}

func part1(inputs []string) (int, error) {
	root := populate(inputs)
	root.calcTotalSizes()
	total := 0
	root.postOrder(func(d *dir) {
		if d.size < 100000 {
			total += d.size
		}
	})
	return total, nil
}

func part2(inputs []string) (int, error) {
	total := 70000000
	required := 30000000
	root := populate(inputs)
	root.calcTotalSizes()

	currentFree := total - root.size
	additionalNeeded := required - currentFree

	min := total
	root.postOrder(func(d *dir) {
		if d.size >= additionalNeeded && d.size < min {
			min = d.size
		}
	})

	return min, nil
}

func run() error {
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
