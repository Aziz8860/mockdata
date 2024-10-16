package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var help bool
	var inputPath, outputPath string

	flag.BoolVar(&help, "h", false, "Tampilkan cara menggunakan")
	flag.BoolVar(&help, "help", false, "Tampilkan cara menggunakan")

	flag.StringVar(&inputPath, "i", "", "Lokasi file JSON sebagai input")
	flag.StringVar(&inputPath, "input", "", "Lokasi file JSON sebagai input")
	flag.StringVar(&outputPath, "o", "", "Lokasi file JSON sebagai output")
	flag.StringVar(&outputPath, "output", "", "Lokasi file JSON sebagai output")

	flag.Parse()

	if help {
		fmt.Println("Cara pakai: mockdata -i input.json -o output.json")
		os.Exit(0)
	}

	if inputPath == "" {
		fmt.Println("Input wajib diisi")
		os.Exit(0) // kalo nol, exit tanpa masalah. Non-zero ada masalah
	}

	if outputPath == "" {
		fmt.Println("Output wajib diisi")
		os.Exit(0)
	}
}
