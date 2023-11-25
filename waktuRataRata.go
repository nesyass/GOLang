package main

import "fmt"

func main() {
	var baris, waktu, totalWaktu, i int
	var rataRata float64

	fmt.Scan(&baris)
	for i = 1; i <= baris; i++ {
		fmt.Scan(&waktu)
		totalWaktu = totalWaktu + waktu
	}
	rataRata = float64(totalWaktu) / float64(baris)
	fmt.Printf("%f", rataRata)
}

//algoritma untuk menghitung rata-rata waktu yang digunakan seorang mahasiswa berlatih ngoding setiap hati
//masukan terdiri dari bebrrapa baris. baris pertama adalah banyaknya hari mahasiswa berlatih, baris berikutnya adlah bilangan bulat yang menyatakan banyaknya jam perhari yang dihabiskan untuk berlatih ngoding
//keluaran berupa rata-rata yang dihabiskan untuk latihan ngoding
