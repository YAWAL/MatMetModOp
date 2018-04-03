package Lab2

func NextFitAlgorithm(loads []Load) []Container {
	var containers []Container
	contNum := 1

	outer:
	for _, load := range loads {

		if max_load/load.Weight < 2 {
			var loadsInContainer = []Load{load}
			container := Container{Num: contNum, Loading: load.Weight, Loads: loadsInContainer}
			containers = append(containers, container)
			contNum++
			continue outer
		}

		if max_load/load.Weight >= 2 {
			if len(containers) == 0 {
				var loadsInContainer = []Load{load}
				container := Container{Num: contNum, Loading: load.Weight, Loads: loadsInContainer}
				containers = append(containers, container)
				contNum++
				continue outer

			}

			for _, container := range containers {
				if (max_load - container.Loading) >= load.Weight {
					loadsInContainer := container.Loads
					loadsInContainer = append(loadsInContainer, load)
					container = Container{Num: contNum, Loading: load.Weight, Loads: loadsInContainer}
					containers = append(containers, container)
					continue outer

				}
			}

		}

	}

	return containers
}
