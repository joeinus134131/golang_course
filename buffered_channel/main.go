package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d: Mulai memproses job %d\n", id, j)
		time.Sleep(50 * time.Millisecond) // Simulasi pekerjaan
		fmt.Printf("Worker %d: Selesai memproses job %d\n", id, j)
		results <- j * 2 // Kirim hasil ke channel results
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)    // Buffered channel untuk jobs
	results := make(chan int, numJobs) // Buffered channel untuk results

	// Memulai beberapa worker goroutine
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Mengirim jobs ke channel jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
		fmt.Printf("Main: Mengirim job %d\n", j)
	}
	close(jobs) // Penting: tutup channel jobs setelah semua job dikirim

	// Menerima hasil dari channel results
	fmt.Println("\nMain: Menerima hasil...")
	for a := 1; a <= numJobs; a++ {
		result := <-results
		fmt.Printf("Main: Menerima hasil: %d\n", result)
	}
	close(results) // Opsional: tutup channel results setelah semua hasil diterima
	fmt.Println("Program selesai.")
}
