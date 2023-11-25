package main

import "fmt"

func main() {
	var n, d int

	fmt.Scan(&n)

	for d = 1; d <= n; d++ {
		if n%d == 0 {
			fmt.Printf("%d: true\n", d)
		} else {
			fmt.Printf("%d: false\n", d)
		}
	}
}

//program untuk menampilkan faktor dari suatu bilangan
//masukan terdiri dari sebuah bilangan bulat positif n
//keluaran terdiri dar kumpulan pasangan nilai bulat d dan boolean s yang dipisahkan oleh newline. dimana d adalah bilangan yang mungkin menjadi faktor dari n
