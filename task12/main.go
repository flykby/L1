package main

import "fmt"

type Set struct {
	data []string
}

func NewSet(elems []string) *Set {
	set := make(map[string]struct{})

	for _, v := range elems {
		set[v] = struct{}{}
	}
	result := make([]string, 0, len(set))
	for k := range set {
		result = append(result, k)
	}
	
	return &Set{data: result}
}

// Other methods
// ...
// ...


func main() {
	animals := []string {"cat", "cat", "dog", "cat", "elephant"}
	setAnimals := *NewSet(animals)
	fmt.Println(setAnimals.data)
}