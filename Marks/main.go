
package main

import "fmt"

func main() {
	marksArr := []int{85, 72, 45, 33, 91, 60, 28}

	fmt.Println(marks(marksArr))
}

func marks(mark []int) any {

	var grade = ""
	var aCount = 0
	var bCount = 0
	var cCount = 0
	var dCount = 0
	var fCount = 0

	for _, val := range mark {

		if val >= 80 {
			grade = "A"
			aCount += 1
		} else if val <= 79 && val >= 70 {
			grade = "B"
			bCount += 1
		} else if val <= 69 && val >= 60 {
			grade = "C"
			cCount += 1
		} else if val <= 59 && val >= 40 {
			grade = "D"
			dCount += 1
		} else if val <= 39 && val > 0 {
			grade = "F"
			fCount += 1
		}
		switch grade {
		case "A":
			fmt.Println(val, "-> A -> Excellent")
		case "B":
			fmt.Println(val, "-> B -> Very Good")
		case "C":
			fmt.Println(val, "-> C -> Good")
		case "D":
			fmt.Println(val, "-> D -> Passed")
		case "F":
			fmt.Println(val, "-> F -> Failed")
		}
	}
	return fmt.Sprintf("Total A grade %d, Total B grade %d, Total C grade %d, Total D grade %d, Total F grade %d", aCount, bCount, cCount, dCount, fCount)
}
