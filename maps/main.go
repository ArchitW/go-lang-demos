package main

import "fmt"

/*
maps similar to object in JS or Dict in python
 keys:value must be exact same type (not that keys must be same type as values.)

when to use struct or maps
Maps								Struct

k,v must have same types.			values can be of different types.
keys are indexed.					does not support indexing
reference type 						value type, changing passed struct doesnt change original
where to use :collection of
closely related properties


signature: map[string]string, map[int]string, map[int]int
*/
func main() {
	/*declaration*/

	// declaring a map "colors" where keys are string and values are string
	colors :=
		map[string]string{
			"red":    "#ff0000",
			"black":  "#000000",
			"orange": "#FFA500",
			"color3": "#CCCCCCC",
			"white":  "#FFFFFF",
		}

	//OR
	var colors2 map[string]string //Declares empty map
	//colors2["green"] = "GREEN"
	//OR
	colors3 := make(map[string]string)
	colors3["White"] = "#ffffff"

	//Delete value
	delete(colors, "red")

	fmt.Println(colors)
	fmt.Println(colors2)
	fmt.Println(colors3)

printMap(colors)
}

//Iterate through a map

func printMap(c map[string]string)  {
	for color, hex := range c{
		println("Color is :",color,"=> [hex] ", hex)
	}

}