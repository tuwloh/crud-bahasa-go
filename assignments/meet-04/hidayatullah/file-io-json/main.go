package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var jsonFile = "nutrition.json"
var reader = bufio.NewReader(os.Stdin)

type Food struct {
	ID          int           `json:"id"`
	Name        string        `json:"name"`
	Category    string        `json:"category"`
	Price       float64       `json:"price"`
	IsHalal     bool          `json:"is_halal"`
	ExpiredAt   time.Time     `json:"expired_at"`
	Nutrition   NutritionInfo `json:"nutrition"`
	Ingredients []Ingredient  `json:"ingredients"`
}

type NutritionInfo struct {
	Calories     int     `json:"calories"`
	Proteins     float64 `json:"proteins"`
	Fat          float64 `json:"fat"`
	Carbohydrate float64 `json:"carbohydrate"`
}

type Ingredient struct {
	Name   string `json:"name"`
	Amount string `json:"amount"`
}

func main() {
	for {
		fmt.Println("\nxxxx- Menu Makanan -xxxx")
		fmt.Println("1. Lihat data")
		fmt.Println("2. Tambah data")
		fmt.Println("3. Ubah data")
		fmt.Println("4. Hapus data")
		fmt.Println("5. Cari data")
		fmt.Println("0. Keluar")

		menu := inputInt("Pilih menu: ")

		switch menu {
		case 1:
			showNutrition()
		case 2:
			addNutrition()
		case 3:
			editNutrition()
		case 4:
			deleteNutrition()
		case 5:
			searchNutrition()
		case 0:
			fmt.Println("Terima kasih telah menggunakan aplikasi ini!")
			return
		default:
			fmt.Println("Menu tidak valid")
		}
	}
}

func showNutrition() {
	foods := readNutrition()
	if len(foods) == 0 {
		fmt.Println("\nBelum ada data makanan dalam file JSON.")
		return
	}

	fmt.Println("\n" + strings.Repeat("=", 120))
	fmt.Printf("%-3s | %-19s | %-13s | %-10s | %-5s | %-10s | %-7s | %-7s | %-7s | %-7s | %s\n",
		"ID", "Nama Makanan", "Kategori", "Harga", "Halal", "Kadaluarsa", "Kalori", "Protein", "Lemak", "Karbo", "Bahan-bahan")
	fmt.Println(strings.Repeat("-", 120))

	for _, f := range foods {
		halalStr := "Tidak"
		if f.IsHalal {
			halalStr = "Ya"
		}
		fmt.Printf("%-3d | %-19s | %-13s | Rp%-8.0f | %-5s | %-10s | %4d kc | %5.1f g | %5.1f g | %5.1f g | %s\n",
			f.ID,
			f.Name,
			f.Category,
			f.Price,
			halalStr,
			f.ExpiredAt.Format("2006-01-02"),
			f.Nutrition.Calories,
			f.Nutrition.Proteins,
			f.Nutrition.Fat,
			f.Nutrition.Carbohydrate,
			formatIngredients(f.Ingredients),
		)
	}
	fmt.Println(strings.Repeat("=", 120))
}

func addNutrition() {
	foods := readNutrition()

	fmt.Println("\n--- Tambah Data Makanan Baru ---")
	name := input("Nama Makanan: ")
	category := input("Kategori (misal: Makanan Pokok / Lauk Pauk / Sayuran / Makanan Khas): ")
	price := inputFloat("Harga (Rp): ")
	isHalal := inputBool("Apakah Halal? (y/n atau true/false): ")
	expiredAt := inputDate("Tanggal Kadaluarsa (format YYYY-MM-DD): ")
	calories := inputInt("Kalori (kcal): ")
	proteins := inputFloat("Protein (g): ")
	fat := inputFloat("Lemak (g): ")
	carbohydrate := inputFloat("Karbohidrat (g): ")
	ingredientsRaw := input("Bahan-bahan (pisahkan koma, contoh: Daging Sapi (150g), Santan (100ml)): ")

	id := 1
	if len(foods) > 0 {
		id = foods[len(foods)-1].ID + 1
	}

	food := Food{
		ID:        id,
		Name:      name,
		Category:  category,
		Price:     price,
		IsHalal:   isHalal,
		ExpiredAt: expiredAt,
		Nutrition: NutritionInfo{
			Calories:     calories,
			Proteins:     proteins,
			Fat:          fat,
			Carbohydrate: carbohydrate,
		},
		Ingredients: parseIngredients(ingredientsRaw),
	}

	foods = append(foods, food)
	writeNutrition(foods)
	fmt.Println("\nData makanan berhasil ditambahkan!")
}

