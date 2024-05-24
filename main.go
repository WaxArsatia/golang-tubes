package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"golang.org/x/text/currency"
)

const NMAX = 100

type Barang struct {
	ID    int
	Nama  string
	Harga int
	Stok  int
}

type SubTransaksi struct {
	ID           int
	NamaBarang   string
	HargaBarang  int
	JumlahBarang int
}

type Transaksi struct {
	ID            int
	Time          string
	Item          [NMAX]SubTransaksi
	NJumlahBarang int
	TotalHarga    int
}

func main() {
	var arrayTransaksi [NMAX]Transaksi
	var arrayBarang [NMAX]Barang
	var nTransaksi, nBarang, indexBarang int
	var choice int

	for choice != 6 {
		mainMenu()
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Input tidak valid!")
		} else {
			switch choice {
			case 1:
				dataBarang(&arrayBarang, &nBarang, &indexBarang)
			case 2:
				tambahTransaksi(&arrayTransaksi, &nTransaksi, &arrayBarang, nBarang)
			case 3:
				ubahTransaksi(&arrayTransaksi, nTransaksi, &arrayBarang, nBarang)
			case 4:
				logTransaksi(arrayTransaksi, nTransaksi)
			case 5:
				omzetTransaksi(arrayTransaksi, nTransaksi)
			case 6:
				fmt.Println("Keluar dari aplikasi")
			default:
				fmt.Println("Pilihan tidak tersedia!")
			}
		}
	}
}

func mainMenu() {
	fmt.Println()
	fmt.Println("==========================")
	fmt.Println("Aplikasi Kasir Minimarket")
	fmt.Println("==========================")
	fmt.Println("1. Data Barang")
	fmt.Println("2. Tambah Transaksi")
	fmt.Println("3. Ubah Transaksi")
	fmt.Println("4. Log Transaksi")
	fmt.Println("5. Omzet Transaksi")
	fmt.Println("6. Exit")
	fmt.Println("==========================")
	fmt.Print("Pilih Menu (1/2/3/4/5/6): ")
}

func dataBarangMenu() {
	fmt.Println()
	fmt.Println("Data Barang")
	fmt.Println(">>>")
	fmt.Println("1. Tambah Barang")
	fmt.Println("2. Ubah Barang")
	fmt.Println("3. Hapus Barang")
	fmt.Println("4. List Barang")
	fmt.Println("5. Tambah Stock Barang")
	fmt.Println("6. Kembali")
	fmt.Print("Pilih Menu (1/2/3/4/5/6): ")
}

func dataBarang(arrayBarang *[NMAX]Barang, nBarang *int, indexBarang *int) {
	var choice int

	for choice != 6 {
		dataBarangMenu()
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Input tidak valid!")
		} else {
			switch choice {
			case 1:
				subTambahBarang(arrayBarang, nBarang, indexBarang)
			case 2:
				subUbahBarang(arrayBarang, *nBarang)
			case 3:
				subHapusBarang(arrayBarang, nBarang)
			case 4:
				subListBarang(*arrayBarang, *nBarang)
			case 5:
				subTambahStock(arrayBarang, *nBarang)
			case 6:
				fmt.Println("Kembali ke menu utama")
			default:
				fmt.Println("Pilihan tidak tersedia!")
			}
		}
	}
}

func subTambahBarang(arrayBarang *[NMAX]Barang, nBarang *int, indexBarang *int) {
	var barangTemp Barang
	barangTemp.ID = *indexBarang + 1

	fmt.Println()
	fmt.Println("Tambah Barang")
	fmt.Println(">>>")

	fmt.Print("Masukkan Nama Barang: ")
	scanner := bufio.NewScanner(os.Stdin)

	var inNama string

	for inNama == "" {
		scanner.Scan()
		inNama = scanner.Text()
	}

	barangTemp.Nama = inNama

	fmt.Print("Masukkan Harga Barang: ")

	_, err := fmt.Scan(&barangTemp.Harga)
	if err != nil {
		fmt.Println("Input Harga Barang tidak valid!")
		return
	}

	fmt.Print("Masukkan Initial Stok Barang: ")

	_, err = fmt.Scan(&barangTemp.Stok)
	if err != nil {
		fmt.Println("Input Stok Barang tidak valid!")
		return
	}

	fmt.Println("Berhasil menambahkan barang!")

	arrayBarang[*nBarang] = barangTemp

	*nBarang++
	*indexBarang++
}

