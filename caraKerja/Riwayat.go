package caraKerja

import (
	"fmt"
	f "fmt"
)

func RiwayatLatihan() {
	var i, x, nLat, tKalori int
	CetakHeader("Riwayat Latihan")
	f.Print("1. Laporan\n2. Sorting berdasarkan durasi\n3. Sorting berdasarkan jumlah kalori\n\n")
	f.Print("Pilihanmu? (1/2/3): ")
	f.Scan(&x)
	nLat = c.lat
	switch x {
	case 1:
		if c.lat > 10 {
			nLat = 10
		}
		SelectionSortDurasiLatihan()
		f.Printf("Laporan %d latihan terakhir\n", nLat)
		PrintLatihan(nLat)
		for i = 0; i < nLat; i++ {
			tKalori = tKalori + A[i].jKalori
		}
		f.Printf("Total kalori yang dibakar selama %d latihan terakhir: %d kcal\n", nLat, tKalori)
	case 2:
		SelectionSortDurasiLatihan()
		PrintLatihan(nLat)
	case 3:
		SelectionSortJumlahKaloriLatihan()
		PrintLatihan(nLat)
	}
}

func SelectionSortTanggalLatihan() {
	var i, j, idx int
	var temp latihan
	for i = 0; i < n-1; i++ {
		idx = i
		for j = i + 1; j < n; j++ {
			if A[j].tanggal.After(A[idx].tanggal) {
				idx = j
			}
		}
		temp = A[i]
		A[i] = A[idx]
		A[idx] = temp
	}
}

func SelectionSortDurasiLatihan() {
	var i, j, idx int
	var temp latihan
	for i = 0; i < n-1; i++ {
		idx = i
		for j = i + 1; j < n; j++ {
			if A[j].durasi > A[idx].durasi {
				idx = j
			}
		}
		temp = A[i]
		A[i] = A[idx]
		A[idx] = temp
	}
}

func SelectionSortJumlahKaloriLatihan() {
	var i, j, idx int
	var temp latihan
	for i = 0; i < n-1; i++ {
		idx = i
		for j = i + 1; j < n; j++ {
			if A[j].jKalori > A[idx].jKalori {
				idx = j
			}
		}
		temp = A[i]
		A[i] = A[idx]
		A[idx] = temp
	}
}

func PrintLatihan(nLat int) {
	var i int
	f.Printf("%-5s %-15s %-15s %-15s %-15s %-8s %-15s\n", "No.", "Nama", "Jenis", "Tipe", "Durasi", "Kalori", "Tanggal")

	for i = 0; i < nLat; i++ {
		fmt.Printf("%-5d %-15s %-15s %-15s %-15s %-8d %-15s\n", i+1, A[i].nama, A[i].tipe.jenisLatihan, A[i].tipe.namaLatihan, A[i].durasi.String(), A[i].jKalori, A[i].tanggal.Format("2006-01-02"))
	}
}
