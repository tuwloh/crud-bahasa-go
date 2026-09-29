package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	students := initialStudents()
	nextID := nextStudentID(students)

	for {
		printMenu()

		choice := readInt(scanner, "Pilih menu: ")
		switch choice {
		case 1:
			printStudents(students)
		case 2:
			students = addStudent(scanner, students, nextID)
			nextID++
		case 3:
			id := readInt(scanner, "ID student: ")
			score := readInt(scanner, "Nilai baru: ")
			if updateScore(students, id, score) {
				fmt.Println("Nilai berhasil diupdate.")
			} else {
				fmt.Println("Student tidak ditemukan.")
			}
		case 4:
			before := len(students)
			students = removeInactive(students)
			fmt.Printf("%d student inactive dihapus.\n", before-len(students))
		case 5:
			major := readString(scanner, "Jurusan: ")
			printStudents(filterByMajor(students, major))
		case 6:
			id := readInt(scanner, "ID student: ")
			printStudentByID(students, id)
		case 7:
			fmt.Printf("Rata-rata nilai: %.2f\n", averageScore(students))
		case 0:
			fmt.Println("Selesai.")
			return
		default:
			fmt.Println("Menu tidak valid.")
		}
	}
}

func addStudent(scanner *bufio.Scanner, students []Student, id int) []Student {
	name := readString(scanner, "Nama: ")
	major := readString(scanner, "Jurusan: ")
	score := readInt(scanner, "Nilai: ")

	students = append(students, Student{
		ID:     id,
		Name:   name,
		Major:  major,
		Score:  score,
		Active: true,
	})
	fmt.Println("Student berhasil ditambahkan.")
	return students
}
