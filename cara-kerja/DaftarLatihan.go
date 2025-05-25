package main

import (
	f "fmt"
	s "strings"

	"github.com/zaidunk/Tubes-Alpro-2/menu"
)

func DataLatihan() {
	var num int

	length := 40
	t1 := "Jenis Latihan"
	sp1 := (length - len(t1)) / 2

	f.Println("+" + s.Repeat("-", length) + "+")
	f.Println("|" + s.Repeat(" ", sp1) + t1 + s.Repeat(" ", sp1) + "|")
	f.Println("+" + s.Repeat("-", length) + "+")

	f.Printf("1. Cari Jenis Latihan\n2. Ubah Daftar Latihan\n4. Kembali\n\n")
	f.Print("Pilihanmu? (1/2/3/4): ")
	f.Scan(&num)

	switch num {
	case 1:
		CariJenisLatihan()
	case 2:
		//UbahDaftarLatihan
	case 3:
		//HapusLatihan
	case 4:
		menu.MenuUtama()
	default:
		DataLatihan()
	}

}

func CariJenisLatihan() {

}
