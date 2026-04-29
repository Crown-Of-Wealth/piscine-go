package piscine

func SortIntegerTable(table []int) {
	for i := 1; i < len(table); i++ {
		for j := 0; j < len(table)-i; j++ {
			if table[j] > table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}
}

func DescIntegerTable(table []int) {
	for i := 1; i < len(table); i++ {
		swapped := false

		for j := 0; j < len(table)-i; j++ {
			if table[j] < table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
}