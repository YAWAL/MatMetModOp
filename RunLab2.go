package main

import (
	"fmt"
	"github.com/YAWAL/MatMetModOp/util"
	"github.com/YAWAL/MatMetModOp/Lab2"
)

var row1_ = []int{40, 29, 31, 17, 97, 71, 81, 75, 19, 27, 67, 56, 97, 53, 86, 65, 16, 83, 19, 24}
var row2_ = []int{28, 71, 32, 29, 3, 19, 70, 68, 8, 15, 40, 49, 96, 23, 18, 45, 46, 51, 21, 55}
var row3_ = []int{79, 88, 64, 28, 41, 50, 93, 51, 34, 64, 24, 14, 87, 56, 43, 91, 27, 65, 59, 36}

var moreThanHalf = []int{51, 61, 72, 83, 97, 71, 81, 75, 67, 56, 97, 53, 86, 65, 83}

func appendSlices(row1, row2, row3 []int) []int {
	var rows1to3 []int
	rows1to3 = append(row1, row2...)
	rows1to3 = append(rows1to3, row3...)
	return rows1to3
}

func main() {
	//allRows := appendSlices(row1_, row2_, row3_)

	loads := util.GetLoads(row1_)
	// loads := util.GetLoads(moreThanHalf)

	fmt.Println("Loads: ", loads)
	//containers := Lab2.NFA(loads)

	//util.PrintContainersInfo(containers)
	Lab2.NFAcontainerCount(row1_,Lab2.Max_load )
}
