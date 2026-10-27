package main

import (
	"fmt"
	"time"
)

func riskyOperation() {
	// Catat waktu mulai untuk benchmarking
	start := time.Now()

	// Satu blok defer func untuk menangani recover DAN benchmark sekaligus
	defer func() {
		// 1. Tangkap panic jika terjadi (Recovery)
		if r := recover(); r != nil {
			fmt.Printf("[RECOVERED] Sistem selamat dari panic dengan pesan: %v\n", r)
		}

		// 2. Hitung durasi eksekusi (Benchmark) 
		// Ini akan tetap tereksekusi meskipun fungsi mengalami panic!
		duration := time.Since(start)
		fmt.Printf("[BENCHMARK] Fungsi berjalan selama: %v\n", duration)
	}()

	fmt.Println("Memulai proses berat...")
	time.Sleep(1 * time.Second) // Simulasi kerja

	// Simulasi kondisi darurat yang memicu panic (crash)
	panic("koneksi database terputus secara mendadak!")

	// Baris ini tidak akan pernah dieksekusi karena panic terjadi di atas
	fmt.Println("Proses selesai dengan normal.")
}

func main() {
	fmt.Println("Program utama dimulai.")
	
	// Memanggil fungsi yang berbahaya
	riskyOperation()

	fmt.Println("Program utama tetap berjalan normal dan tidak crash!")
}