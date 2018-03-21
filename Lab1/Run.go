package main

//var row1 = []int{40, 29, 31, 17, 97, 71, 81, 75, 19, 27, 67, 56, 97, 53, 86, 65, 16, 83, 19, 24}
var row1 = []int{79, 95, 14, 37, 92, 95, 12, 52, 70, 14}

func main() {

	alts := GetAlternatives(row1)
	//PrintAlternatives(alts)

	//pareto := Pareto(alts)
	 ParetoAlt(alts)
	//PrintAlternatives(pareto)
}
