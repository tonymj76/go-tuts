package main

import "fmt"

func peroidicTable() map[int]map[string]map[string]string {

	elements := map[int]map[string]map[string]string {
		1: map[string]map[string]string {
			"H": map[string]string {
				"name": "Hydrogen",
				"state": "gas",
			}, 
		},

		2: map[string]map[string]string {
			"He": map[string]string {
				"name": "Helium",
				"state": "gas",
			},
		},

		3: map[string]map[string]string {
			"Li": map[string]string {
				"name": "Lithium",
				"state": "solid",
			},
		},

		4: map[string]map[string]string {
			"Be": map[string]string {
				"name": "Beryllium",
				"state": "solid",
			},
		},

		5: map[string]map[string]string {
			"B": map[string]string {
				"name": "Boron",
				"state": "solid",
			},
		}, 

		6: map[string]map[string]string {
			"C": map[string]string {
				"name": "Carbon",
				"state": "solid",
			},
		},

		7: map[string]map[string]string {
			"N": map[string]string {
				"name": "Nitrogen",
				"state": "gas",
			},
		},

		8: map[string]map[string]string {
			"O": map[string]string {
				"name": "Oxygen",
				"state": "gas",
			},
		},

		9: map[string]map[string]string {
			"F": map[string]string {
				"name": "Fluorine",
				"state": "gas",
			},
		},

		10: map[string]map[string]string {
			"Ne": map[string]string {
				"name": "Neon",
				"state": "gas",
			},
		},
	}
	return elements
}
// for me the best way is to use struct type accomplish this

// func niceTable(no int) string{
// 	elements := peroidicTable()
// 	if ele, ok := elements[no]; ok {
// 		return fmt.Sprint("Atomic NO: %d\nSymbol: %s\nName:\t%s\nState:\t%s\n", no, )
// 	}
// }

func main() {
	elements := peroidicTable()
	if el, ok := elements[2]; ok {
		fmt.Println(el)
	}
	fmt.Println(len(elements))
}

