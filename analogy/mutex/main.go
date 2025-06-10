package main

import (
	"fmt"
	"sync"
	"time"
)

// Simbol 'kamar mandi' kita
var kamarMandiCounter int = 0
var kamarMandiLock sync.Mutex // Kunci kamar mandi

// Fungsi 'pengguna kamar mandi'
func penggunaKamarMandi(id int) {
	fmt.Printf("Pengguna %d: Ingin masuk kamar mandi...\n", id)
	kamarMandiLock.Lock() // Mengunci pintu kamar mandi

	fmt.Printf("Pengguna %d: Berhasil masuk kamar mandi. Menggunakan kamar mandi...\n", id)
	kamarMandiCounter++                // Mengubah data di dalam kamar mandi (kritis)
	time.Sleep(500 * time.Millisecond) // Simulasi penggunaan kamar mandi
	fmt.Printf("Pengguna %d: Selesai menggunakan kamar mandi. Counter: %d\n", id, kamarMandiCounter)

	kamarMandiLock.Unlock() // Membuka kunci pintu kamar mandi
	fmt.Printf("Pengguna %d: Keluar dari kamar mandi.\n", id)
}

func main() {
	fmt.Println("Rumah dengan satu kamar mandi dan banyak pengguna.")

	var wg sync.WaitGroup // WaitGroup untuk menunggu semua goroutine selesai

	for i := 1; i <= 5; i++ {
		wg.Add(1) // Menambahkan satu goroutine ke WaitGroup
		go func(id int) {
			defer wg.Done() // Pastikan wg.Done() dipanggil saat goroutine selesai
			penggunaKamarMandi(id)
		}(i)
	}

	wg.Wait() // Menunggu semua goroutine selesai
	fmt.Printf("\nSemua pengguna selesai. Total penggunaan kamar mandi: %d\n", kamarMandiCounter)
}
