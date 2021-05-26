package main

import (
	"./utils"
	"fmt"
)

func main() {
	result := utils.TwoSum([]int{2,7,11,15}, 9)
	fmt.Printf("%v", result)
}