package main

import "fmt"

func main() {
	// รับ input
	var number int
	fmt.Print("ป้อนตัวเลข = ")
	fmt.Scanf("%d", &number)

	if number%2 == 0 {
		fmt.Println("เลขคู่")
	} else {
		fmt.Println("เลขคี่")

	}
}