func subUbahBarang(arrayBarang *[NMAX]Barang, nBarang int) {
	fmt.Println()
	fmt.Println("Ubah Barang")
	fmt.Println(">>>")

	var IDBarang int
	fmt.Print("Masukkan ID Barang: ")

	_, err := fmt.Scan(&IDBarang)
	if err != nil {
		fmt.Println("Input ID Barang tidak valid!")
		return
	}

	indexBarang := IDtoIndexBarang(*arrayBarang, nBarang, IDBarang)
	if indexBarang == -1 {
		fmt.Println("ID Barang tidak ditemukan!")
		return
	}

	var choiceUbah int

	for choiceUbah != 3 {

		fmt.Println()
		fmt.Println("1. Ubah Nama Barang")
		fmt.Println("2. Ubah Harga Barang")
		fmt.Println("3. Kembali")

		fmt.Print("Pilih Menu (1/2/3): ")
		_, err = fmt.Scan(&choiceUbah)
		if err != nil {
			fmt.Println("Input tidak valid!")
		} else {

			switch choiceUbah {
			case 1:
				subUbahNamaBarang(arrayBarang, indexBarang)

			case 2:
				subUbahHargaBarang(arrayBarang, indexBarang)

			case 3:
				fmt.Println("Kembali ke menu Data Barang")

			default:
				fmt.Println("Pilihan tidak tersedia!")
			}
		}
	}
}

func subUbahNamaBarang(arrayBarang *[NMAX]Barang, indexBarang int) {
	fmt.Println()
	fmt.Print("Masukkan Nama Barang: ")
	scanner := bufio.NewScanner(os.Stdin)

	var inNama string

	for inNama == "" {
		scanner.Scan()
		inNama = scanner.Text()
	}

	arrayBarang[indexBarang].Nama = inNama
	fmt.Println("Nama barang berhasil diubah!")
}

func subUbahHargaBarang(arrayBarang *[NMAX]Barang, indexBarang int) {
	fmt.Println()
	fmt.Print("Masukkan Harga Barang: ")
	var tempHarga int
	_, err := fmt.Scan(&tempHarga)
	if err != nil {
		fmt.Println("Input Harga Barang tidak valid!")
		return
	}

	arrayBarang[indexBarang].Harga = tempHarga

	fmt.Println("Harga barang berhasil diubah!")
}

func subHapusBarang(arrayBarang *[NMAX]Barang, nBarang *int) {
	fmt.Println()
	fmt.Println("Hapus Barang")
	fmt.Println(">>>")

	var IDBarang int
	fmt.Print("Masukkan ID Barang: ")
	fmt.Scan(&IDBarang)

	indexBarang := IDtoIndexBarang(*arrayBarang, *nBarang, IDBarang)
	if indexBarang == -1 {
		fmt.Println("ID Barang tidak ditemukan!")
		return
	}

	for i := indexBarang; i < *nBarang-1; i++ {
		arrayBarang[i] = arrayBarang[i+1]
	}
	*nBarang--

	fmt.Println("Barang berhasil dihapus!")
}

