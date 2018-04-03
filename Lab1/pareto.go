package Lab1

import (
	"log"
)

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
	var paretoDom []Alternative
	log.Printf("Paretos before 1st for: %v", pareto)

	log.Println("START PROCEDURE")

	for i := 0; i < len(pareto); i++ {

		for j := 1; j < len(pareto)-1; j++ {

			if pareto[i].Q1 <= pareto[j].Q1 && pareto[i].Q2 <= pareto[j].Q2 {
				pareto[i].IsDominated = false
				//pareto := append(pareto[:i], pareto[i+1:]...)
				log.Printf("Iteration %v, Paretos %v ", i, pareto)

			}

		}
	}

	for _, pd := range pareto {
		if pd.IsDominated == true {
			paretoDom = append(paretoDom, pd)
		}

	}

	log.Println("paretoDom: ", paretoDom)

	log.Println("Paretos after for: ", pareto)

	log.Println("END PROCEDURE")

	return pareto
}

