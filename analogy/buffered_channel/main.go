package main

import (
	"fmt"
	"time"
)

// Koki yang mengirim makanan ke keranjang (buffered channel)
func kokiPengirimBuffered(nama string, keranjangMakanan chan string) {
	for i := 1; i <= 3; i++ {
		makanan := fmt.Sprintf("Sup Ayam %s #%d", nama, i)
		fmt.Printf("Koki %s: Menaruh %s ke keranjang.\n", nama, makanan)
		keranjangMakanan <- makanan        // Mengirim makanan ke buffered channel
		time.Sleep(100 * time.Millisecond) // jeda sedikit
	}
	// Penting: Koki selesai mengirim, tutup keranjang
	close(keranjangMakanan)
	fmt.Printf("Koki %s: Selesai menaruh semua sup dan menutup keranjang.\n", nama)
}

// Koki yang menerima makanan dari keranjang
func kokiPenerimaBuffered(keranjangMakanan chan string) {
	fmt.Println("\nKoki Penerima: Mulai mengambil makanan dari keranjang...")
	// Menggunakan for...range untuk terus menerima sampai channel ditutup
	for makanan := range keranjangMakanan {
		fmt.Printf("Koki Penerima: Mengambil %s dari keranjang.\n", makanan)
		time.Sleep(300 * time.Millisecond) // Simulasi waktu mengambil
	}
	fmt.Println("Koki Penerima: Keranjang sudah kosong dan ditutup.")
}

func main() {
	// Membuat 'keranjang makanan' (buffered channel) dengan kapasitas 2
	keranjangMakanan := make(chan string, 2)

	fmt.Println("Manajer: Membuka restoran dan keranjang makanan (kapasitas 2).")

	// Koki pengirim bekerja
	go kokiPengirimBuffered("Lina", keranjangMakanan)

	// Koki penerima bekerja
	go kokiPenerimaBuffered(keranjangMakanan)

	// Memberi waktu yang cukup agar semua proses selesai
	time.Sleep(3 * time.Second)

	fmt.Println("\nManajer: Restoran tutup. Semua makanan terdistribusi.")
}
