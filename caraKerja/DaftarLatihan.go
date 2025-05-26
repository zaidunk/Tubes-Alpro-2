package caraKerja

import (
	f "fmt"
	s "strings"
)

type datLat struct {
	namaLatihan  string
	jenisLatihan string
}

var listLatihan = []datLat{
	{"Lari", "Cardio"},
	{"Jogging", "Cardio"},
	{"Sprint", "Cardio"},
	{"Jalan cepat", "Cardio"},
	{"Bersepeda", "Cardio"},
	{"Spinning", "Cardio"},
	{"Lompat tali", "Cardio"},
	{"Zumba", "Cardio"},
	{"Aerobik", "Cardio"},
	{"HIIT", "Cardio"},
	{"Tabata", "Cardio"},
	{"Stair Climbing", "Cardio"},
	{"Hiking", "Cardio"},
	{"Kickboxing", "Cardio"},
	{"Berenang", "Cardio"},
	{"Dancing", "Cardio"},
	{"Rowing", "Cardio"},
	{"Rollerblading", "Cardio"},
	{"Mountain Climber", "Cardio"},
	{"Burpees", "Cardio"},
	{"Jumping Jacks", "Cardio"},
	{"Jump Squats", "Cardio"},
	{"Shadow Boxing", "Cardio"},
	{"Battle Ropes", "Cardio"},
	{"Circuit Training", "Cardio"},

	{"Push-up", "Strength"},
	{"Pull-up", "Strength"},
	{"Sit-up", "Strength"},
	{"Plank", "Strength"},
	{"Squat", "Strength"},
	{"Deadlift", "Strength"},
	{"Bench Press", "Strength"},
	{"Overhead Press", "Strength"},
	{"Bent-over Row", "Strength"},
	{"Lunges", "Strength"},
	{"Bicep Curl", "Strength"},
	{"Tricep Dip", "Strength"},
	{"Russian Twist", "Strength"},
	{"Leg Press", "Strength"},
	{"Dumbbell Fly", "Strength"},
	{"Cable Row", "Strength"},
	{"Kettlebell Swing", "Strength"},
	{"Farmer's Walk", "Strength"},
	{"Chin-up", "Strength"},
	{"Barbell Clean & Press", "Strength"},
	{"Hip Thrust", "Strength"},
	{"Wall Sit", "Strength"},
	{"Shoulder Shrugs", "Strength"},
	{"Resistance Band Training", "Strength"},
	{"Calf Raises", "Strength"},

	{"Static Stretching", "Flexibility"},
	{"Dynamic Stretching", "Flexibility"},
	{"Yoga", "Flexibility"},
	{"Pilates", "Flexibility"},
	{"Foam Rolling", "Flexibility"},
	{"Cat-Cow Stretch", "Flexibility"},
	{"Cobra Stretch", "Flexibility"},
	{"Child’s Pose", "Flexibility"},
	{"Seated Forward Fold", "Flexibility"},
	{"Downward Dog", "Flexibility"},
	{"Butterfly Stretch", "Flexibility"},
	{"Hamstring Stretch", "Flexibility"},
	{"Hip Flexor Stretch", "Flexibility"},
	{"Side Bend Stretch", "Flexibility"},
	{"Shoulder Stretch", "Flexibility"},
	{"Neck Stretch", "Flexibility"},
	{"Quad Stretch", "Flexibility"},
	{"Lizard Pose", "Flexibility"},
	{"Pigeon Pose", "Flexibility"},
	{"Spinal Twist", "Flexibility"},
	{"Chest Opener", "Flexibility"},
	{"Ankle Mobility Drills", "Flexibility"},
	{"Wrist Stretching", "Flexibility"},
	{"Tai Chi", "Flexibility"},
	{"Resistance Band Stretching", "Flexibility"},
}

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
		return
	default:
		f.Println("Opsi tidak valid")
		DataLatihan()
	}

}

func LinearSearch(nama string) int {
	for i := 0; i < len(listLatihan); i++ {
		if listLatihan[i].namaLatihan == nama {
			return i
		}
	}
	return -1
}

func InsertionSortNamaLatihan() {
	var i, j int
	var key datLat
	for i = 1; i < len(listLatihan); i++ {
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
	InsertionSortNamaLatihan()
	low, high = 0, len(listLatihan)-1

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
	f.Print("Latihan apa yang ingin kamu cari? ")
	f.Scan(&latihan)

	var idx int
	if len(listLatihan) >= 75 {
		idx = BinarySearch(latihan)
	} else {
		idx = LinearSearch(latihan)
	}

	if idx != -1 {
		f.Println("Latihan ditemukan!")
		f.Println("Nama:", listLatihan[idx].namaLatihan)
		f.Println("Jenis:", listLatihan[idx].jenisLatihan)
	} else {
		f.Println("Latihan tidak ditemukan.")
	}
}
