package caraKerja

import (
	f "fmt"
	s "strings"
)

type datLat struct {
	namaLatihan  string
	jenisLatihan string
}

var listLatihan = [NMAX]datLat{
	{"Lari", "Cardio"},
	{"Jogging", "Cardio"},
	{"Sprint", "Cardio"},
	{"Jalan_cepat", "Cardio"},
	{"Bersepeda", "Cardio"},
	{"Spinning", "Cardio"},
	{"Lompat_tali", "Cardio"},
	{"Zumba", "Cardio"},
	{"Aerobik", "Cardio"},
	{"HIIT", "Cardio"},
	{"Tabata", "Cardio"},
	{"Stair_Climbing", "Cardio"},
	{"Hiking", "Cardio"},
	{"Kickboxing", "Cardio"},
	{"Berenang", "Cardio"},
	{"Dancing", "Cardio"},
	{"Rowing", "Cardio"},
	{"Rollerblading", "Cardio"},
	{"Mountain_Climber", "Cardio"},
	{"Burpees", "Cardio"},
	{"Jumping_Jacks", "Cardio"},
	{"Jump_Squats", "Cardio"},
	{"Shadow_Boxing", "Cardio"},
	{"Battle_Ropes", "Cardio"},
	{"Circuit_Training", "Cardio"},

	{"Push-up", "Strength"},
	{"Pull-up", "Strength"},
	{"Sit-up", "Strength"},
	{"Plank", "Strength"},
	{"Squat", "Strength"},
	{"Deadlift", "Strength"},
	{"Bench_Press", "Strength"},
	{"Overhead_Press", "Strength"},
	{"Bent-over_Row", "Strength"},
	{"Lunges", "Strength"},
	{"Bicep_Curl", "Strength"},
	{"Tricep_Dip", "Strength"},
	{"Russian_Twist", "Strength"},
	{"Leg_Press", "Strength"},
	{"Dumbbell_Fly", "Strength"},
	{"Cable_Row", "Strength"},
	{"Kettlebell_Swing", "Strength"},
	{"Farmer's_Walk", "Strength"},
	{"Chin-up", "Strength"},
	{"Barbell_Clean_&_Press", "Strength"},
	{"Hip_Thrust", "Strength"},
	{"Wall_Sit", "Strength"},
	{"Shoulder_Shrugs", "Strength"},
	{"Resistance_Band_Training", "Strength"},
	{"Calf_Raises", "Strength"},

	{"Static_Stretching", "Flexibility"},
	{"Dynamic_Stretching", "Flexibility"},
	{"Yoga", "Flexibility"},
	{"Pilates", "Flexibility"},
	{"Foam_Rolling", "Flexibility"},
	{"Cat-Cow_Stretch", "Flexibility"},
	{"Cobra_Stretch", "Flexibility"},
	{"Child’s_Pose", "Flexibility"},
	{"Seated_Forward_Fold", "Flexibility"},
	{"Downward_Dog", "Flexibility"},
	{"Butterfly_Stretch", "Flexibility"},
	{"Hamstring_Stretch", "Flexibility"},
	{"Hip_Flexor_Stretch", "Flexibility"},
	{"Side_Bend_Stretch", "Flexibility"},
	{"Shoulder_Stretch", "Flexibility"},
	{"Neck_Stretch", "Flexibility"},
	{"Quad_Stretch", "Flexibility"},
	{"Lizard_Pose", "Flexibility"},
	{"Pigeon_Pose", "Flexibility"},
	{"Spinal_Twist", "Flexibility"},
	{"Chest_Opener", "Flexibility"},
	{"Ankle_Mobility_Drills", "Flexibility"},
	{"Wrist_Stretching", "Flexibility"},
	{"Tai_Chi", "Flexibility"},
	{"Resistance_Band_Stretching", "Flexibility"},
}

// angka yang digunakan untuk menampilkan banyak data
var n int = 75

func DataDaftarLatihan() {
	InsertionSortNamaLatihan()
	var num int

	length := 40
	t1 := "Menu Daftar Latihan "
	sp1 := (length - len(t1)) / 2

	f.Println("+" + s.Repeat("-", length) + "+")
	f.Println("|" + s.Repeat(" ", sp1) + t1 + s.Repeat(" ", sp1) + "|")
	f.Println("+" + s.Repeat("-", length) + "+")

	f.Printf("1. Daftar Latihan\n2. Cari Jenis Latihan\n3. Ubah Daftar Latihan\n4. Hapus Daftar Latihan\n5. Tambah Daftar Latihan\n6.Kembali\n\n")
	f.Print("Pilihanmu? (1/2/3/4/5/6): ")
	f.Scan(&num)

	switch num {
	case 1:
		OutputJenisLatihan()
		f.Printf("\n\n")
		DataDaftarLatihan()
	case 2:
		CariJenisLatihan()
		f.Printf("\n\n")
		DataDaftarLatihan()
	case 3:
		UbahDaftarLatihan()
		f.Printf("\n\n")
		DataDaftarLatihan()
	case 4:
		HapusDaftarLatihan()
		f.Printf("\n\n")
		DataDaftarLatihan()
	case 5:
		TambahDaftarLatihan()
		f.Printf("\n\n")
		DataDaftarLatihan()
	case 6:
		return
	default:
		f.Println("Opsi tidak valid")
		DataDaftarLatihan()
	}
}

