package main

import (
	"fmt"
	"github.com/YAWAL/MatMetModOp/util"
)

var loads = []int{40, 29, 31, 17, 97, 71, 81, 75, 19, 27, 67, 56, 97, 53, 86, 65, 16, 83, 19, 24}

func main()  {

	fmt.Println(util.GetLoads(loads))


}
