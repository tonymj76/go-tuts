package main

import "fmt"

func peroidicTable() map[int]map[string]map[string]string {

	elements := map[int]map[string]map[string]string{
		1: map[string]map[string]string{
			"H": map[string]string{
				"name":  "Hydrogen",
				"state": "gas",
			},
		},

		2: map[string]map[string]string{
			"He": map[string]string{
				"name":  "Helium",
				"state": "gas",
			},
		},

		3: map[string]map[string]string{
			"Li": map[string]string{
				"name":  "Lithium",
				"state": "solid",
			},
		},

		4: map[string]map[string]string{
			"Be": map[string]string{
				"name":  "Beryllium",
				"state": "solid",
			},
		},

		5: map[string]map[string]string{
			"B": map[string]string{
				"name":  "Boron",
				"state": "solid",
			},
		},

		6: map[string]map[string]string{
			"C": map[string]string{
				"name":  "Carbon",
				"state": "solid",
			},
		},

		7: map[string]map[string]string{
			"N": map[string]string{
				"name":  "Nitrogen",
				"state": "gas",
			},
		},

		8: map[string]map[string]string{
			"O": map[string]string{
				"name":  "Oxygen",
				"state": "gas",
			},
		},

		9: map[string]map[string]string{
			"F": map[string]string{
				"name":  "Fluorine",
				"state": "gas",
			},
		},

		10: map[string]map[string]string{
			"Ne": map[string]string{
				"name":  "Neon",
				"state": "gas",
			},
		},
	}
	return elements
}

// for me the best way is to use struct type accomplish this
func niceFmt(no int) string {
	elements := peroidicTable()
	ele, ok := elements[no]

	// check to know if element is pressent
	if !ok {
		return fmt.Sprintf("Element with atomic No: %d Not Found", no)
	}

	var (
		// global variable to return
		syEl string
		nEl  string
		stEl string
	)

	// for loop for the map items
	for symbol, valueMap := range ele {
		for name, state := range valueMap {
			syEl = symbol
			nEl = name
			stEl = state
		}
	}

	return fmt.Sprintf("Atomic NO: %d\nSymbol: %s\nName:\t%s\nState:\t%s\n", no, syEl, nEl, stEl)
}

// switch func that calls everything together
func allElement(ele interface{}) string {
	// variable for the items to be returned
	var (
		atomicNumber int
		symbol       string
		name         string
		state        string
	)

	switch f := ele.(type) {
	case string:
		if ele == "full" {
			element := peroidicTable()
			for atomic, valueMap := range element {
				for sym, valueMap2 := range valueMap {
					for nEl, sEl := range valueMap2 {
						atomicNumber = atomic
						symbol = sym
						name = nEl
						state = sEl
						return fmt.Sprintf("Atomic NO: %d\nSymbol: %s\nName:\t%s\nState:\t%s\n", atomicNumber, symbol, name, state)
					}
				}
			}
		}
		return fmt.Sprintf("%s is not a valid input try `full`", ele)

	case int:
		return niceFmt(int(f))
	default:
		return "Not found"
	}
}

func main() {
	fmt.Println(niceFmt(4))
	fmt.Println(allElement("full"))

	buffer := make([]int, 200)
	for i := 1; i < len(buffer); i++ {
		buffer[i]++
	}
	fmt.Println(buffer)
	fmt.Printf("%T-- %T\n", "something"[0], int("something"[0]))
}
