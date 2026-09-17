
// একটি n integer দেওয়া থাকবে।

// তোমার কাজ:

// for loop দিয়ে 1 থেকে n পর্যন্ত প্রতিটি সংখ্যা দেখবে।
// if/else দিয়ে প্রতিটি সংখ্যা even নাকি odd সেটা বের করবে।
// switch দিয়ে:
// even হলে "Even" print করবে
// odd হলে "Odd" print করবে
// শেষে মোট কতগুলো even এবং কতগুলো odd আছে সেটা দেখাবে।



package main

import "fmt"

func main() {
	fmt.Println(number(61))
}

func number(n int) any {
	var evenCount = 0
	var oddCount = 0
	var evenOdd = ""
	for i := 1; i <= n; i++ {

		if i%2 == 0 {
			oddCount += 1
			evenOdd = "Odd"
		} else {
			evenCount += 1
			evenOdd = "Even"
		}

		switch evenOdd {
		case "Odd":
			fmt.Println(i, " -> Odd")
		case "Even":
			fmt.Println(i, " -> Even")
		}

	}
	return fmt.Sprintf("Total Even: %d Total Odd: %d", evenCount, oddCount)

}
