package main

import "fmt"

func main() {
	slice := [][]int{
		{88}, {22}, {88}, {22},
	}
	fmt.Println(maxDistance(slice))
}
