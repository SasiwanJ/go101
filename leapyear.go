package main

import "fmt"

// func main() {
// 	var year int
// 	fmt.Print("Insert Year = ")
// 	fmt.Scanf("%d", &number)

// 	if year%4 == 0 {
// 		fmt.Println("เลขคู่")
// 	}
// 	else {
// 		year%100 == 0
// 		fmt.Println("เลขคี่")
// 	}
// 	if year%400 == 0 {
//         fmt.Println("เรรรรรร")
//     }

// }

func isLeapYear(year int) bool {
	var ans bool = false
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				ans = true
			}
		}
	}
	return ans
}

func main() {
	year := 2025 // Example year to check
	if isLeapYear(year) {
		fmt.Println(year, "ใช่จ้า")
	} else {
		fmt.Println(year, "ไม่ใช่จ้า")
	}
}
