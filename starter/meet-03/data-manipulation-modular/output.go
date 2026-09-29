package main

import (
	"fmt"
	"sort"
)

func printMenu() {
	fmt.Println("\n=== Menu Student ===")
	fmt.Println("1. Lihat semua student")
	fmt.Println("2. Tambah student")
	fmt.Println("3. Update nilai")
	fmt.Println("4. Hapus student inactive")
	fmt.Println("5. Filter berdasarkan jurusan")
	fmt.Println("6. Cari student berdasarkan ID")
	fmt.Println("7. Lihat rata-rata nilai")
	fmt.Println("0. Keluar")
}

func printStudents(students []Student) {
	if len(students) == 0 {
		fmt.Println("Data student kosong.")
		return
	}

	sort.Slice(students, func(i, j int) bool {
		return students[i].Score > students[j].Score
	})

	fmt.Println("Daftar student:")
	for _, student := range students {
		status := "inactive"
		if student.Active {
			status = "active"
		}
		fmt.Printf("- ID %d | %s | %s | nilai %d | grade %s | %s\n",
			student.ID, student.Name, student.Major, student.Score, grade(student.Score), status)
	}
}

func printStudentByID(students []Student, id int) {
	studentByID := indexByID(students)
	if student, ok := studentByID[id]; ok {
		printStudents([]Student{student})
		return
	}
	fmt.Println("Student tidak ditemukan.")
}
