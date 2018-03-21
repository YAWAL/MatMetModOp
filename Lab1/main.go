package main

import (
	"math"
	"strconv"
	"log"
)

type Alternative struct {
	Name        string
	Q1, Q2      int
	IsDominated bool
}

func Pareto(alts []Alternative) []Alternative {
	var pareto []Alternative
	//counter := 0
	log.Println("START PROCEDURE")

	log.Println("Paretos before 1st for: ", pareto)
	for ind, alt := range alts {

		for i := ind; i < len(alts)-1; i++ {

			if alt.Q1 >= alts[i+1].Q1 && alt.Q2 >= alts[i+1].Q2 {
				log.Printf("Iteration %v INDEX %v", i, ind)

				log.Printf("Alternative %v: Q1 %v , Q2 %v", alt.Name, alt.Q1, alt.Q2)
				log.Printf("Alternative %v: Q1 %v , Q2 %v", alts[i+1].Name, alts[i+1].Q1, alts[i+1].Q2)
				log.Printf("Alternative %v dominated on %v", alt.Name, alts[i+1].Name)
				alt.IsDominated = true
				alts[i+1].IsDominated = false
			} else {
				//log.Printf("Alternative %v: dominant on %v", alts[i].Name , alt.Name )

				alt.IsDominated = false

			}

		}

		if alt.IsDominated == true {
			pareto = append(pareto, alt)
		}
		//log.Println("Paretos in 1st for: ", pareto)

	}
	log.Println("Paretos after 1st for: ", pareto)

	log.Println("END PROCEDURE")

	return pareto
}

func ParetoAlt(alts []Alternative) []Alternative {
	var pareto = alts
	log.Printf("Paretos before 1st for: %v", pareto)

	log.Println("START PROCEDURE")

	for i := 0; i < len(alts)-1; i++ {

		//for j := 1; j < len(pareto)-1; j++ {

			if pareto[i].Q1 <= pareto[i+1].Q1 && pareto[i].Q2 <= pareto[i+1].Q2 {
				pareto[i+1].IsDominated = true
				pareto := append(pareto[:i], pareto[i+1:]...)
				log.Printf("Iteration %v, Paretos %v ", i, pareto)

			}

		//}
		//log.Println("Paretos in for: ", pareto)
	}

	log.Println("Paretos after for: ", pareto)

	log.Println("END PROCEDURE")

	return pareto
}

func GetAlternatives(row []int) []Alternative {
	var alts []Alternative
	for ind, rw := range row {
		var alt Alternative
		name := "A" + strconv.Itoa(ind+1)
		alt.Name = name
		Q2 := math.Mod(float64(rw), 10)
		Q1 := (float64(rw) - Q2) / 10
		alt.Q1 = int(Q1)
		alt.Q2 = int(Q2)
		alt.IsDominated = false
		alts = append(alts, alt)
	}
	return alts
}

func PrintAlternatives(alts []Alternative) {
	for _, alt := range alts {
		log.Printf("Alternateve %v, Criterias: Q1=%v , Q2=%v , IsDominated: %v", alt.Name, alt.Q1, alt.Q2, alt.IsDominated)
	}
}
