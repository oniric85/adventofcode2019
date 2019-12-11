package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// Object represents a space "thing" that orbits around some other Object
type Object struct {
	name    string
	objects []*Object
}

// We n
func (obj Object) sumPathLenghts(level int) int {
	sum := 0

	for _, child := range obj.objects {
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

		obj2, ok := objects[name2]

		if !ok {
			obj2 = &Object{name2, nil}
		}

		obj1, ok = objects[name1]

		if !ok {
			obj1 = &Object{name1, nil}
		}

		obj1.objects = append(obj1.objects, obj2)

		objects[name1] = obj1
		objects[name2] = obj2
	}

	return objects
}

func solve1(objectsMap map[string]*Object) int {
	root := objectsMap["COM"]

	// The result is the sum of the path lenghts from the root (COM)
	// to each node (the objects in the map)
	return root.sumPathLenghts(1)
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	objectsMap := convertInputToObjectsMap(string(content))

	result := solve1(objectsMap)

	fmt.Println("Checksum is", result)
}