func OutputJenisLatihan() {
	var i int
	f.Println("Daftar latihan: ")
	for i = 0; i < n; i++ {
		f.Printf("%d. %s %s\n", i+1, listLatihan[i].namaLatihan, listLatihan[i].jenisLatihan)
	}
}
func LinearSearchNamaLatihan(nama string) int {
	for i := 0; i < n; i++ {
		if listLatihan[i].namaLatihan == nama || s.ToLower(listLatihan[i].namaLatihan) == nama || s.ToUpper(listLatihan[i].namaLatihan) == nama {
			return i
		}
	}
	return -1
}

func InsertionSortNamaLatihan() {
	var i, j int
	var key datLat
	for i = 1; i < n; i++ {
		key = listLatihan[i]
		j = i - 1
		for j >= 0 && listLatihan[j].namaLatihan > key.namaLatihan {
			listLatihan[j+1] = listLatihan[j]
			j--
		}
		listLatihan[j+1] = key
	}
}

func BinarySearch(nama string) int {
	var low, high, mid int
	var target string
	low, high = 0, n-1

	for low <= high {
		mid = (low + high) / 2
		target = listLatihan[mid].namaLatihan

		if target == nama {
			return mid
		} else if target < nama {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func CariJenisLatihan() {
	var latihan string
	f.Println("Latihan apa yang ingin kamu cari? (Gunakan underscore sebagai spasi) ")
	f.Scan(&latihan)

	var idx int
	if n >= 75 {
		idx = BinarySearch(latihan)
	} else {
		idx = LinearSearchNamaLatihan(latihan)
	}

	if idx != -1 {
		f.Println("Latihan ditemukan!")
		f.Println("Nama:", listLatihan[idx].namaLatihan)
		f.Println("Jenis:", listLatihan[idx].jenisLatihan)
	} else {
		f.Println("Latihan tidak ditemukan.")
	}
}

func UbahDaftarLatihan() {
	var i, idx int
	f.Println("Daftar latihan: ")
	for i = 0; i < n; i++ {
		f.Printf("%d. %s %s\n", i+1, listLatihan[i].namaLatihan, listLatihan[i].jenisLatihan)
	}

	f.Printf("\n\nNo. Latihan apa yang ingin kamu ubah? ")
	f.Scan(&idx)
	if idx < 1 || idx > n {
		f.Println("Input tidak valid")
		return
	}
	idx = idx - 1

	f.Printf("Data yang ingin diubah: %s %s\n", listLatihan[idx].namaLatihan, listLatihan[idx].jenisLatihan)
	f.Println("Masukkan nama latihan dan jenisnya:")
	f.Scan(&listLatihan[idx].namaLatihan, &listLatihan[idx].jenisLatihan)
	f.Println("Data berhasil diubah!")
}

func HapusDaftarLatihan() {
	var i, idx int
	f.Println("Daftar latihan: ")
	for i = 0; i < n; i++ {
		f.Printf("%d. %s %s\n", i+1, listLatihan[i].namaLatihan, listLatihan[i].jenisLatihan)
	}

	f.Print("\n\nNo. Latihan yang ingin kamu hapus? ")
	f.Scan(&idx)
	if idx < 1 || idx > n {
		f.Println("Input tidak valid")
		return
	}
	idx = idx - 1

	for i = idx; i < n-1; i++ {
		listLatihan[i] = listLatihan[i+1]
	}
	n--
	f.Println("Latihan berhasil dihapus!")
}

func TambahDaftarLatihan() {
	var nama string
	var jenis string

	if n >= NMAX {
		f.Println("Daftar latihan sudah penuh!")
		return
	}

	f.Println("Masukkan nama latihan baru: (Gunakan underscore sebagai spasi) ")
	f.Scan(&nama)

	f.Print("Masukkan jenis latihan (Cardio/Strength/Flexibility): ")
	f.Scan(&jenis)

	listLatihan[n].namaLatihan = nama
	listLatihan[n].jenisLatihan = jenis
	n++

	f.Println("Latihan berhasil ditambahkan.")
}