func subListBarang(arrayBarang [NMAX]Barang, nBarang int) {
	fmt.Println()
	fmt.Println("List Barang")
	fmt.Println(">>>")

	var offset, limit, page int
	const limitPerPage = 5
	page = 1
	var endAvailablePage = nBarang / limitPerPage
	if nBarang%limitPerPage != 0 {
		endAvailablePage++
	}

	offset = (page - 1) * limitPerPage
	limit = offset + limitPerPage
	if limit > nBarang {
		limit = nBarang
	}

	if endAvailablePage == 0 || endAvailablePage == 1 {
		subListBarangPagination(arrayBarang, offset, limit)
	} else {
		for page != 0 {
			offset = (page - 1) * limitPerPage
			limit = offset + limitPerPage
			if limit > nBarang {
				limit = nBarang
			}

			subListBarangPagination(arrayBarang, offset, limit)
			fmt.Println("Halaman", page, "dari", endAvailablePage, "(Total:", nBarang, "barang)")
			fmt.Println()
			fmt.Println("1-" + strconv.Itoa(endAvailablePage) + ". Pilih Halaman")
			fmt.Println("0. Kembali")
			fmt.Print("Pilih Menu (1-" + strconv.Itoa(endAvailablePage) + "/0): ")

			fmt.Scan(&page)

			for page < 0 || page > endAvailablePage {
				fmt.Println("Halaman tidak tersedia!")
				fmt.Print("Pilih Menu (1-" + strconv.Itoa(endAvailablePage) + "/0): ")
				fmt.Scan(&page)
			}
		}
	}
}

func subListBarangPagination(arrayBarang [NMAX]Barang, offset int, limit int) {
	fmt.Println()
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Nama", "Harga", "Stok"})
	for i := offset; i < limit; i++ {
		t.AppendRow([]interface{}{arrayBarang[i].ID, arrayBarang[i].Nama, currency.IDR.Amount(arrayBarang[i].Harga), arrayBarang[i].Stok})
	}
	t.Render()
}

func subTambahStock(arrayBarang *[NMAX]Barang, nBarang int) {
	fmt.Println()
	fmt.Println("Tambah Stock Barang")
	fmt.Println(">>>")

	var IDBarang, JumlahBarang int
	fmt.Print("Masukkan ID Barang: ")
	_, err := fmt.Scan(&IDBarang)
	if err != nil {
		fmt.Println("Input ID Barang tidak valid!")
		return
	}

	indexBarang := IDtoIndexBarang(*arrayBarang, nBarang, IDBarang)
	if indexBarang == -1 {
		fmt.Println("ID Barang tidak ditemukan!")
		return
	}

	fmt.Print("Masukkan Jumlah Barang: ")
	_, err = fmt.Scan(&JumlahBarang)
	if err != nil {
		fmt.Println("Input Jumlah Barang tidak valid!")
		return
	}

	arrayBarang[indexBarang].Stok += JumlahBarang

	fmt.Println("Stok barang berhasil ditambahkan!")
}

