package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// Object represents a space "thing" that orbits around some other Object
type Object struct {
	name     string
	parent   *Object
	children []*Object
}

func (obj Object) sumPathLenghts(level int) int {
	sum := 0

	for _, child := range obj.children {
		sum += level + child.sumPathLenghts(level+1)
	}

	return sum
}

func convertInputToObjectsMap(input string) map[string]*Object {
	orbits := strings.Split(input, "\n")

	objects := make(map[string]*Object)

	for _, line := range orbits {
		s := strings.Split(line, ")")
		name1, name2 := s[0], s[1]

		var obj1, obj2 *Object

		obj1, ok := objects[name1]

		if !ok {
			obj1 = &Object{name1, nil, nil}
		}

		obj2, ok = objects[name2]

		if !ok {
			obj2 = &Object{name2, nil, nil}
		}

		obj2.parent = obj1

		obj1.children = append(obj1.children, obj2)

		objects[name1] = obj1
		objects[name2] = obj2
	}

	return objects
}

func shortestPathBetweenTwoNodes(node1, node2 *Object) int {
	visitedNodes := make(map[string]int)
	firstBranchPathLength := 0

	// first we mark all the visited nodes from first node
	// to the root of the tree
	for node1 != nil {
		visitedNodes[node1.name] = firstBranchPathLength
		firstBranchPathLength++
		node1 = node1.parent
	}

	// then from the second node we start traversing the tree to the root
	// but we stop as we soon a node visited from the first node
	secondBranchPathLength := 0
	var lca string
	for node2 != nil {
		if _, ok := visitedNodes[node2.name]; ok {
			lca = node2.name
			break
		}
		node2 = node2.parent
		secondBranchPathLength++
	}

	return visitedNodes[lca] + secondBranchPathLength
}

func solve1(objectsMap map[string]*Object) int {
	root := objectsMap["COM"]

	// The result is the sum of the path lenghts from the root (COM)
	// to each node (the objects in the map)
	return root.sumPathLenghts(1)
}

func solve2(objectsMap map[string]*Object) int {
	node1 := objectsMap["YOU"].parent
	node2 := objectsMap["SAN"].parent

	return shortestPathBetweenTwoNodes(node1, node2)
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	objectsMap := convertInputToObjectsMap(string(content))

	checksum := solve1(objectsMap)
	fmt.Println("Checksum is", checksum)

	orbitalTransfers := solve2(objectsMap)
	fmt.Println("Required number of orbital transfers is", orbitalTransfers)
}
