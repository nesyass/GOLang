package main

import "fmt"

func main() {
	var x int

	fmt.Scan(&x)
	if x < 0 || x > 999 {
		return
	}

	d1 := x / 100
	d2 := (x / 10) % 10
	d3 := x % 10

	fmt.Println(d1, d2, d3)
}

//program digunakan untuk menghitung tiga digit nilai yang terdapat pada suatu bilangan bulat positif x
//masukan berupa bilangan positif x yanf kurang atau sama dengan 999
//keluaran terdiri dari tiga digit bilangan d1, d2, dan d3 yang menyatakan digit pertama, kedua, dan ketiga dari x
