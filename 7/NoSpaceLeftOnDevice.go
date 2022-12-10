package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type node struct {
	name string
	size int
	isFile bool
	leafs map[string]*node
	root *node
}

func (n *node) GetRoot() *node {
	if n.root == nil {
		return n
	} else {
		return n.root.GetRoot()
	}
}

func (n *node) calculateSizeWithSizeDetect(sizeAtMost int, sizeMore int, searchResult *[]*node) int {
	if n.isFile {
		return n.size
	} else {
		size := 0
		if sizeMore == 0 {
			for _, v := range n.leafs {
				size += v.calculateSizeWithSizeDetect(sizeAtMost, 0, searchResult)
			}
			if size <= sizeAtMost {
				*searchResult = append(*searchResult, n)
			} 
		} else {
			for _, v := range n.leafs {
				size += v.calculateSizeWithSizeDetect(sizeAtMost, sizeMore, searchResult)
			}
			if size >= sizeMore {
				*searchResult = append(*searchResult, n)
			}
		}
		n.size = size
		return size
	}
}

func main(){
	
	//structure := []*node{}
	var actualNode *node

	readFile, err := os.Open("input.txt")
	
    if err != nil {
        fmt.Println(err)
    }

	fs := bufio.NewScanner(readFile)
	
	for fs.Scan() {
		line := strings.Fields(fs.Text())
		if line[0] == "$" {
			if line[1] == "cd" {
				if line[2] == ".." {
					actualNode = actualNode.root
				} else if line[2] == "/" {
					actualNode = &node{"/", 0, false, make(map[string]*node), nil}
				} else {
					actualNode = actualNode.leafs[line[2]]
				}
			}
		} else {
			if line[0] == "dir" {
				actualNode.leafs[line[1]] = &node{line[1], -1, false, make(map[string]*node), actualNode}
			} else {
				size, _ := strconv.Atoi(line[0])
				actualNode.leafs[line[1]] = &node{line[1], size, true, nil, actualNode}
			}
		}
	}
	
	readFile.Close()

	actualNode = actualNode.GetRoot()

	// Part One
	sizeAtMost := 100000
	searchResult := make([]*node, 0)
	totalSize := actualNode.calculateSizeWithSizeDetect(sizeAtMost, 0, &searchResult)
	fmt.Println("Total size:", totalSize)
	sum := 0
	for _,val := range searchResult {
		sum += val.size
	}
	fmt.Println("Total files size with size at most", sizeAtMost, "bytes:", sum)

	// Part Two
	searchResult = make([]*node, 0)
	sizeMore := totalSize - 40000000
	actualNode.calculateSizeWithSizeDetect(0, sizeMore, &searchResult)

	max := totalSize
	name := ""
	for _,val := range searchResult {
		if val.size < max {
			max = val.size
			name = string(val.name)
		}
	}
	fmt.Println("File with size more than", sizeMore, "bytes:", name, "with size", max)

}