func tambahTransaksi(arrayTransaksi *[NMAX]Transaksi, nTransaksi *int, arrayBarang *[NMAX]Barang, nBarang int) {
	var transaksiTemp Transaksi
	var TotalHarga int
	transaksiTemp.ID = *nTransaksi + 1

	fmt.Println()
	fmt.Println("Tambah Transaksi")
	fmt.Println(">>>")

	fmt.Print("Masukkan ID Barang (jika banyak pisahkan dengan spasi): ")
	scanner := bufio.NewScanner(os.Stdin)

	var inIDBarang string

	for inIDBarang == "" {
		scanner.Scan()
		inIDBarang = scanner.Text()
	}

	arrayIDBarang := strings.Fields(inIDBarang)

	for i := 0; i < len(arrayIDBarang); i++ {
		var err error
		transaksiTemp.Item[i].ID, err = strconv.Atoi(arrayIDBarang[i])
		if err != nil {
			fmt.Println("Input ID Barang tidak valid!")
			return
		}

	}

	fmt.Print("Masukkan Jumlah per Barang (jika banyak pisahkan dengan spasi): ")

	var inJumlahPerBarang string

	for inJumlahPerBarang == "" {
		scanner.Scan()
		inJumlahPerBarang = scanner.Text()
	}

	arrayJumlahPerBarang := strings.Fields(inJumlahPerBarang)

	if len(arrayIDBarang) != len(arrayJumlahPerBarang) {
		fmt.Println()
		fmt.Println("Input tidak valid. Jumlah ID Barang dan Jumlah per Barang tidak sama!")
		return
	}

	for i := 0; i < len(arrayJumlahPerBarang); i++ {
		var err error
		transaksiTemp.Item[i].JumlahBarang, err = strconv.Atoi(arrayJumlahPerBarang[i])
		if err != nil {
			fmt.Println("Input Jumlah per Barang tidak valid!")
			return
		}
	}

	transaksiTemp.NJumlahBarang = len(arrayIDBarang)

	var tempArrayBarang = *arrayBarang

	for i := 0; i < transaksiTemp.NJumlahBarang; i++ {
		indexBarang := IDtoIndexBarang(*arrayBarang, nBarang, transaksiTemp.Item[i].ID)
		if indexBarang == -1 {
			fmt.Println()
			fmt.Println("ID Barang", transaksiTemp.Item[i].ID, "tidak ditemukan!")
			return
		}

		if transaksiTemp.Item[i].JumlahBarang <= 0 {
			fmt.Println()
			fmt.Println("Jumlah per Barang", arrayBarang[indexBarang].Nama, "tidak valid!")
			return
		}

		if transaksiTemp.Item[i].JumlahBarang > tempArrayBarang[indexBarang].Stok {
			fmt.Println()
			fmt.Println("Stok barang", arrayBarang[indexBarang].Nama, "tidak mencukupi!")
			return
		}

		transaksiTemp.Item[i].NamaBarang = arrayBarang[indexBarang].Nama

		tempArrayBarang[indexBarang].Stok -= transaksiTemp.Item[i].JumlahBarang

		transaksiTemp.Item[i].HargaBarang = arrayBarang[indexBarang].Harga * transaksiTemp.Item[i].JumlahBarang

		TotalHarga += transaksiTemp.Item[i].HargaBarang
	}

	for i := 0; i < transaksiTemp.NJumlahBarang; i++ {
		indexBarang := IDtoIndexBarang(*arrayBarang, nBarang, transaksiTemp.Item[i].ID)
		arrayBarang[indexBarang].Stok = tempArrayBarang[indexBarang].Stok
	}

	transaksiTemp.Time = time.Now().Local().Format("15:04:05")

	transaksiTemp.TotalHarga = TotalHarga

	arrayTransaksi[*nTransaksi] = transaksiTemp

	*nTransaksi++
}

func ubahTransaksi(arrayTransaksi *[NMAX]Transaksi, nTransaksi int, arrayBarang *[NMAX]Barang, nBarang int) {
	fmt.Println()
	fmt.Println("Ubah Transaksi")
	fmt.Println(">>>")

	var IDTransaksi int
	fmt.Print("Masukkan ID Transaksi: ")
	_, err := fmt.Scan(&IDTransaksi)
	if err != nil {
		fmt.Println("Input ID Transaksi tidak valid!")
		return
	}

	indexTransaksi := IDtoIndexTransaksi(*arrayTransaksi, nTransaksi, IDTransaksi)
	if indexTransaksi == -1 {
		fmt.Println("ID Transaksi tidak ditemukan!")
		return
	}

	var choiceUbah int

	for choiceUbah != 3 {

		fmt.Println()
		fmt.Println("1. Ubah Barang")
		fmt.Println("2. Hapus Barang")
		fmt.Println("3. Kembali")

		fmt.Print("Pilih Menu (1/2/3): ")
		_, err = fmt.Scan(&choiceUbah)
		if err != nil {
			fmt.Println("Input tidak valid!")
		} else {

			switch choiceUbah {
			case 1:
				subUbahJumlahBarangTransaksi(arrayTransaksi, indexTransaksi, arrayBarang, nBarang)

			case 2:
				subHapusBarangTransaksi(arrayTransaksi, indexTransaksi, arrayBarang, nBarang)

			case 3:
				fmt.Println("Kembali ke menu utama")

			default:
				fmt.Println("Pilihan tidak tersedia!")
			}
		}
	}
}

