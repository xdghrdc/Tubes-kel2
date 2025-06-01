package main

import ("fmt")

const NMAX int = 10
var banyakData int

type startup struct {
	idStartUp    string
	nama         string
	bidang       string
	tahunBerdiri int
	pendanaan    float64
}
type tabStartUp [NMAX]startup

func main() {
	menu()
	pilihanMenu()
	penutup()
}

func menu(){
	fmt.Println("------------------------")
	fmt.Println("|  MANAJEMEN START-UP  |")
	fmt.Println("------------------------")
	fmt.Println("| 		Pilih menu   	|")
	fmt.Println("|1. Baca Data          |")
	fmt.Println("|2. Cetak Data         |")
	fmt.Println("|3. Cari by ID         |")
	fmt.Println("|4. Cari by Nama       |")
	// fmt.Println("|5. Urutkan Tahun      |")
	fmt.Println("|6. Pendanaan Terbesar |")
	fmt.Println("|7.Keluar			 	|")
	fmt.Println("------------------------")
	fmt.Println("Pilih menu: ")
}

func pilihanMenu(){
	var data tabStartUp
	var pilihan int
	for{
		fmt.Scanln(&pilihan)
		fmt.Println()
		
		switch pilihan {
		case 1:
			dataStartUp(&data)
		case 2: 
			cetakStartUp(data)
		case 3:
			cariById(&data)
		case 4:
			cariByNama(data)
		case 5:
			urutkan(&data)
		case 6:
			pendanaan(&data)
		case 7:
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
		menu()
	}
}

func dataStartUp(A *tabStartUp) {
	var i, n int
	fmt.Print("\nBerapa banyak data yang ingin ditambahkan? ")
	fmt.Scan(&n)
	
	if n > NMAX {
		n = NMAX
	}

	for i = 0; i < n; i++ {
		fmt.Print("Masukkan ID: ")
		fmt.Scan(&A[i].idStartUp)
		fmt.Print("Masukkan Nama: ")
		fmt.Scan(&A[i].nama)
		fmt.Print("Masukkan Bidang: ")
		fmt.Scan(&A[i].bidang)
		fmt.Print("Masukkan Tahun Berdiri: ")
		fmt.Scan(&A[i].tahunBerdiri)
		fmt.Print("Masukkan Jumlah Pendanaan: ")
		fmt.Scan(&A[i].pendanaan)
		fmt.Println()
		Fmt.println("Data Ditambahkan")
	}
	banyakData = n
}

func cetakStartUp(A tabStartUp) {
	var i, nomor int

	nomor = 1
	fmt.Println("Data tercatat: ")
	for i = 0; i < banyakData; i++ {
		if A[i].nama != ""{
		fmt.Printf("%d. %s, %s, %d, %s, %.2f\n",nomor, A[i].idStartUp, A[i].nama, A[i].bidang, A[i].tahunBerdiri, A[i].pendanaan)
		fmt.Println()
		nomor++
		}
	}
}

func cariById(A *tabStartUp){
	var id string
	var idx int
	
	fmt.Print("Masukkan ID yang dicari: ")
	fmt.Scanln(&id)
	idx = binarySearch(*A, banyakData, id)
	if idx != -1{
		fmt.Println("Data Ditemukan: ")
		fmt.Printf("%s, %s, %s, %d, %.2f\n", A[idx].idStartUp, A[idx].nama, A[idx].bidang, A[idx].tahunBerdiri, A[idx].pendanaan)
	}else{
		fmt.Println("Data Tidak Ditemukan !")
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
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func urutkan(A *tabStartUp){
	pilihanSort(A)
}

func pilihanSort(A *tabStartUp){
	var pilih int
    for {
		fmt.Println("==========================")
        fmt.Println("1. Urutkan Stok Descending")
        fmt.Println("2. Urutkan Stok Ascending")
        fmt.Println("0. Kembali ke Menu Utama")
		fmt.Println("==========================")
        fmt.Print("Pilih (Masukkan Angka): ")
        fmt.Scan(&pilih)
        fmt.Println()
		
        switch pilih {
        case 1:
            insertionSortDesc(A, banyakData)
            cetakStartUp(*A)
        case 2:
            insertionSortAsc(A, banyakData)
            cetakStartUp(*A)
        case 0:
            return
        default:
            fmt.Println("Pilihan tidak valid!")
        }
    }
}

//insertion sort//
func insertionSortDesc(A *tabStartUp, n int) {
	var pass, i int
	var temp startup

	for pass = 1; pass < n; pass++ {
		temp = A[pass]
		i = pass
		for i > 0 && temp.tahunBerdiri > A[i-1].tahunBerdiri {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
	}
}

func insertionSortAsc(A *tabStartUp, n int) {
	var pass, i int
	var temp startup

	for pass = 1; pass < n; pass++ {
		temp = A[pass]
		i = pass
		for i > 0 && temp.tahunBerdiri < A[i-1].tahunBerdiri {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
	}
}

func cariByNama(A tabStartUp){
	var nama string
	var idx int
	
	fmt.Print("Masukkan nama yang dicari: ")
	fmt.Scanln(&nama)
	
	idx = sequentialSearch(A, banyakData, nama)
	if idx != -1{
		fmt.Println("Data Ditemukan: ")
		fmt.Printf("%s, %s, %s, %s, %.2f\n", A[idx].idStartUp, A[idx].nama, A[idx].bidang, A[idx].tahunBerdiri, A[idx].pendanaan)
	}else{
		fmt.Println("Data Tidak Ditemukan !")
	}
}

// sequential search
func sequentialSearch(A tabStartUp, n int, key string) int {
	for i := 0; i < n; i++ {
		if A[i].nama == key {
			return i
		}
	}
	return -1
}

func pendanaan(A *tabStartUp){
	pilihanNilaiEkstrim(A)
}

func pilihanNilaiEkstrim(A *tabStartUp){
	var pilih int
    for {
		fmt.Println("==========================")
        fmt.Println("1. Cari Pendanaan Maksimum")
        fmt.Println("2. Cari Pendanaan Minimum")
        fmt.Println("0. Kembali ke Menu Utama")
		fmt.Println("==========================")
        fmt.Print("Pilih (Masukkan Angka): ")
        fmt.Scan(&pilih)
        fmt.Println()
		
        switch pilih {
        case 1:
            pendanaanMaksimum(*A)
        case 2:
            pendanaanMinimum(*A)
        case 0:
            return
        default:
            fmt.Println("Pilihan tidak valid!")
        }
    }
}

func pendanaanMaksimum(A tabStartUp){
	var idx int
	idx = cariPendanaanMaksimum(A, banyakData)
	fmt.Println("Pendanaan Terbesar Dimiliki Oleh: ")
	fmt.Printf("%s, %s, %s, %d, %.2f\n", A[idx].idStartUp, A[idx].nama, A[idx].bidang, A[idx].tahunBerdiri, A[idx].pendanaan)
}

// pencarian nilai ekstrim
func cariPendanaanMaksimum(A tabStartUp, n int) int {
	var i, maxIdx int
	maxIdx = 0
	for i = 1; i < n; i++ {
		if A[i].pendanaan > A[maxIdx].pendanaan {
			maxIdx = i
		}
	}
	return maxIdx
}

func pendanaanMinimum(A tabStartUp){
	var idx int
	idx = cariPendanaanMaksimum(A, banyakData)
	fmt.Println("Pendanaan Terbesar Dimiliki Oleh: ")
	fmt.Printf("%s, %s, %s, %d, %.2f\n", A[idx].idStartUp, A[idx].nama, A[idx].bidang, A[idx].tahunBerdiri, A[idx].pendanaan)
}

func cariPendanaanMinimum(A tabStartUp, n int) int {
	var i, minIdx int
	minIdx = 0
	for i = 1; i < n; i++ {
		if A[i].pendanaan < A[minIdx].pendanaan {
			minIdx = i
		}
	}
	return minIdx
}

func penutup(){
	fmt.Println("=========================================================")
	fmt.Println()
	fmt.Println("Terima kasih telah menggunakan aplikasi.")
	fmt.Println()
	fmt.Println("=========================================================")
}