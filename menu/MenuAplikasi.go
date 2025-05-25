package menu

import (
	f "fmt"
	s "strings"
	"github.com/zaidunk/Tubes-Alpro-2/menu"
)

var acc, pass string

// Tampilan menu
func MenuLogin() {
	var num int
	f.Println("+-------------------------------------------------+")
	f.Println("|       Manajemen dan Pencatatan WorkOut          |")
	f.Println("+-------------------------------------------------+")
	f.Printf("1. Masuk\n2. Daftar\n3. Exit\n\n")
	f.Print("Pilihanmu? (1/2/3): ")

	// Login
	f.Scan(&num)
	if num == 1 {
		Login()
	} else if num == 2 {
		Daftar()
	} else if num == 3 {
		Exit()
	}

}

// Tampilan menu saat login
func Login() {
	var accIn, passIn string
	var count int

	f.Printf("\n+-------------------+\n")
	f.Println("       Masuk")
	f.Println("+-------------------+")

	// Proses Login
	f.Println("Masukkan Username-mu!")
	f.Scan(&accIn)

	for count < 2 {
		f.Println("Masukkan Passwordnya")
		f.Scan(&passIn)
		if accIn == acc && passIn == pass {
			f.Printf("Login berhasil!\n")
			break
		}
		count++
		f.Printf("Username atau password tidak sesuai!\n\n")
		f.Println("Masukkan Username-mu!")
		f.Scan(&accIn)
	}

	// Kondisi ketika gagal login lebih dari 4 kali
	if count >= 2 {
		f.Printf("Sepertinya kamu lupa atau memang tidak punya akun deh!\n\n")
		MenuLogin()
	}

	MenuUtama()

}

// Fitur pembuatan akun
func Daftar() {
	f.Println("\n\n+------------------+")
	f.Println("|      Daftar      |")
	f.Println("+------------------+")

	f.Print("Username: ")
	f.Scan(&acc)
	f.Print("Password: ")
	f.Scan(&pass)

	f.Printf("\nPembuatan akun berhasil!\n\n")
	MenuLogin()
}

// Tampilan menu utama
func MenuUtama() {
	var num int
	f.Printf("-------------------------------------------\n")
	f.Printf("\nHalo %s, sudah latihan apa saja hari ini?\n", acc)
	for num != 4 {
		f.Printf("-------------------------------------------\n")
		f.Printf("1. Jadwal\n2. Catatan data latihan\n3. Riwayat latihan\n4. Keluar\n\n")
		f.Print("Pilihanmu: (1/2/3/4): ")
		f.Scan(&num)

		switch num {
		case 1:
			Jadwal()
		case 2:
			Record()
		case 3:
			History()

		}
	}

	Exit()
}

// Pesan keluar dari program
func Exit() {
	length := 40
	t1 := "See you next time!"
	sp1 := (length - len(t1)) / 2
	t2 := "Manajemen dan Pencatatan WorkOut"
	sp2 := (length - len(t2)) / 2
	t3 := "Dibuat oleh Fagian & Zaidan "
	sp3 := (length - len(t3)) / 2

	f.Println()
	f.Println("+" + s.Repeat("-", length) + "+")
	f.Println("|" + s.Repeat(" ", sp1) + t1 + s.Repeat(" ", sp1) + "|")
	f.Println("|" + s.Repeat(" ", sp2) + t2 + s.Repeat(" ", sp2) + "|")
	f.Println("|" + s.Repeat(" ", sp3) + t3 + s.Repeat(" ", sp3) + "|")
	f.Println("+" + s.Repeat("-", length) + "+")

}
