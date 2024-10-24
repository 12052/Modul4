package main

import "fmt"

// Prosedur untuk mencetak deret bilangan
func cetakDeret(n_2222 int) {
	for n_2222 != 1 {
		fmt.Print(n_2222, " ")
		if n_2222%2 == 0 {
			n_2222 = n_2222 / 2
		} else {
			n_2222 = 3*n_2222 + 1
		}
	}
	fmt.Println(1)
}
func main() {
	var n int
	fmt.Scan(&n)
	cetakDeret(n)
}
