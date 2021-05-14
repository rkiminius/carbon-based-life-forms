package main

import "fmt"

func main() {
	fmt.Println("hello")
	var min = Mineral{
		"topaz",
		MINERAL_STATE_LIQUID,
		10,
	}
	fmt.Println(min)
	err := min.Melt()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(min)
}