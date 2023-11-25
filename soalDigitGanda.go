package main

import "fmt"

func main() {
	var x, hasil string
	fmt.Scan(&x)
	hasil = string(x[0]) + string(x[0]) + string(x[1]) + string(x[1])
	fmt.Println(hasil)
}

//program menggandakan digit suatu bilangan, misal 15 menjadi 1155
