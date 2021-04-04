package main

import (
	"fmt"
)

func main() {
	fmt.Println("DAFTAR HARGA BARANG")
	fmt.Println("Sabun Cuci Piring : 50.000")
	fmt.Println("Sabun Mandi : 25.000")
	fmt.Println("Shampoo : 35.000")
	fmt.Println("Pasta Gigi : 10.000")
	fmt.Println("Tissue : 7.500")
	fmt.Println("----------------------")

	//INPUT
	var (
		k1, k2, k3, k4, k5 int
	)
	fmt.Print("Masukkan Jumlah Sabun Cuci Piring")
	fmt.Scanln(&k1)
	fmt.Print("Masukkan Jumlah Sabun Mandi")
	fmt.Scanln(&k2)
	fmt.Print("Masukkan Jumlah Shampoo")
	fmt.Scanln(&k3)
	fmt.Print("Masukkan Jumlah Pasta Gigi")
	fmt.Scanln(&k4)
	fmt.Print("Masukkan Jumlah Tissue")
	fmt.Scanln(&k5)
	fmt.Println("----------------------")

	//Struct
	type hargabarang struct {
		SabunCuciPiring, SabunMandi, Shampoo, PastaGigi, Tissue int
	}
	var h1 = hargabarang{}
	h1.SabunCuciPiring = 50000
	h1.SabunMandi = 25000
	h1.Shampoo = 35000
	h1.PastaGigi = 10000
	h1.Tissue = 7500

	type totalbarang struct {
		total1, total2, total3, total4, total5 int
	}
	var t1 = totalbarang{}
	t1.total1 = k1
	t1.total2 = k2
	t1.total3 = k3
	t1.total4 = k4
	t1.total5 = k5

	fmt.Println("(DETAIL BELANJA)")
	fmt.Println("Sabun Cuci Piring : Rp", h1.SabunCuciPiring, "X", t1.total1)
	fmt.Println("Sabun Mandi : Rp", h1.SabunMandi, "X", t1.total2)
	fmt.Println("Shampoo : Rp", h1.Shampoo, "X", t1.total3)
	fmt.Println("Pasta Gigi : Rp", h1.PastaGigi, "X", t1.total4)
	fmt.Println("Tissue : Rp", h1.Tissue, "X", t1.total5)
	fmt.Println("----------------------")

	//RUMUS TOTAL
	var (
		hargattl1, hargattl2, hargattl3, hargattl4, hargattl5, total int
	)
	fmt.Println("(INVOICE)")
	fmt.Println("Harga Sabun Cuci Piring : Rp", h1.SabunCuciPiring*k1)
	fmt.Println("Harga Sabun Mandi : Rp", h1.SabunMandi*k2)
	fmt.Println("Harga Shampoo : Rp", h1.Shampoo*k3)
	fmt.Println("Harga Pasta Gigi : Rp", h1.PastaGigi*k4)
	fmt.Println("Tissue : Rp", h1.Tissue*k5)
	fmt.Println("----------------------")

	hargattl1 = h1.SabunCuciPiring * k1
	hargattl2 = h1.SabunMandi * k2
	hargattl3 = h1.Shampoo * k3
	hargattl4 = h1.PastaGigi * k4
	hargattl5 = h1.Tissue * k5
	total = hargattl1 + hargattl2 + hargattl3 + hargattl4 + hargattl5
	fmt.Println("TOTAL YANG HARUS DIBAYARKAN : Rp", total)
}
