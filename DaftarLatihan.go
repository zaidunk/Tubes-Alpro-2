package main

import (
	f "fmt"
	s "strings"
)

func TipeLatihanMenu() {
	var num int

	length := 40
	t1 := "Jenis Latihan"
	sp1 := (length - len(t1)) / 2

	f.Println("+" + s.Repeat("-", length) + "+")
	f.Println("|" + s.Repeat(" ", sp1) + t1 + s.Repeat(" ", sp1) + "|")
	f.Println("+" + s.Repeat("-", length) + "+")

	f.Printf("1. Cardio\n2. Strength\n3. Flexibility\n\n")
	f.Print("Pilihanmu? (1/2/3): ")
	f.Scan(&num)

	switch num {
	case 1:
		TipeCardio()
	case 2:
		TipeStrength()
	case 3:
		TipeFlexibility()
	default:
		f.Println("Pilihan tidak valid.")
	}
	Counter()
}

func TipeCardio() {
	var num int
	f.Printf("1. Running\n 2. Jogging\n3. Cycling\n4. Rope Jumping\n5. Swimming\n6. Zumba\n\n")
	f.Print("Pilihanmu? (1/2/3/4/5/6): ")
	f.Scan(&num)
}

func TipeStrength() {
	var num int
	f.Printf("1. Callisthenic\n2. Weight Lifting\n")
	f.Print("Pilihanmu? (1/2): ")
	f.Scan(&num)
}

func TipeFlexibility() {
	var num int
	f.Printf("1. Static Stretching\n2. Dynamic Stretching\n3. Yoga\n4. Pilates\n5. Foam Rolling\n")
	f.Print("Pilihanmu? (1/2/3/4/5): ")
	f.Scan(&num)
}
