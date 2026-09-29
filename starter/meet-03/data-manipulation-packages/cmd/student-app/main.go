package main

import (
	"bufio"
	"fmt"
	"os"

	"data-manipulation-packages/internal/input"
	"data-manipulation-packages/internal/student"
	"data-manipulation-packages/internal/ui"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	students := student.InitialData()
	nextID := student.NextID(students)

	for {
		ui.PrintMenu()

		choice := input.ReadInt(scanner, "Pilih menu: ")
		switch choice {
		case 1:
			ui.PrintStudents(students)
		case 2:
			students = addStudent(scanner, students, nextID)
			nextID++
		case 3:
			id := input.ReadInt(scanner, "ID student: ")
			score := input.ReadInt(scanner, "Nilai baru: ")
			if student.UpdateScore(students, id, score) {
				fmt.Println("Nilai berhasil diupdate.")
			} else {
				fmt.Println("Student tidak ditemukan.")
			}
		case 4:
			before := len(students)
			students = student.RemoveInactive(students)
			fmt.Printf("%d student inactive dihapus.\n", before-len(students))
		case 5:
			major := input.ReadString(scanner, "Jurusan: ")
			ui.PrintStudents(student.FilterByMajor(students, major))
		case 6:
			id := input.ReadInt(scanner, "ID student: ")
			ui.PrintStudentByID(students, id)
		case 7:
			fmt.Printf("Rata-rata nilai: %.2f\n", student.AverageScore(students))
		case 0:
			fmt.Println("Selesai.")
			return
		default:
			fmt.Println("Menu tidak valid.")
		}
	}
}

func addStudent(scanner *bufio.Scanner, students []student.Student, id int) []student.Student {
	name := input.ReadString(scanner, "Nama: ")
	major := input.ReadString(scanner, "Jurusan: ")
	score := input.ReadInt(scanner, "Nilai: ")

	fmt.Println("Student berhasil ditambahkan.")
	return append(students, student.Student{
		ID:     id,
		Name:   name,
		Major:  major,
		Score:  score,
		Active: true,
	})
}