func subUbahJumlahBarangTransaksi(arrayTransaksi *[NMAX]Transaksi, indexTransaksi int, arrayBarang *[NMAX]Barang, nBarang int) {
	fmt.Println()
	fmt.Println("Ubah Barang dari Transaksi")
	fmt.Println(">>>")

	var IDBarang, JumlahBarang int
	fmt.Print("Masukkan ID Barang: ")
	_, err := fmt.Scan(&IDBarang)
	if err != nil {
		fmt.Println("Input ID Barang tidak valid!")
		return
	}

	indexBarang := IDtoIndexBarang(*arrayBarang, nBarang, IDBarang)
	if indexBarang == -1 {
		fmt.Println("ID Barang tidak ditemukan!")
		return
	}

	var indexSubTransaksi int = -1
	for i := 0; i < arrayTransaksi[indexTransaksi].NJumlahBarang; i++ {
		if arrayTransaksi[indexTransaksi].Item[i].ID == IDBarang {
			indexSubTransaksi = i
		}
	}

	if indexSubTransaksi == -1 {
		fmt.Println("ID Barang tidak ditemukan pada transaksi ini!")
		return
	}

	var selectedItem = &arrayTransaksi[indexTransaksi].Item[indexSubTransaksi]

	fmt.Print("Masukkan Jumlah Barang: ")
	_, err = fmt.Scan(&JumlahBarang)
	if err != nil {
		fmt.Println("Input Jumlah Barang tidak valid!")
		return
	}

	if JumlahBarang <= 0 {
		fmt.Println("Jumlah Barang tidak valid!")
		return
	}

	if JumlahBarang > (arrayBarang[indexBarang].Stok + selectedItem.JumlahBarang) {
		fmt.Println("Stok barang tidak mencukupi!")
		return
	}

	arrayBarang[indexBarang].Stok += selectedItem.JumlahBarang
	arrayBarang[indexBarang].Stok -= JumlahBarang

	selectedItem.JumlahBarang = JumlahBarang
	selectedItem.HargaBarang = arrayBarang[indexBarang].Harga * JumlahBarang

	arrayTransaksi[indexTransaksi].TotalHarga = 0
	for i := 0; i < arrayTransaksi[indexTransaksi].NJumlahBarang; i++ {
		arrayTransaksi[indexTransaksi].TotalHarga += arrayTransaksi[indexTransaksi].Item[i].HargaBarang
	}

	fmt.Println("Jumlah Barang berhasil diubah dari Transaksi !")
}

func subHapusBarangTransaksi(arrayTransaksi *[NMAX]Transaksi, indexTransaksi int, arrayBarang *[NMAX]Barang, nBarang int) {
	fmt.Println()
	fmt.Println("Hapus Barangd dari Transaksi")
	fmt.Println(">>>")

	var IDBarang int
	fmt.Print("Masukkan ID Barang: ")
	_, err := fmt.Scan(&IDBarang)
	if err != nil {
		fmt.Println("Input ID Barang tidak valid!")
		return
	}

	indexBarang := IDtoIndexBarang(*arrayBarang, nBarang, IDBarang)
	if indexBarang == -1 {
		fmt.Println("ID Barang tidak ditemukan!")
		return
	}

	var indexSubTransaksi int = -1
	for i := 0; i < arrayTransaksi[indexTransaksi].NJumlahBarang; i++ {
		if arrayTransaksi[indexTransaksi].Item[i].ID == IDBarang {
			indexSubTransaksi = i
		}
	}

	if indexSubTransaksi == -1 {
		fmt.Println("ID Barang tidak ditemukan pada transaksi ini!")
		return
	}

	var selectedItem = &arrayTransaksi[indexTransaksi].Item[indexSubTransaksi]

	arrayBarang[indexBarang].Stok += selectedItem.JumlahBarang

	for i := indexSubTransaksi; i < arrayTransaksi[indexTransaksi].NJumlahBarang-1; i++ {
		arrayTransaksi[indexTransaksi].Item[i] = arrayTransaksi[indexTransaksi].Item[i+1]
	}

	arrayTransaksi[indexTransaksi].NJumlahBarang--

	arrayTransaksi[indexTransaksi].TotalHarga = 0
	for i := 0; i < arrayTransaksi[indexTransaksi].NJumlahBarang; i++ {
		arrayTransaksi[indexTransaksi].TotalHarga += arrayTransaksi[indexTransaksi].Item[i].HargaBarang
	}

	fmt.Println("Barang berhasil dihapus dari Transaksi!")
}

