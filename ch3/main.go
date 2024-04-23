package main

import "fmt"

func main() {
	//	len cap functions
	var x []int
	fmt.Println(x, len(x), cap(x))
	x = append(x, 10)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 20)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 30)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 40)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 50)
	fmt.Println(x, len(x), cap(x))

	//	make
	y := make([]int, 10)
	fmt.Println(y, len(y), cap(y))
	y = append(y, 66)
	fmt.Println(y, len(y), cap(y))
	// appending this to an already zero populated slice will double the capacity as per slice rules

	// you can specify capacity using make, and also a slice with length 0 but capacity >0
	// so you must append this slice in order to add values, you can't directly index it
	y = make([]int, 0, 10)
	fmt.Println(y, len(y), cap(y))
	y = append(y, 1, 2, 3, 4, 5)
	fmt.Println(y, len(y), cap(y))

	//	clear
	//	takes a slice, sets all of its elements to their zero value, length remains unchanged
	//	some examples with strings and ints below
	clear(y)
	fmt.Println(y, len(y), cap(y))

	s := []string{"a", "b", "c"}
	fmt.Println(s)
	clear(s)
	fmt.Println(s)

	//	declaring your slice
	//	The primary goal is to minimise the number of times teh slice needs to grow

	//	slice nil

	//	If slice will never grow, use var with no assigned value, to create a nil slice
	var storageSlice []int

	// this differs from an empty slice literal
	//		var x = []int{}
	//	this has a special use case for converting a slice to JSON only when encoding/decoding JSON
	fmt.Println(storageSlice, len(storageSlice), cap(storageSlice))

	// slice literal
	//	if you have some default or starting values, use a slice literal
	data := []int{2, 4, 6, 8}
	fmt.Println(data)

	//	if you have a good idea what size you will need but not the immediate values, use make

	//	using slice as a buffer, special use case - io operations
	buffer := make([]int, 50)
	fmt.Println(len(buffer), cap(buffer))

	//	if you are sure you know the exact size you want - usually transforming values from one to another
	//	specify length and index into the slice to set values

	// in other scenarios use make with a zero length, specified capacity
	//	there are no extra zero values, and number of items is larger than expected code wont panic
	alternative := make([]int, 0, 50)
	fmt.Println(alternative, len(alternative), cap(alternative))

}
