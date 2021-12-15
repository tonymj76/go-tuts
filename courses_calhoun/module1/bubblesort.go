package module1

func BubbleSort(list []int) []int {
	for x := 0; x < len(list); x++ {
		sweep := false
		for i := 0; i < len(list)-1-x; i++ {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
				sweep = true
			}
		}
		if !sweep {
			break
		}
	}
	return list
}

func sweep(list []int) {
	for i := 0; i < len(list)-1; i++ {
		if list[i] > list[i+1] {

			// this works because right hand side get evaluated first before assigning
			list[i], list[i+1] = list[i+1], list[i]

			// this won't work because you have change what list[i] is
			//list[i] = list[i+1]
			//list[i+1] = list[i]
		}
	}
}
