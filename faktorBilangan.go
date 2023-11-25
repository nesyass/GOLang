package main

import "fmt"

func main() {

	var x, y int
	var hasil bool

	fmt.Scan(&x, &y)
	hasil = (y%x == 0)

	fmt.Println(hasil)

}

//program menentukan faktor bilangan
//bilangan x adalah faktor dari bilangan y apabila x habis membagi y
