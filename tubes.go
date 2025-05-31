package main

import "fmt"

const NMAX int = 10

type startup struct {
	idStartUp    string
	nama         string
	bidang       string
	tahunBerdiri string
	pendanaan    float64
}
type tabStartUp [NMAX]startup

func main() {
	var data tabStartUp
	var nData int
	fmt.Println("-----------------------")
	fmt.Println("|  MANAJEMEN START-UP |")
	fmt.Println("-----------------------")
	fmt.Println("| 		Pilih menu     |")
	fmt.Println("|1.                   |")
	fmt.Println("|2.                   |")
	fmt.Println("|3.                   |")
	fmt.Println("|4.                   |")
	fmt.Println("|5.                   |")
	fmt.Println("|6.                   |")
	fmt.Println("|7.Keluar			   |")
	fmt.Println("-----------------------")
	fmt.Println("Pilih menu: ")
	var pilihan string
	fmt.Scanln(&pilihan)
	switch pilihan {
	case "1":
		dataStartUp(&data, nData)

	case "7":
		return
	}

}

func dataStartUp(A *tabStartUp, n int) {
	var i int
	if n > NMAX {
		n = NMAX
	}

	for i = 0; i < n; i++ {
		fmt.Println("Masukkan ID: ")
		fmt.Scanln(&A[i].idStartUp)
		fmt.Println("Masukkan Nama: ")
		fmt.Scanln(&A[i].nama)
		fmt.Println("Masukkan Bidang: ")
		fmt.Scanln(&A[i].bidang)
		fmt.Println("Masukkan Tahun Berdiri: ")
		fmt.Scanln(&A[i].tahunBerdiri)
		fmt.Println("Masukkan Jumlah Pendanaan: ")
		fmt.Scanln(&A[i].pendanaan)
	}
}

func cetakStartUp(A tabStartUp, n int) {
	var i int

	for i = 0; i < n; i++ {
		fmt.Println(A[i].idStartUp, A[i].nama, A[i].bidang, A[i].tahunBerdiri, A[i].pendanaan)
	}
}

//binary searching//
func binarySearch(A tabStartUp, n int, id string) int {
	var left, right, mid int

	left = 0
	right = n - 1

	for left <= right {
		mid = (left + right) / 2
		if A[mid].idStartUp == id {
			return mid
		} else if id < A[mid].idStartUp {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

//insertion sort//
func insertionSort(A *tabStartUp, n int) {
	var pass, i int
	var temp int

	for pass = 1; pass < n; pass++ {
		i = pass
		temp = A[pass].tahunBerdiri
		for i > 0 && temp < A[i-1].tahunBerdiri {
			A[i].tahunBerdiri = A[i-1].tahunBerdiri
			i = i - 1
		}
		A[i].tahunBerdiri = temp
	}
}


// sequential search
func sequentialSearch(A tabStartup, n int, key string) int {
	for i := 0; i < n; i++ {
		if A[i].nama == key {
			return i
		}
	}
	return -1
}

// pencarian nilai ekstrim
func cariPendanaanMaksimum(A tabStartUp, n int) int {
	maxIdx := 0
	for i := 1; i < n; i++ {
		if A[i].pendanaan > A[maxIdx].pendanaan {
			maxIdx = i
		}
	}
	return maxIdx
}