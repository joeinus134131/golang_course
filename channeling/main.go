package main

import (
	"fmt"
	"time"
)

func pengirim(pesan chan string) {
	fmt.Println("Pengirim: Mengirim pesan...")
	pesan <- "Halo dari Goroutine!" // Mengirim pesan ke channel
	fmt.Println("Pengirim: Pesan terkirim.")
}

func main() {
	// Membuat channel bertipe string (unbuffered)
	komunikasiChan := make(chan string)

	start := time.Now()

	// Menjalankan goroutine pengirim
	go pengirim(komunikasiChan)

	// Menunggu sebentar untuk memastikan goroutine pengirim berjalan
	time.Sleep(1 * time.Second)

	elapsed := time.Since(start)

	fmt.Println("Main: Menerima pesan...")
	// Menerima pesan dari channel
	// Ini akan memblokir hingga pengirim mengirim nilai
	receivedPesan := <-komunikasiChan
	fmt.Printf("Main: Pesan diterima: \"%s\"\n", receivedPesan)
	fmt.Printf("Waktu eksekusi channel: %s\n", elapsed)
	fmt.Println("Program selesai.")
}
