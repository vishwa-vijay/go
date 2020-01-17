package main

import (
	"container/list"
	"fmt"

	"github.com/vishwa-vijay/go/raja"
)

type cord struct {
	x int
	y int
}

var successPath = list.New()

func main() {
	io := raja.NewIo()
	array, rows, cols := io.ReadArray("Enter 2 diamention array.")
	println("Array is ____")
	raja.PrintArray(array)
	fmt.Println()

	path := list.New()
	start := &cord{x: 0, y: 0}

	explore(array, path, start, rows, cols)
	printPaths(rows * cols)
}

func explore(array [][]int, path *list.List, position *cord, rows int, cols int) {
	element := path.PushBack(position)
	if position.x > -1 && position.y > -1 && position.x < rows && position.y < cols && !isCycle(path) {
		if array[position.x][position.y] == 9 {
			newPath := list.New()
			for e := path.Front(); e != nil; e = e.Next() {
				newPath.PushBack(e.Value)
			}
			successPath.PushBack(newPath)
		} else if array[position.x][position.y] == 1 {
			explore(array, path, &cord{x: position.x - 1, y: position.y}, rows, cols)
			explore(array, path, &cord{x: position.x, y: position.y - 1}, rows, cols)
			explore(array, path, &cord{x: position.x + 1, y: position.y}, rows, cols)
			explore(array, path, &cord{x: position.x, y: position.y + 1}, rows, cols)
		}
	}
	path.Remove(element)
}

func isCycle(path *list.List) bool {
	len := path.Len()
	if len > 1 {
		i := 0
		cx := path.Back().Value.(*cord).x
		cy := path.Back().Value.(*cord).y
		for e := path.Front(); e != nil && i < len-1; e = e.Next() {
			cp := e.Value.(*cord)
			if cp.x == cx && cp.y == cy {
				return true
			}
			i++
		}
	}
	return false
}

func printPaths(shortest int) {

	if successPath.Len() > 0 {
		fmt.Println("Following are the valid path to reach 9")
		for e := successPath.Front(); e != nil; e = e.Next() {
			lp := e.Value.(*list.List)
			clen := lp.Len() - 1
			if shortest > clen {
				shortest = clen
			}
			printList(lp)
		}
		fmt.Println("\nShortest path to reach 9 needs ", shortest, "steps")
		for e := successPath.Front(); e != nil; e = e.Next() {
			lp := e.Value.(*list.List)
			clen := lp.Len() - 1
			if clen == shortest {
				printList(lp)
			}
		}
	} else {
		fmt.Println("There is no way to reach")
	}
}

func printList(lp *list.List) {
	first := true
	fmt.Print("Steps ", (lp.Len() - 1), " ")
	for l := lp.Front(); l != nil; l = l.Next() {
		if first {
			first = false
		} else {
			fmt.Print("->")
		}
		fmt.Print(*l.Value.(*cord))
	}
	fmt.Println()
}
