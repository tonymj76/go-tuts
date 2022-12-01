package module1

func InsertionSort(list []int) {
	var sorted []int
	for _, item := range list {
		sorted = insert(sorted, item)
	}
	for i, sortedItem := range sorted {
		list[i] = sortedItem
	}
}

func insert(sortItems []int, item int) []int {
	for i, sortItem := range sortItems {
		if item < sortItem {
			return append(sortItems[:i], append([]int{item}, sortItems[:i]...)...)
		}
	}
	return append(sortItems, item)
}

// https://github.com/golang/go/wiki/SliceTricks
// https://ueokande.github.io/go-slice-tricks/
