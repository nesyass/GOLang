package main

import "fmt"

func main() {
	var n, i, luas, keliling int

	fmt.Scan(&n)
	// menyimpan dalam slice
	var sisiPersegi []int
	for i = 0; i < n; i++ {
		var sisi int
		fmt.Scan(&sisi)
		sisiPersegi = append(sisiPersegi, sisi)
	}

	// menghitung berdasarkan masukan yang sudah di simpan
	for i = 0; i < n; i++ {
		luas = sisiPersegi[i] * sisiPersegi[i]
		keliling = 4 * sisiPersegi[i]
		fmt.Printf("%d %d\n", luas, keliling)
	}
}

//algoritma menghitung keliling dan luas persegi
