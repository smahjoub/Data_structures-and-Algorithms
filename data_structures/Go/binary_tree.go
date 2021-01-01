package main

import (
	"fmt"
)

type node struct {
	value int
	left  *node
	right *node
}

func main() {
	t := &node{value: 8}

	fmt.Println("Building the tree...")
	t.insert(1)
	t.insert(2)
	t.insert(3)
	t.insert(4)
	t.insert(5)
	t.insert(6)
	t.insert(7)
	t.print()

	fmt.Println()
	fmt.Println()

	fmt.Println("Looking for 5...")
	if t.contains(5) {
		fmt.Println("5 was found")
	} else {
		fmt.Println("5 was not found")
	}

	fmt.Println()

	fmt.Println("Looking for 11...")
	if t.contains(11) {
		fmt.Println("11 was found")
	} else {
		fmt.Println("11 was not found")
	}

	fmt.Println()

	fmt.Println("Deleting 5 and 7...")
	t.delete(5)
	t.delete(7)

	fmt.Println()
	fmt.Println("Tree content after deleting 5 and 7.")
	t.print()

	fmt.Println()
	fmt.Println()

	fmt.Printf("min is %d\n", t.findMinimum())
	fmt.Printf("max is %d\n", t.findMaximum())
}

func (t *node) print() {

	if t != nil {
		t.left.print()
		fmt.Printf("%d", t.value)
		t.right.print()
	}
}

func (t *node) insert(value int) {

	if t.value == value {
		return
	}

	if t.value > value {
		if t.left == nil {

			t.left = &node{value: value}
		}

		t.left.insert(value)
	}

	if t.value < value {

		if t.right == nil {

			t.right = &node{value: value}
		}

		t.right.insert(value)
	}
}

func (t *node) contains(value int) bool {

	if t == nil {
		return false
	} else if value < t.value {
		return t.left.contains(value)
	} else if value > t.value {
		return t.right.contains(value)
	} else {
		return true
	}
}

func (t *node) delete(value int) *node {

	if t == nil {
		return nil
	}

	if value < t.value {
		t.left = t.left.delete(value)
		return t
	}
	if value > t.value {
		t.right = t.right.delete(value)
		return t
	}

	if t.left == nil && t.right == nil {
		t = nil
		return nil
	}

	if t.left == nil {
		t = t.right
		return t
	}
	if t.right == nil {
		t = t.left
		return t
	}

	smallestValOnRight := t.right
	for {
		//find smallest value on the right side
		if smallestValOnRight != nil && smallestValOnRight.left != nil {
			smallestValOnRight = smallestValOnRight.left
		} else {
			break
		}
	}

	t.value = smallestValOnRight.value
	t.right = t.right.delete(t.value)
	return t
}

func (t *node) findMaximum() int {
	if t.right == nil {
		return t.value
	}
	return t.right.findMaximum()
}

func (t *node) findMinimum() int {
	if t.left == nil {
		return t.value
	}
	return t.left.findMinimum()
}
