package main

import (
	f "fmt"
	s "strings"
	"time"
)

const NMAX int = 100

type latihan struct {
	nama string
	//jenis	daftarlatihan?
	tanggal time.Time
	durasi  time.Duration
	jKalori int
}

type tabLatihan [NMAX]latihan

func main() {
	var A tabLatihan
	var counter int
	RecordProcess(&A, counter)
}

func RecordMenu() {
	length := 40
	t1 := "Workout"
	sp1 := (length - len(t1)) / 2

	f.Println("+" + s.Repeat("-", length) + "+")
	f.Println("|" + s.Repeat(" ", sp1) + t1 + s.Repeat(" ", sp1) + "|")
	f.Println("+" + s.Repeat("-", length) + "+")

	f.Printf("1. Latihan\n2. Durasi\n3. Jumlah Kalori\n4. Submit\n\n")
}

func RecordDurasi(A *tabLatihan, counter int) {
	var x string
	f.Print("Masukkan durasi latihan (contoh: 1h30m17s): ")
	f.Scan(&x)

	durasi, err := time.ParseDuration(x)
	if err != nil {
		f.Println("Format durasi tidak valid")
	} else {
		A[counter].durasi = durasi
	}

}

func RecordKalori(A *tabLatihan, counter int) {
	f.Print("Masukkan jumlah kalori: ")
	f.Scan(&A[counter].jKalori)
}

func RecordSubmit() bool {
	var num int
	f.Println("Apakah kamu yakin ingin menyelesaikan workout ini?")
	f.Printf("1. Yes\n2. No\n\n")
	f.Print("Pilihanmu? (1/2): ")
	f.Scan(&num)

	return num == 1
}

func RecordProcess(A *tabLatihan, counter int) {
	var num int
	A[counter].tanggal = time.Now()

	for {
		RecordMenu()
		f.Print("Pilihanmu? (1/2/3/4): ")
		f.Scan(&num)

		switch num {
		case 1:
			//DaftarLatihan()
		case 2:
			RecordDurasi(A, counter)
		case 3:
			RecordKalori(A, counter)
		case 4:
			f.Print("Nama workout: ")
			f.Scan(&A[counter].nama)
			if RecordSubmit() {
				f.Println("Workout berhasil disubmit.")
				//f.Print(A[counter].tanggal)
				return
			}
		default:
			f.Println("Pilihan tidak valid.")
		}
	}
}
