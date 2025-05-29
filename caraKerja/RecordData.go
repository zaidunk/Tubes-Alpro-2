package caraKerja

import (
	f "fmt"
	"math/rand"
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
	tipe    datLat
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

func Record() {
	var num int
	CetakHeader("Input Data")
	f.Println("1. Input Jadwal\n2. Daftar Jadwal\n3. Hapus Jadwal\n4. Input Latihan\n5. Ubah Latihan\n6. Hapus Latihan\n7. Kembali")
	f.Print("Pilihanmu? (1/2/3/4/5/6/7): ")
	f.Scan(&num)
	switch num {
	case 1:
		InputJadwal()
		Record()
	case 2:
		OutputJadwal()
		Record()
	case 3:
		HapusJadwal()
		Record()
	case 4:
		InputLatihan()
		Record()
	case 5:
		UbahLatihan()
		Record()
	case 6:
		HapusLatihan()
		Record()
	case 7:
		return
	default:
		f.Println("Input tidak valid")
		Record()
	}
}
func InputJadwal() {
	var i, x int
	CetakHeader("Input Data Jadwal")
	f.Print("Masukkan hari\n1. Senin\n2. Selasa\n3. Rabu\n4. Kamis\n5. Jumat\n6. Sabtu\n7. Minggu\n\n")
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

func OutputJadwal() {
	CetakHeader("Daftar Jadwal Latihan")
	var hari = [7]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	var jenis = [4]string{"-", "Cardio", "Strength", "Flexibility"}

	for i := 0; i < 7; i++ {
		f.Printf("%s: %s\n", hari[i], jenis[B[i]])
	}
	f.Println()
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
	PrintLatihan(c.lat)
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
	PrintLatihan(c.lat)

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
	var d, x, i int
	var jenis string
	var t bool
	t = false
	x = 0

	f.Print("Masukkan nama (latihan Gunakan underscore sebagai spasi): ")
	f.Scan(&l.nama)

	f.Print("Masukkan jenis latihan\n1. Cardio\n2. Strength\n3. Flexibility\n")
	f.Print("Pilihanmu? (1/2/3): ")
	f.Scan(&l.jenis)
	switch l.jenis {
	case 1:
		jenis = "Cardio"
	case 2:
		jenis = "Strength"
	case 3:
		jenis = "Flexibility"
	}

	f.Print("Masukkan tipe latihan\n")
	f.Println("Daftar latihan: ")
	for i = 0; i < n; i++ {
		if listLatihan[i].jenisLatihan == jenis {
			f.Printf("%d. %s\n", i+1, listLatihan[i].namaLatihan)
		}
	}
	f.Print("Pilihanmu? (1/2/3/...): ")

	for !t {
		f.Scan(&x)
		x = x - 1
		if listLatihan[x].jenisLatihan == jenis {
			l.tipe = listLatihan[x]
			t = true
		} else {
			f.Println("Index tidak valid")
			f.Print("Pilihanmu? (1/2/3/...): ")
		}
	}

	f.Print("Masukkan durasi latihan (dalam menit): ")
	f.Scan(&d)
	l.durasi = time.Duration(d) * time.Minute

	f.Print("Masukkan jumlah kalori latihan: ")
	f.Scan(&l.jKalori)
}

func Recommend() {
	var num int
	var kata string
	if c.lat == 0 || c.cardio == c.flexibility && c.flexibility == c.strength {
		num = rand.Intn(3)
		switch num {
		case 0:
			kata = "Cardio"
		case 1:
			kata = "Strength"
		case 2:
			kata = "Flexibility"

		}
		f.Printf("Mau Latihan apa? Kalo aku sih nyaranin latihan %s\n", kata)
	}

	if c.cardio > c.strength && c.cardio > c.flexibility {
		f.Println("Kamu terlalu banyak latihan cardio! Sebaiknya kamu coba tipe latihan lain.")
	} else if c.strength > c.cardio && c.strength > c.flexibility {
		f.Println("Kamu terlalu banyak latihan strength! Sebaiknya kamu coba tipe latihan lain.")
	} else if c.flexibility > c.cardio && c.flexibility > c.strength {
		f.Println("Kamu terlalu banyak latihan flexibility! Sebaiknya kamu coba tipe latihan lain.")
	}
}