func editNutrition() {
	showNutrition()
	foods := readNutrition()
	id := inputInt("\nMasukkan ID makanan yang ingin diubah: ")

	for i, f := range foods {
		if f.ID == id {
			fmt.Printf("\n--- Mengubah Data: %s (ID: %d) ---\n", f.Name, f.ID)
			fmt.Println("(Tekan [Enter] jika ingin mempertahankan nilai lama)")

			name := inputWithDefault("Nama baru", f.Name)
			category := inputWithDefault("Kategori baru", f.Category)
			price := inputFloatWithDefault("Harga baru (Rp)", f.Price)
			isHalal := inputBoolWithDefault("Apakah Halal (y/n)", f.IsHalal)
			expiredAt := inputDateWithDefault("Tanggal Kadaluarsa baru (YYYY-MM-DD)", f.ExpiredAt)
			calories := inputIntWithDefault("Kalori baru (kcal)", f.Nutrition.Calories)
			proteins := inputFloatWithDefault("Protein baru (g)", f.Nutrition.Proteins)
			fat := inputFloatWithDefault("Lemak baru (g)", f.Nutrition.Fat)
			carbohydrate := inputFloatWithDefault("Karbohidrat baru (g)", f.Nutrition.Carbohydrate)
			ingredientsRaw := inputWithDefault("Bahan-bahan baru", formatIngredients(f.Ingredients))

			foods[i] = Food{
				ID:        id,
				Name:      name,
				Category:  category,
				Price:     price,
				IsHalal:   isHalal,
				ExpiredAt: expiredAt,
				Nutrition: NutritionInfo{
					Calories:     calories,
					Proteins:     proteins,
					Fat:          fat,
					Carbohydrate: carbohydrate,
				},
				Ingredients: parseIngredients(ingredientsRaw),
			}

			writeNutrition(foods)
			fmt.Println("\nData makanan berhasil diubah!")
			return
		}
	}
	fmt.Println("Data tidak ditemukan")
}

func deleteNutrition() {
	showNutrition()
	foods := readNutrition()
	id := inputInt("\nMasukkan ID makanan yang ingin dihapus: ")

	for i, f := range foods {
		if f.ID == id {
			foods = append(foods[:i], foods[i+1:]...)
			writeNutrition(foods)
			fmt.Println("\nData makanan berhasil dihapus!")
			return
		}
	}
	fmt.Println("Data tidak ditemukan")
}

func searchNutrition() {
	foods := readNutrition()
	keyword := input("\nCari berdasarkan Nama / Kategori / Bahan Baku: ")
	keyword = strings.ToLower(strings.TrimSpace(keyword))

	if keyword == "" {
		fmt.Println("Kata kunci pencarian kosong.")
		return
	}

	var results []Food
	for _, f := range foods {
		name := strings.ToLower(f.Name)
		category := strings.ToLower(f.Category)
		ingredients := strings.ToLower(formatIngredients(f.Ingredients))

		if strings.Contains(name, keyword) ||
			strings.Contains(category, keyword) ||
			strings.Contains(ingredients, keyword) {
			results = append(results, f)
		}
	}

	if len(results) == 0 {
		fmt.Printf("Tidak ditemukan makanan dengan kata kunci '%s'.\n", keyword)
		return
	}

	fmt.Printf("\nHasil pencarian untuk '%s' (%d ditemukan):\n", keyword, len(results))
	fmt.Println(strings.Repeat("=", 120))
	fmt.Printf("%-3s | %-19s | %-13s | %-10s | %-5s | %-10s | %-7s | %-7s | %-7s | %-7s | %s\n",
		"ID", "Nama Makanan", "Kategori", "Harga", "Halal", "Kadaluarsa", "Kalori", "Protein", "Lemak", "Karbo", "Bahan-bahan")
	fmt.Println(strings.Repeat("-", 120))

	for _, f := range results {
		halalStr := "Tidak"
		if f.IsHalal {
			halalStr = "Ya"
		}
		fmt.Printf("%-3d | %-19s | %-13s | Rp%-8.0f | %-5s | %-10s | %4d kc | %5.1f g | %5.1f g | %5.1f g | %s\n",
			f.ID,
			f.Name,
			f.Category,
			f.Price,
			halalStr,
			f.ExpiredAt.Format("2006-01-02"),
			f.Nutrition.Calories,
			f.Nutrition.Proteins,
			f.Nutrition.Fat,
			f.Nutrition.Carbohydrate,
			formatIngredients(f.Ingredients),
		)
	}
	fmt.Println(strings.Repeat("=", 120))
}

