package main

import (
	"fmt"
	"time"
)

// Koki yang mengirim makanan
func kokiPengirim(nama string, pipaMakanan chan string) {
	makanan := "Nasi Goreng Spesial " + nama
	fmt.Printf("Koki %s: Menaruh %s ke pipa...\n", nama, makanan)
	pipaMakanan <- makanan // Mengirim makanan ke channel
	fmt.Printf("Koki %s: Selesai menaruh %s.\n", nama, makanan)
}

// Koki yang menerima makanan
func kokiPenerima(pipaMakanan chan string) {
	fmt.Println("Koki Penerima: Menunggu makanan dari pipa...")
	makananDiterima := <-pipaMakanan // Menerima makanan dari channel
	fmt.Printf("Koki Penerima: Mengambil %s dari pipa. Siap diantar!\n", makananDiterima)
}

func main() {
	// Membuat 'pipa makanan' (channel unbuffered)
	// Artinya, pengirim harus menunggu penerima, dan sebaliknya.
	pipaMakanan := make(chan string)

	fmt.Println("Manajer: Membuka restoran dan pipa komunikasi.")

	// Koki pengirim bekerja (goroutine)
	go kokiPengirim("Chef Bintang", pipaMakanan)

	// Koki penerima bekerja (goroutine)
	go kokiPenerima(pipaMakanan)

	// Memberi waktu agar kedua goroutine sempat berkomunikasi
	time.Sleep(2 * time.Second)

	fmt.Println("Manajer: Proses komunikasi selesai.")
}
