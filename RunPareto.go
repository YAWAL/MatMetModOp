package main

import (
	"github.com/YAWAL/MatMetModOp/util"
	"github.com/YAWAL/MatMetModOp/Lab1"
)

//var row1 = []int{40, 29, 31, 17, 97, 71, 81, 75, 19, 27, 67, 56, 97, 53, 86, 65, 16, 83, 19, 24}
var row1 = []string{"79", "95", "04", "37", "92", "95", "12", "52", "70", "14"}

func main() {

	alts := util.GetAlternatives(row1)
	util.PrintAlternatives(alts)

	//pareto := Pareto(alts)
	 Lab1.ParetoAlt(alts)
	//PrintAlternatives(pareto)
}