func readNutrition() []Food {
	file, err := os.Open(jsonFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Food{}
		}
		log.Fatal(err)
	}
	defer file.Close()

	var foods []Food
	if err := json.NewDecoder(file).Decode(&foods); err != nil {
		if err.Error() == "EOF" {
			return []Food{}
		}
		log.Fatal(err)
	}
	return foods
}

func writeNutrition(foods []Food) {
	file, err := os.Create(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(foods); err != nil {
		log.Fatal(err)
	}
}

func formatIngredients(ingredients []Ingredient) string {
	var labels []string
	for _, ing := range ingredients {
		if ing.Amount != "" && ing.Amount != "-" {
			labels = append(labels, fmt.Sprintf("%s (%s)", ing.Name, ing.Amount))
		} else {
			labels = append(labels, ing.Name)
		}
	}
	return strings.Join(labels, ", ")
}

func parseIngredients(raw string) []Ingredient {
	var ingredients []Ingredient
	if strings.TrimSpace(raw) == "" {
		return ingredients
	}

	items := strings.Split(raw, ",")
	for _, item := range items {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}

		if idxOpen := strings.Index(item, "("); idxOpen != -1 {
			if idxClose := strings.Index(item, ")"); idxClose > idxOpen {
				name := strings.TrimSpace(item[:idxOpen])
				amount := strings.TrimSpace(item[idxOpen+1 : idxClose])
				ingredients = append(ingredients, Ingredient{Name: name, Amount: amount})
				continue
			}
		}

		if parts := strings.SplitN(item, ":", 2); len(parts) == 2 {
			ingredients = append(ingredients, Ingredient{
				Name:   strings.TrimSpace(parts[0]),
				Amount: strings.TrimSpace(parts[1]),
			})
			continue
		}

		ingredients = append(ingredients, Ingredient{Name: item, Amount: "-"})
	}
	return ingredients
}

func input(label string) string {
	fmt.Print(label)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func inputInt(label string) int {
	val, _ := strconv.Atoi(input(label))
	return val
}

func inputFloat(label string) float64 {
	val, _ := strconv.ParseFloat(input(label), 64)
	return val
}

func inputBool(label string) bool {
	val := strings.ToLower(input(label))
	return val == "y" || val == "ya" || val == "yes" || val == "true" || val == "1"
}

func inputDate(label string) time.Time {
	val := input(label)
	t, err := time.Parse("2006-01-02", val)
	if err != nil {
		return time.Now().AddDate(0, 1, 0).Truncate(24 * time.Hour)
	}
	return t
}

func inputWithDefault(label, defaultVal string) string {
	text := input(fmt.Sprintf("%s [%s]: ", label, defaultVal))
	if text == "" {
		return defaultVal
	}
	return text
}

func inputIntWithDefault(label string, defaultVal int) int {
	text := input(fmt.Sprintf("%s [%d]: ", label, defaultVal))
	if text == "" {
		return defaultVal
	}
	val, err := strconv.Atoi(text)
	if err != nil {
		return defaultVal
	}
	return val
}

func inputFloatWithDefault(label string, defaultVal float64) float64 {
	text := input(fmt.Sprintf("%s [%.2f]: ", label, defaultVal))
	if text == "" {
		return defaultVal
	}
	val, err := strconv.ParseFloat(text, 64)
	if err != nil {
		return defaultVal
	}
	return val
}

func inputBoolWithDefault(label string, defaultVal bool) bool {
	strDefault := "n"
	if defaultVal {
		strDefault = "y"
	}
	text := strings.ToLower(input(fmt.Sprintf("%s [%s]: ", label, strDefault)))
	if text == "" {
		return defaultVal
	}
	return text == "y" || text == "ya" || text == "yes" || text == "true" || text == "1"
}

func inputDateWithDefault(label string, defaultVal time.Time) time.Time {
	strDefault := defaultVal.Format("2006-01-02")
	text := input(fmt.Sprintf("%s [%s]: ", label, strDefault))
	if text == "" {
		return defaultVal
	}
	t, err := time.Parse("2006-01-02", text)
	if err != nil {
		return defaultVal
	}
	return t
}
