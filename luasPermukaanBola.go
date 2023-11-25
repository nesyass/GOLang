package main

import "fmt"

func main() {
	var r, L float64

	fmt.Scan(&r)
	L = (4 * 3.14 * float64(r) * float64(r))

	fmt.Println(L)

}

//program menghitung luas permukaan bola
//masukan berupa jari jari bola dalam bilangan real
//keluaran berupa luad permukaan bola dalam bilangan real
