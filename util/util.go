package util

import (
	"math"
	"log"
	"strconv"
	"github.com/YAWAL/MatMetModOp/Lab1"
	"github.com/YAWAL/MatMetModOp/Lab2"
)


func GetAlternatives(stringRow []string) []Lab1.Alternative {
	var intRow []int
	for _, row := range stringRow {
		num, _ := strconv.Atoi(row)
		intRow = append(intRow, num)
	}
	var alts []Lab1.Alternative
	for ind, rw := range intRow {
		var alt Lab1.Alternative
		name := "A" + strconv.Itoa(ind+1)
		alt.Name = name
		Q2 := math.Mod(float64(rw), 10)
		Q1 := (float64(rw) - Q2) / 10
		alt.Q1 = int(Q1)
		alt.Q2 = int(Q2)
		alt.IsDominated = true
		alts = append(alts, alt)
	}
	return alts
}

func PrintAlternatives(alts []Lab1.Alternative) {
	for _, alt := range alts {
		log.Printf("Alternative %v, Criterias: Q1=%v , Q2=%v , IsDominated: %v", alt.Name, alt.Q1, alt.Q2, alt.IsDominated)
	}
}

// GetLoads take input and return slice of loads for processing
func GetLoads(input []int) []Lab2.Load {
	var loads []Lab2.Load
	for num, inp := range input {
		load := Lab2.Load{Num:num+1, Load:inp}
		loads = append(loads, load)
	}
	return loads
}