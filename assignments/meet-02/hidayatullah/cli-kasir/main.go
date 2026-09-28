package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const minimumDiskon = 100000

type Item struct {
	Nama     string
	Harga    int
	Jumlah   int
	Stok     int
	Subtotal int
}

var katalog []Item

func hitungSubtotal(harga int, jumlah int) int {
	return harga * jumlah
}

func hitungDiskon(total int) int {
	if total >= minimumDiskon {
		return total * 10 / 100
	}

	return 0
}

func hitungTotal(items []Item) int {
	total := 0
	for _, item := range items {
		total += item.Subtotal
	}

	return total
}

func bacaTeks(scanner *bufio.Scanner, label string) (string, error) {
	fmt.Print(label)
	if !scanner.Scan() {
		return "", fmt.Errorf("input teks tidak valid")
	}

	value := strings.TrimSpace(scanner.Text())
	if value == "" {
		return "", fmt.Errorf("input tidak boleh kosong")
	}

	return value, nil
}

func bacaAngkaPositif(scanner *bufio.Scanner, label string) (int, error) {
	fmt.Print(label)
	if !scanner.Scan() {
		return 0, fmt.Errorf("input angka tidak valid")
	}

	value, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		return 0, fmt.Errorf("input angka tidak valid")
	}
	if value <= 0 {
		return 0, fmt.Errorf("input harus lebih dari 0")
	}

	return value, nil
}

func tampilkanKeranjang(items []Item) {
	if len(items) == 0 {
		fmt.Println("Keranjang masih kosong")
		return
	}

	fmt.Println("\n=== Keranjang ===")
	for _, item := range items {
		fmt.Printf("%s x%d @Rp%d = Rp%d\n", item.Nama, item.Jumlah, item.Harga, item.Subtotal)
	}
}

func cetakStruk(items []Item) {
	if len(items) == 0 {
		fmt.Println("Keranjang masih kosong")
		return
	}

	total := hitungTotal(items)
	diskon := hitungDiskon(total)
	totalBayar := total - diskon

	tampilkanKeranjang(items)
	fmt.Println("---------------------")
	fmt.Println("Total:", total)
	fmt.Println("Diskon:", diskon)
	fmt.Println("Total Bayar:", totalBayar)
}

func tampilkanMenu() {
	fmt.Println("\n=== CLI Kasir Sederhana ===")
	fmt.Println("1. Tambah barang")
	fmt.Println("2. Lihat katalog")
	fmt.Println("3. Tambah ke keranjang")
	fmt.Println("4. Lihat keranjang")
	fmt.Println("5. Checkout")
	fmt.Println("6. Reset keranjang")
	fmt.Println("7. Keluar")
}

func tambahBarang(scanner *bufio.Scanner) {
	fmt.Println("\n=== Tambah Barang ke Katalog ===")

	nama, err := bacaTeks(scanner, "Nama barang: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	harga, err := bacaAngkaPositif(scanner, "Harga barang: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	stok, err := bacaAngkaPositif(scanner, "Stok barang: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	item := Item{
		Nama:  nama,
		Harga: harga,
		Stok:  stok,
	}

	katalog = append(katalog, item)

	fmt.Println("Barang berhasil ditambahkan ke katalog")
}

func tampilkanKatalog() {
	if len(katalog) == 0 {
		fmt.Println("Katalog masih kosong")
		return
	}

	fmt.Println("\n=== Katalog Barang ===")

	for i, item := range katalog {
		fmt.Printf("%d. %s - Rp%d - Stok: %d\n",
			i+1,
			item.Nama,
			item.Harga,
			item.Stok,
		)
	}
}

func tambahKeKeranjang(scanner *bufio.Scanner, items *[]Item) {
	if len(katalog) == 0 {
		fmt.Println("Katalog masih kosong")
		return
	}

	tampilkanKatalog()

	nomor, err := bacaAngkaPositif(scanner, "Pilih nomor barang: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if nomor > len(katalog) {
		fmt.Println("Nomor barang tidak tersedia")
		return
	}

	index := nomor - 1
	item := &katalog[index]

	if item.Stok == 0 {
		fmt.Println("Stok Habis")
		return
	}

	jumlah, err := bacaAngkaPositif(scanner, "Jumlah barang: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if jumlah > item.Stok {
		fmt.Println("Stok tidak mencukupi")
		return
	}

	item.Stok -= jumlah

	itemKeranjang := Item{
		Nama:     item.Nama,
		Harga:    item.Harga,
		Jumlah:   jumlah,
		Subtotal: hitungSubtotal(item.Harga, jumlah),
	}

	*items = append(*items, itemKeranjang)

	fmt.Println("Barang berhasil ditambahkan ke keranjang")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	items := []Item{}

	for {
		tampilkanMenu()

		pilihan, err := bacaAngkaPositif(scanner, "Pilih menu: ")
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		switch pilihan {
		case 1:
			tambahBarang(scanner)

		case 2:
			tampilkanKatalog()

		case 3:
			tambahKeKeranjang(scanner, &items)

		case 4:
			tampilkanKeranjang(items)

		case 5:
			cetakStruk(items)

		case 6:
			items = []Item{}
			fmt.Println("Keranjang berhasil direset")

		case 7:
			fmt.Println("Terima kasih")
			return

		default:
			fmt.Println("Error: menu tidak tersedia")
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
}
