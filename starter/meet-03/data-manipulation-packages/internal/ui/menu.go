package ui

import "fmt"

func PrintMenu() {
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
