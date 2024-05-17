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

type Transaksi struct {
	ID              int
	Time            string
	IDBarang        [NMAX]int
	JumlahPerBarang [NMAX]int
	HargaPerJumlah  [NMAX]int
	NJumlahBarang   int
	TotalHarga      int
}

func main() {
	var arrayTransaksi [NMAX]Transaksi
	var arrayBarang [NMAX]Barang
	var nTransaksi, nBarang int
	var choice int
outerLoop:
	for choice != 5 {
		mainMenu()
		fmt.Scan(&choice)
	innerSwitch:
		switch choice {
		case 1:
			dataBarang(&arrayBarang, &nBarang)
			break innerSwitch
		case 2:
			tambahTransaksi(&arrayTransaksi, &nTransaksi, &arrayBarang, nBarang)
			break innerSwitch
		case 3:
			logTransaksi(arrayTransaksi, nTransaksi, arrayBarang, nBarang)
			break innerSwitch
		case 4:
			omzetTransaksi(arrayTransaksi, nTransaksi)
			break innerSwitch
		case 5:
			break outerLoop
		default:
			fmt.Println("Pilihan tidak tersedia!")
			break innerSwitch
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
	fmt.Println("3. Log Transaksi")
	fmt.Println("4. Omzet Transaksi")
	fmt.Println("5. Exit")
	fmt.Println("==========================")
	fmt.Print("Pilih Menu (1/2/3/4/5): ")
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

func dataBarang(arrayBarang *[NMAX]Barang, nBarang *int) {
	var choice int

outerLoop:
	for choice != 6 {
		dataBarangMenu()
		fmt.Scan(&choice)
	innerSwitch:
		switch choice {
		case 1:
			subTambahBarang(arrayBarang, nBarang)
			break innerSwitch
		case 2:
			subUbahBarang(arrayBarang, *nBarang)
			break innerSwitch
		case 3:
			subHapusBarang(arrayBarang, nBarang)
			break innerSwitch
		case 4:
			subListBarang(*arrayBarang, *nBarang)
			break innerSwitch
		case 5:
			subTambahStock(arrayBarang, *nBarang)
			break innerSwitch
		case 6:
			break outerLoop
		default:
			fmt.Println("Pilihan tidak tersedia!")
			break innerSwitch
		}
	}
}

func subTambahBarang(arrayBarang *[NMAX]Barang, nBarang *int) {
	var barangTemp Barang
	barangTemp.ID = *nBarang + 1

	fmt.Println()
	fmt.Println("Tambah Barang")
	fmt.Println(">>>")

	fmt.Print("Masukkan Nama Barang: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	barangTemp.Nama = scanner.Text()

	if barangTemp.Nama == "" {
		fmt.Println("Nama Barang tidak boleh kosong!")
		return
	}

	fmt.Print("Masukkan Harga Barang: ")
	scanner.Scan()
	plainHargaBarang := scanner.Text()
	arrayHargaBarang := strings.Fields(plainHargaBarang)

	if len(arrayHargaBarang) != 1 {
		fmt.Println("Input Harga Barang tidak valid!")
		return
	}

	var err error
	barangTemp.Harga, err = strconv.Atoi(arrayHargaBarang[0])
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
}

func subUbahBarang(arrayBarang *[NMAX]Barang, nBarang int) {
	fmt.Println()
	fmt.Println("Ubah Barang")
	fmt.Println(">>>")

	var IDBarang int
	fmt.Print("Masukkan ID Barang: ")
	fmt.Scan(&IDBarang)

	indexBarang := IDtoIndexBarang(*arrayBarang, nBarang, IDBarang)
	if indexBarang == -1 {
		fmt.Println("ID Barang tidak ditemukan!")
		return
	}

	var tempBarang Barang

	fmt.Print("Masukkan Nama Barang (Kosongkan jika tidak berubah): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	tempBarang.Nama = scanner.Text()

	if tempBarang.Nama == "" {
		tempBarang.Nama = arrayBarang[indexBarang].Nama
	}

	fmt.Print("Masukkan Harga Barang (Kosongkan jika tidak berubah): ")
	scanner.Scan()
	plainHargaBarang := scanner.Text()
	arrayHargaBarang := strings.Fields(plainHargaBarang)

	if len(arrayHargaBarang) == 0 {
		tempBarang.Harga = arrayBarang[indexBarang].Harga
	} else {
		if len(arrayHargaBarang) != 1 {
			fmt.Println("Input Harga Barang tidak valid!")
			return
		}

		var err error
		tempBarang.Harga, err = strconv.Atoi(arrayHargaBarang[0])
		if err != nil {
			fmt.Println("Input Harga Barang tidak valid!")
			return
		}
	}

	if tempBarang.Nama == arrayBarang[indexBarang].Nama && tempBarang.Harga == arrayBarang[indexBarang].Harga {
		fmt.Println("Tidak ada perubahan!")
		return
	}

	arrayBarang[indexBarang].Nama = tempBarang.Nama
	arrayBarang[indexBarang].Harga = tempBarang.Harga

	fmt.Println("Barang berhasil diubah!")
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
	fmt.Scan(&IDBarang)
	fmt.Print("Masukkan Jumlah Barang: ")
	fmt.Scan(&JumlahBarang)

	indexBarang := IDtoIndexBarang(*arrayBarang, nBarang, IDBarang)
	if indexBarang == -1 {
		fmt.Println("ID Barang tidak ditemukan!")
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
	scanner.Scan()
	plainIDBarang := scanner.Text()
	arrayIDBarang := strings.Fields(plainIDBarang)
	for i := 0; i < len(arrayIDBarang); i++ {
		var err error
		transaksiTemp.IDBarang[i], err = strconv.Atoi(arrayIDBarang[i])
		if err != nil {
			fmt.Println("Input ID Barang tidak valid!")
			return
		}

	}

	fmt.Print("Masukkan Jumlah per Barang (jika banyak pisahkan dengan spasi): ")
	scanner.Scan()
	plainJumlahPerBarang := scanner.Text()
	arrayJumlahPerBarang := strings.Fields(plainJumlahPerBarang)
	for i := 0; i < len(arrayJumlahPerBarang); i++ {
		var err error
		transaksiTemp.JumlahPerBarang[i], err = strconv.Atoi(arrayJumlahPerBarang[i])
		if err != nil {
			fmt.Println("Input Jumlah per Barang tidak valid!")
			return
		}
	}

	if len(arrayIDBarang) != len(arrayJumlahPerBarang) {
		fmt.Println()
		fmt.Println("Input tidak valid. Jumlah ID Barang dan Jumlah per Barang tidak sama!")
		return
	}

	transaksiTemp.NJumlahBarang = len(arrayIDBarang)

	for i := 0; i < len(arrayIDBarang); i++ {
		indexBarang := IDtoIndexBarang(*arrayBarang, nBarang, transaksiTemp.IDBarang[i])
		if indexBarang == -1 {
			fmt.Println()
			fmt.Println("ID Barang tidak ditemukan!")
			return
		}
		if transaksiTemp.JumlahPerBarang[i] > arrayBarang[indexBarang].Stok {
			fmt.Println()
			fmt.Println("Stok barang tidak mencukupi!")
			return
		}

		arrayBarang[indexBarang].Stok -= transaksiTemp.JumlahPerBarang[i]

		transaksiTemp.HargaPerJumlah[i] = arrayBarang[indexBarang].Harga * transaksiTemp.JumlahPerBarang[i]
		TotalHarga += transaksiTemp.HargaPerJumlah[i]
	}

	transaksiTemp.Time = time.Now().Local().Format("15:04:05")

	transaksiTemp.TotalHarga = TotalHarga

	arrayTransaksi[*nTransaksi] = transaksiTemp

	*nTransaksi++
}

func logTransaksi(arrayTransaksi [NMAX]Transaksi, nTransaksi int, arrayBarang [NMAX]Barang, nBarang int) {
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
		subLogTransaksiPagination(arrayTransaksi, arrayBarang, nBarang, offset, limit)
	} else {
		for page != 0 {
			offset = (page - 1) * limitPerPage
			limit = offset + limitPerPage
			if limit > nTransaksi {
				limit = nTransaksi
			}

			subLogTransaksiPagination(arrayTransaksi, arrayBarang, nBarang, offset, limit)
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

func subLogTransaksiPagination(arrayTransaksi [NMAX]Transaksi, arrayBarang [NMAX]Barang, nBarang int, offset int, limit int) {
	fmt.Println()
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Time", "Nama Barang", "Jumlah", "Harga", "Total Harga"})
	for i := offset; i < limit; i++ {
		for j := 0; j < arrayTransaksi[i].NJumlahBarang; j++ {
			indexBarang := IDtoIndexBarang(arrayBarang, nBarang, arrayTransaksi[i].IDBarang[j])
			var rowData = table.Row{"", "", arrayBarang[indexBarang].Nama, arrayTransaksi[i].JumlahPerBarang[j], currency.IDR.Amount(arrayTransaksi[i].HargaPerJumlah[j]), ""}
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
