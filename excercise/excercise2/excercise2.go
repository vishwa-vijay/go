package main

import (
	"container/list"
	"fmt"
	"strings"

	"github.com/vishwa-vijay/go/raja"
)

type marker struct {
	city1    int
	city2    int
	distance int
}

func main() {
	io := raja.NewIo()
	array, rows, _ := io.ReadArray("Ender City-1 index, City-2 index, Distance Between the Cities")
	//array := [][]int{{1, 2, 3}, {2, 3, 4}, {2, 6, 2}, {6, 4, 6}, {6, 5, 5}}
	//rows := 5

	markers := list.New()
	cities := []int{}

	for i := 0; i < rows; i++ {
		markers.PushBack(marker{city1: array[i][0], city2: array[i][1], distance: array[i][2]})
		cities = addCity(&cities, array[i][0])
		cities = addCity(&cities, array[i][1])
	}

	for _, city := range cities {
		explore(markers, city, 0, []int{})
	}

	max := 0
	for _, v := range pathmap {
		if max < v {
			max = v
		}
	}

	for k, v := range pathmap {
		if max == v {
			fmt.Println("distance", v, " : ", k)
		}
	}

}

var pathmap = make(map[string]int, 15)

func explore(markers *list.List, city int, distance int, way []int) {
	for _, c := range way {
		if c == city {
			//fmt.Println("Cycle is created..", strings.Trim(strings.Replace(fmt.Sprint(way), " ", "->", -1), "[]"), city)
			return
		}
	}
	way = append(way, city)
	//fmt.Println("exploring", "city", city, "distance", distance, "path", way)
	if distance > 0 {
		pathmap[strings.Trim(strings.Replace(fmt.Sprint(way), " ", " -> ", -1), "[]")] = distance
	}

	for markerElement := markers.Front(); markerElement != nil; markerElement = markerElement.Next() {
		mark := markerElement.Value.(marker)
		if mark.city1 == city {
			explore(markers, mark.city2, mark.distance+distance, way)
		}
		if mark.city2 == city {
			explore(markers, mark.city1, mark.distance+distance, way)
		}
	}

}

func addCity(cities *[]int, city int) []int {
	citieso := *cities
	for _, i := range citieso {
		if i == city {
			return citieso
		}
	}
	newCities := append(citieso, city)
	return newCities
}
