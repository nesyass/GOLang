package main

import "fmt"

func main() {
	var jpa, jkl, ji, jkm, je, jp int

	fmt.Scan(&jpa, &jkl, &ji, &jkm, &je)
	jp = (jpa + jkl + ji) - (jkm + je)

	fmt.Println(jp)
}

//program ini digunakan untuk menghitung jumlah penduduk suatu negara
//jpa (jumlah penduduk awal)
//jkl(jumlah kelahiran)
//ji (jumlah imigrasi)
//jkm (jumlah kematian)
//je (jumlah emigrasi)
//jp (jumlah penduduk)
