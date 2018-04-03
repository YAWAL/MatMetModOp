package main

import (
	"fmt"
	"github.com/YAWAL/MatMetModOp/util"
	"github.com/YAWAL/MatMetModOp/Lab2"
)

var input = []int{40, 29, 31, 17, 97, 71, 81, 75, 19, 27, 67, 56, 97, 53, 86, 65, 16, 83, 19, 24}

func main() {

	loads := util.GetLoads(input)

	fmt.Println("Loads: ", loads)
	containers := Lab2.NextFitAlgorithm(loads)

	for _, cont := range containers{
		fmt.Println(cont.Num, " :", cont.Loads)

	}
}
