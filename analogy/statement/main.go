package main

import (
	"fmt"
	"time"
)

// Jalur Makanan Panas
func jalurMakananPanas(ch chan string) {
	time.Sleep(1 * time.Second) // Makanan siap setelah 1 detik
	ch <- "Steak Panas"
}

// Jalur Minuman Dingin
func jalurMinumanDingin(ch chan string) {
	time.Sleep(500 * time.Millisecond) // Minuman siap lebih cepat (0.5 detik)
	ch <- "Es Teh Lemon"
}

func main() {
	// Membuat channel untuk masing-masing jalur
	makananPanas := make(chan string)
	minumanDingin := make(chan string)

	fmt.Println("Pengantar: Siap menunggu pesanan dari jalur manapun.")

	// Memulai goroutine untuk setiap jalur
	go jalurMakananPanas(makananPanas)
	go jalurMinumanDingin(minumanDingin)

	// Pengantar menunggu dan memilih pesanan yang pertama siap
	for i := 0; i < 2; i++ { // Akan mengambil 2 pesanan
		select {
		case pesananMakanan := <-makananPanas:
			fmt.Printf("Pengantar: Menerima pesanan: %s (dari Jalur Makanan Panas)\n", pesananMakanan)
		case pesananMinuman := <-minumanDingin:
			fmt.Printf("Pengantar: Menerima pesanan: %s (dari Jalur Minuman Dingin)\n", pesananMinuman)
		default: // Opsional: jika tidak ada channel yang siap segera
			// fmt.Println("Pengantar: Belum ada pesanan yang siap, masih menunggu...")
			// time.Sleep(100 * time.Millisecond) // Bisa ditambahkan jeda
		}
	}
	fmt.Println("Pengantar: Semua pesanan diambil. Selesai bekerja.")
}
