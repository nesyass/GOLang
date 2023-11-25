package main

import "fmt"

func main() {
	var n, i int
	var masukan string

	fmt.Scan(&n, &masukan)
	for i = 1; i < n; i++ {
		fmt.Printf("%s\n", masukan)
	}
}

//program cetak string sebanya n kali
