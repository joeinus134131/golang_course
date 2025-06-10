package main

import (
	"fmt"
	"time"
)

// Ini adalah fungsi 'koki' kita
func koki(nama string) {
	fmt.Printf("Koki %s: Mulai memasak...\n", nama)
	time.Sleep(2 * time.Second) // Koki sedang sibuk memasak (simulasi)
	fmt.Printf("Koki %s: Selesai memasak!\n", nama)
}

func main() {
	fmt.Println("Manajer: Membuka restoran dan merekrut koki!")

	// Merekrut koki pertama (goroutine)
	go koki("Budi") // Ini akan berjalan di latar belakang

	// Merekrut koki kedua (goroutine)
	go koki("Andi") // Ini juga akan berjalan di latar belakang

	// Manajer melanjutkan pekerjaan lain tanpa menunggu koki selesai
	fmt.Println("Manajer: Mengurus meja dan pelanggan...")
	time.Sleep(3 * time.Second) // Memberi waktu koki untuk bekerja

	fmt.Println("Manajer: Restoran tutup. Semua koki selesai.")
	// Perhatikan: main goroutine harus berjalan cukup lama agar goroutine lain sempat selesai.
	// Jika main goroutine selesai terlalu cepat, goroutine lain mungkin tidak sempat tereksekusi penuh.
}