func logTransaksi(arrayTransaksi [NMAX]Transaksi, nTransaksi int) {
	fmt.Println()
	fmt.Println("Log Transaksi")
	fmt.Println(">>>")

	var offset, limit, page int
	const limitPerPage = 5
	page = 1
	var endAvailablePage = nTransaksi / limitPerPage
	if nTransaksi%limitPerPage != 0 {
		endAvailablePage++
	}

	offset = (page - 1) * limitPerPage
	limit = offset + limitPerPage
	if limit > nTransaksi {
		limit = nTransaksi
	}

	if endAvailablePage == 0 || endAvailablePage == 1 {
		subLogTransaksiPagination(arrayTransaksi, offset, limit)
	} else {
		for page != 0 {
			offset = (page - 1) * limitPerPage
			limit = offset + limitPerPage
			if limit > nTransaksi {
				limit = nTransaksi
			}

			subLogTransaksiPagination(arrayTransaksi, offset, limit)
			fmt.Println("Halaman", page, "dari", endAvailablePage, "(Total:", nTransaksi, "transaksi)")
			fmt.Println()
			fmt.Println("1-" + strconv.Itoa(endAvailablePage) + ". Pilih Halaman")
			fmt.Println("0. Kembali")
			fmt.Print("Pilih Menu (1-" + strconv.Itoa(endAvailablePage) + "/0): ")

			fmt.Scan(&page)

			for page < 0 || page > endAvailablePage {
				fmt.Println("Halaman tidak tersedia!")
				fmt.Print("Pilih Menu (1-" + strconv.Itoa(endAvailablePage) + "/0): ")
				fmt.Scan(&page)
			}
		}
	}
}

func subLogTransaksiPagination(arrayTransaksi [NMAX]Transaksi, offset int, limit int) {
	fmt.Println()
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Time", "Nama Barang", "Jumlah", "Harga", "Total Harga"})
	for i := offset; i < limit; i++ {
		for j := 0; j < arrayTransaksi[i].NJumlahBarang; j++ {
			var rowData = table.Row{"", "", arrayTransaksi[i].Item[j].NamaBarang, arrayTransaksi[i].Item[j].JumlahBarang, currency.IDR.Amount(arrayTransaksi[i].Item[j].HargaBarang), ""}
			if j == 0 {
				rowData[0] = arrayTransaksi[i].ID
				rowData[1] = arrayTransaksi[i].Time
				rowData[5] = currency.IDR.Amount(arrayTransaksi[i].TotalHarga)
			}
			t.AppendRow(rowData)
		}
		t.AppendSeparator()
	}
	t.Render()
}

func omzetTransaksi(arrayTransaksi [NMAX]Transaksi, nTransaksi int) {
	fmt.Println()
	fmt.Println("Omzet Transaksi")
	fmt.Println(">>>")

	var TotalOmzet int
	for i := 0; i < nTransaksi; i++ {
		TotalOmzet += arrayTransaksi[i].TotalHarga
	}

	fmt.Println("Banyak Transaksi:", nTransaksi)
	fmt.Println("Total Omzet:", currency.IDR.Amount(TotalOmzet))
}

func IDtoIndexBarang(arrayBarang [NMAX]Barang, nBarang int, ID int) int {
	var index int = -1
	var left, mid, right int
	left = 0
	right = nBarang - 1

	for left <= right && index == -1 {
		mid = (left + right) / 2
		if arrayBarang[mid].ID == ID {
			index = mid
		} else if ID > arrayBarang[mid].ID {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return index
}

func IDtoIndexTransaksi(arrayTransaksi [NMAX]Transaksi, nTransaksi int, ID int) int {
	// Use Binary Search Algorithm
	var index int = -1
	var left, mid, right int
	left = 0
	right = nTransaksi - 1

	for left <= right && index == -1 {
		mid = (left + right) / 2
		if arrayTransaksi[mid].ID == ID {
			index = mid
		} else if ID > arrayTransaksi[mid].ID {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return index
}
