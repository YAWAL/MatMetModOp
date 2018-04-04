package Lab2

import (
	"log"
	"fmt"
)

// NFAcontainerCount counts how many containers are needed
func NFAcontainerCount(input []int, size int) {
	containerCount := 1
	counter := 0
	s := size
	for i := 0; i < len(input); i++ {
		if (s-input[i] > 0) {
			s -= input[i]
			continue
		} else {
			containerCount++
			s = size
			i--
		}
		counter++
	}
	fmt.Println("Обчислювальна складність: ", len(input))
	fmt.Println("Containers needed: ", containerCount)
	fmt.Println("Counter: ", counter)
}

// NFA implements Next Fit Algorithm for resolving Bin Packing Problem
func NFA(loads []Load) []Container {
	var containers []Container
	log.Printf("Init container state:%v", containers)
	contCount := 1
	counter := 0
outer:
	for _, load := range loads {
		log.Println("Entering for-cycle=====counter=>", counter)
		if Max_load/load.Weight < 2 {
			var loadsInContainer = []Load{load}
			container := Container{Num: contCount, Loading: load.Weight, Loads: loadsInContainer}
			containers = append(containers, container)
			contCount++
			continue outer
		} else {
			// if Max_load/load.Weight >= 2 {
			if len(containers) == 0 {
				var loadsInContainer = []Load{load}
				container := Container{Num: contCount, Loading: load.Weight, Loads: loadsInContainer}
				containers = append(containers, container)
				contCount++
				continue outer

			} else {
				for _, container := range containers {
					if (Max_load - container.Loading) >= load.Weight {
						loadsInContainer := container.Loads
						loadsInContainer = append(loadsInContainer, load)
						container = Container{Num: contCount, Loading: container.Loading + load.Weight, Loads: loadsInContainer}
						containers = append(containers, container)
						continue outer
					}
				}
			}

			// }
		}

		counter++
		log.Println("Exit for-cycle=====Counter=>", counter)

	}

	return containers
}
