package main

func main() {

	// colors := make(map[int]string)
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}
	colors["white"] = "#ffffff"
	delete(colors, "red")

	printMap(colors)

}

func printMap(c map[string]string) {
	for key, value := range c {
		println(key, value)
	}
}
