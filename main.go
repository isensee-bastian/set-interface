package main

import (
	"fmt"
	"sort"
	"strings"
)

type Address struct {
	name   string
	street string
	zip    int
}

func (a Address) String() string {
	return fmt.Sprintf("%v | %v | %v", a.name, a.street, a.zip)
}

type Set struct {
	internalMap map[interface{}]bool
}

func NewSet(elements ...interface{}) Set {
	internalMap := map[interface{}]bool{}

	for _, element := range elements {
		internalMap[element] = true
	}

	return Set{internalMap}
}

func (s Set) Size() int {
	return len(s.internalMap)
}

func (s *Set) Add(element interface{}) {
	s.internalMap[element] = true
}

func (s *Set) Remove(element interface{}) {
	delete(s.internalMap, element)
}

func (s Set) Contains(element interface{}) bool {
	// Since we only store true values for keys that are existing,
	// we can just return the value. Missing key will result in default
	// value which is false for bool types.
	return s.internalMap[element]
}

func (s Set) Slice() []interface{} {
	slice := make([]interface{}, 0, len(s.internalMap))

	for element := range s.internalMap {
		slice = append(slice, element)
	}

	return slice
}

// Return elements ordered ascending.
func (s Set) String() string {
	slice := []string{}

	for element := range s.internalMap {
		switch element.(type) {
		case string:
			slice = append(slice, element.(string))
		case Address:
			slice = append(slice, element.(Address).String())
		default:
			fmt.Printf("Cannot convert %v to string\n", element)
		}
	}

	sort.Strings(slice)

	return strings.Join(slice, ", ")
}
