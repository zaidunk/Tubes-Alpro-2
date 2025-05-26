package caraKerja

import (
	f "fmt"
	s "strings"
	"time"
)

const NMAX int = 999

type latihan struct {
	nama    string
	jenis   int
	tanggal time.Time
	durasi  time.Duration
	jKalori int
}

type counter struct {
	lat         int
	cardio      int
	strength    int
	flexibility int
}

type tabLatihan [NMAX]latihan
type tabJadwal [7]int

var A tabLatihan
var B tabJadwal
var c counter

func InputJadwal() {
	var i, x int
	CetakHeader("Input Data Jadwal")
	f.Print("Masukkan hari\n1. Senin\n2. Selasa\n3. Rabu\n4. Kamis\n 5. Jumat\n 6. Sabtu\n 7. Minggu\n\n")
	f.Print("Pilihanmu? (1/2/3/4/5/6/7): ")
	f.Scan(&x)

	if x < 1 || x > 7 {
		f.Println("Index tidak valid.")
		return
	}

	i = x - 1

	f.Print("Masukkan jenis latihan\n1. Cardio\n2. Strength\n3. Flexibility\n")
	f.Print("Pilihanmu? (1/2/3): ")
	f.Scan(&B[i])
}

func HapusJadwal() {
	var i, x int
	CetakHeader("Hapus Data Jadwal")
	f.Print("Masukkan hari\n1. Senin\n2. Selasa\n3. Rabu\n4. Kamis\n 5. Jumat\n 6. Sabtu\n 7. Minggu\n\n")
	f.Print("Pilihanmu? (1/2/3/4/5/6/7): ")
	f.Scan(&x)

	if x < 1 || x > 7 {
		f.Println("Index tidak valid.")
		return
	}

	i = x - 1

	B[i] = 0
}

func InputLatihan() {
	CetakHeader("Input Data Latihan")

	InputSatuLatihan(&A[c.lat])
	A[c.lat].tanggal = time.Now()

	// Update counter
	switch A[c.lat].jenis {
	case 1:
		c.cardio++
	case 2:
		c.strength++
	case 3:
		c.flexibility++
	}
	c.lat++
}

func UbahLatihan() {
	var x, i int
	CetakHeader("Ubah Data Latihan")

	//Menampilkan riwayat latihan
	//Riwayat()
	f.Print("Masukkan nomor latihan yang ingin diubah: ")
	f.Scan(&x)

	if x < 1 || x > c.lat {
		f.Println("Index tidak valid.")
		return
	}

	i = x - 1

	switch A[i].jenis {
	case 1:
		c.cardio--
	case 2:
		c.strength--
	case 3:
		c.flexibility--
	}

	InputSatuLatihan(&A[i])

	switch A[i].jenis {
	case 1:
		c.cardio++
	case 2:
		c.strength++
	case 3:
		c.flexibility++
	}
}

func HapusLatihan() {
	var x, i int
	CetakHeader("Hapus Data Latihan")
	//Menampilkan riwayat latihan
	//Riwayat()

	//Input index latihan yg ingin dihapus
	f.Scan(&x)
	if x < 1 || x > c.lat {
		f.Println("Index tidak valid.")
		return
	}

	for i = x - 1; i < c.lat-1; i++ {
		A[i] = A[i+1]
	}

	//Update counter
	switch A[x-1].jenis {
	case 1:
		c.cardio--
	case 2:
		c.strength--
	case 3:
		c.flexibility--
	}
	c.lat--
}

func CetakHeader(text string) {
	length := 40
	sp := (length - len(text)) / 2
	f.Println("+" + s.Repeat("-", length) + "+")
	f.Println("|" + s.Repeat(" ", sp) + text + s.Repeat(" ", length-len(text)-sp) + "|")
	f.Println("+" + s.Repeat("-", length) + "+")
}

func InputSatuLatihan(l *latihan) {
	var d int

	f.Print("Masukkan nama latihan: ")
	f.Scan(&l.nama)

	f.Print("Masukkan jenis latihan\n1. Cardio\n2. Strength\n3. Flexibility\n")
	f.Print("Pilihanmu? (1/2/3): ")
	f.Scan(&l.jenis)

	f.Print("Masukkan durasi latihan (dalam menit): ")
	f.Scan(&d)
	l.durasi = time.Duration(d) * time.Minute

	f.Print("Masukkan jumlah kalori latihan: ")
	f.Scan(&l.jKalori)
}
