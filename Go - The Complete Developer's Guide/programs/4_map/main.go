package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// fmt.Println(colors)
	printMap(colors)
}

func printMap(colorMap map[string]string) {
	for color, hex := range colorMap {
		fmt.Println(color, ": ", hex)
	}
}
