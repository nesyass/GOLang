package main

import "fmt"

func main() {
	var E0, c, E1, cc int

	fmt.Scan(&E0, &c, &E1)
	cc = (E0 - E1) / c

	fmt.Println(cc)

}

//didefinisikan sumber energi sebesar E0, kapasitas baterasi sebesar c, energi yang tersissa pada baterai adalah E1, dan cycle count (cc) adalah pengurangan energi baterai sebesar c
//program ini untuk menghitung cc yang diperoleh
