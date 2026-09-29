package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readString(scanner *bufio.Scanner, label string) string {
	for {
		fmt.Print(label)
		if !scanner.Scan() {
			fmt.Println()
			os.Exit(0)
		}

		value := strings.TrimSpace(scanner.Text())
		if value != "" {
			return value
		}
		fmt.Println("Input tidak boleh kosong.")
	}
}

func readInt(scanner *bufio.Scanner, label string) int {
	for {
		value := readString(scanner, label)
		number, err := strconv.Atoi(value)
		if err == nil {
			return number
		}
		fmt.Println("Input harus angka.")
	}
}
