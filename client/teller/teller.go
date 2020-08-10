package teller

import (
	"fmt"
	"time"
	"golang.org/x/net/context"

	ent "bds/entities"
	conf "bds/config"
	bank "bds/proto"
)

func MenuTeller(user ent.User){
	//fmt.Println("\n=== Menu Teller ===")
	// fmt.Println(user)
	// fmt.Println("Teller Masuk pak ekoooooooooo")
	
	var pil int
	menu := ("=== Menu Teller ===\n" +
		"1.  Setor Tunai\n" +
		"2.  Tarik Tunai\n" +
		"3.  Pindah Buku\n" +
		"4.  Cetak Buku\n" +
		"99. Logout")

	for pil != 99 {
		fmt.Println("\nId User =", user.Id_user)
		fmt.Println(menu)
		fmt.Print("Pilihan\t: ")
		fmt.Scan(&pil)

		switch pil {
		case 1:
			// fmt.Println("Setor Tunai")
			MenuSetorTunai(user.Id_user)
		case 2:
			//fmt.Println("Tarik Tunai")
			MenuTarikTunai(user.Id_user)
		case 3:
			//fmt.Println("Pindah Buku")
			MenuPindahBuku(user.Id_user)
		case 4:
			//fmt.Println("Cetak Buku")
			MenuCetakBuku()
		case 99:
			fmt.Println("Berhasil Logout")
		default:
			fmt.Println("Pilihan salah")			
		}
	}
	return
}

func CariNasabah(rekTujuan int64) (ent.NasabahDetail, error) {
	//Koneksi ke grpc
	conn, err := conf.KoneksiGrpc()
	if err != nil {
		fmt.Println(err)
		return ent.NasabahDetail{}, err
	}
	defer conn.Close()

	//Memanggil funtcion CarNasabahDetail() dari server
	s := bank.NewBankServiceClient(conn)
	response, err := s.CariNasabahDetail(context.Background(), &bank.NasabahDetail{
		NoRekening: rekTujuan,
	})

	//Memasukan nilai yang didapat
	nasabah := ent.NasabahDetail{
		Cif : response.Cif,
		Nama : response.Nama,
		No_rekening : response.NoRekening,
		Saldo : response.Saldo,
	}

	return nasabah, nil
}

func MenuSetorTunai(idUser int64) {
	var (
		//nasabah		ent.NasabahDetail
		rekTujuan	int64
		nominal		int64
		proses		string
		berita		string
		tanggal		string
	)

	//Input id user dan password
	fmt.Print("\n=== Setor Tunai ===\n")
	fmt.Print("No Rekening\t: ")
	fmt.Scan(&rekTujuan)
	fmt.Print("Nominal\t\t: ")
	fmt.Scan(&nominal)
	fmt.Print("Berita\t\t: ")
	fmt.Scan(&berita)

	// Cek Data Rekening
	response, err := CariNasabah(rekTujuan)
	if err != nil {
		fmt.Println(err)
		//return ent.User{}, err
	} else if response.Nama == "" {
		fmt.Println("Nomor Rekening Salah...")
		return
	}

	//Buat objek transaksi
	tanggal = time.Now().Format("2006-01-02 15:04:05")
	transaksi := ent.Transaksi {
		Id_user : idUser,
		No_rekening : rekTujuan,
		Tanggal : tanggal,
		Jenis_transaksi : "st",
		Nominal	: nominal,
		Berita : berita,
	}

	for !(proses == "y") && !(proses == "n") {
		fmt.Print("\n=== Konfirmasi ===\n")
		fmt.Println("Rek Tujuan\t:", rekTujuan)
		fmt.Println("Nama\t\t:", response.Nama)
		fmt.Println("Nominal\t\t:", nominal)
		fmt.Println("Berita\t\t:", berita)
		fmt.Print("Proses Setor Tunai? (y/n) : ")
		fmt.Scan(&proses)

		switch proses {
			case "y":
				transaksi, err := SetorTunai(transaksi)
				if (err != nil) {
					fmt.Println("Setor tunai gagal...")
					fmt.Println(err)
					return
				} else {
					// fmt.Println(transaksi)
					fmt.Println("\nStatus\t\t: Setor Tunai Berhasil")
					fmt.Println("Id Teller\t:", transaksi.Id_user)
					fmt.Println("No Rekening\t:", transaksi.No_rekening)
					fmt.Println("Tanggal\t\t:", transaksi.Tanggal)
					fmt.Println("Kode Transaksi\t:", transaksi.Jenis_transaksi)
					fmt.Println("Nominal\t\t:", transaksi.Nominal)
					fmt.Println("Berita\t\t:", transaksi.Berita)
				}				
			case "n":
				fmt.Println("Batal")
				return
			default:
				fmt.Println("Pilihan salah")			
			}
	}
	return
}

func MenuTarikTunai(idUser int64) {
	
	var (
		//nasabah		ent.NasabahDetail
		rekTujuan	int64
		nominal		int64
		proses		string
		berita		string
		tanggal		string
	)

	//Input id user dan password
	fmt.Print("\n=== Tarik Tunai ===\n")
	fmt.Print("No Rekening\t: ")
	fmt.Scan(&rekTujuan)
	fmt.Print("Nominal\t\t: ")
	fmt.Scan(&nominal)
	fmt.Print("Berita\t\t: ")
	fmt.Scan(&berita)

	// Cek Data Rekening
	response, err := CariNasabah(rekTujuan)
	if err != nil {
		fmt.Println(err)
		//return ent.User{}, err
	} else if response.Nama == "" {
		fmt.Println("Nomor Rekening Salah...")
		return
	}

	//Buat objek transaksi
	tanggal = time.Now().Format("2006-01-02 15:04:05")
	transaksi := ent.Transaksi {
		Id_user : idUser,
		No_rekening : rekTujuan,
		Tanggal : tanggal,
		Jenis_transaksi : "tt",
		Nominal	: nominal,
		Berita : berita,
	}

	for !(proses == "y") && !(proses == "n") {
		fmt.Print("\n=== Konfirmasi ===\n")
		fmt.Println("Rek Tujuan\t:", rekTujuan)
		fmt.Println("Nama\t\t:", response.Nama)
		fmt.Println("Nominal\t\t:", nominal)
		fmt.Println("Berita\t\t:", berita)
		fmt.Print("Proses Tarik Tunai? (y/n) : ")
		fmt.Scan(&proses)

		switch proses {
			case "y":
				transaksi, err := TarikTunai(transaksi)
				//jika error
				if (err != nil) {
					fmt.Println("\nTarik tunai gagal...")
					fmt.Println(err)
				}
				
				//berhasil namun saldo tidak cukup
				if transaksi.Berita == "Saldo Tidak Cukup"{
					fmt.Println("\nStatus :", transaksi.Berita)				
				} else { //berhasil
					// fmt.Println(transaksi)
					fmt.Println("\nStatus\t\t: Tarik Tunai Berhasil")
					fmt.Println("Id Teller\t:", transaksi.Id_user)
					fmt.Println("No Rekening\t:", transaksi.No_rekening)
					fmt.Println("Tanggal\t\t:", transaksi.Tanggal)
					fmt.Println("Kode Transaksi\t:", transaksi.Jenis_transaksi)
					fmt.Println("Nominal\t\t:", transaksi.Nominal)
					fmt.Println("Berita\t\t:", transaksi.Berita)
				}				
			case "n":
				fmt.Println("Batal")
			default:
				fmt.Println("Pilihan salah")			
			}
	}
	return
}

func MenuCetakBuku() {
	var (
		//nasabah		ent.NasabahDetail
		rekening	int64
		// nominal		int64
		// proses		string
		// berita		string
		// tanggal		string
	)

	//Input id user dan password
	fmt.Print("\n=== Cetak Buku ===\n")
	fmt.Print("No Rekening\t: ")
	fmt.Scan(&rekening)

	// Cek Data Rekening
	response, err := CariNasabah(rekening)
	if err != nil {
		fmt.Println(err)
		//return ent.User{}, err
		return
	} else if response.Nama == "" {
		fmt.Println("Nomor Rekening Salah...")
		return
	}
	
	transaksi := ent.Transaksi{No_rekening : rekening}
	err2 := CetakBuku(transaksi)
	if err2 != nil {
		fmt.Println("error : ", err2)
	}
}

func MenuPindahBuku(idUser int64) {
	
	var (
		//nasabah		ent.NasabahDetail
		rekeningDebit	int64
		rekeningKredit	int64
		nominal			int64
		proses			string
		berita			string
		tanggal			string
	)

	//Input id user dan password
	fmt.Print("\n=== Pindah Buku ===\n")
	fmt.Print("Rekening Debit\t: ")
	fmt.Scan(&rekeningDebit)
	fmt.Print("Nominal\t\t: ")
	fmt.Scan(&nominal)
	fmt.Print("Berita\t\t: ")
	fmt.Scan(&berita)
	fmt.Print("Rekening Kredit\t: ")
	fmt.Scan(&rekeningKredit)

	// Cek Data Rekening Debit
	response, err := CariNasabah(rekeningDebit)
	if err != nil {
		fmt.Println(err)
		//return ent.User{}, err
	} else if response.Nama == "" {
		fmt.Println("Nomor rekening debit salah...")
		return
	}

	
	// Cek Data Rekening Kredit
	response2, err2 := CariNasabah(rekeningKredit)
	if err2 != nil {
		fmt.Println(err2)
		//return ent.User{}, err
	} else if response.Nama == "" {
		fmt.Println("Nomor rekening tujuan salah....")
		return
	}

	//Buat objek transaksi
	tanggal = time.Now().Format("2006-01-02 15:04:05")
	transaksiPB := ent.TransaksiPB {
		Id_user 		: idUser,
		Tanggal			: tanggal,
		Nasabah_debit 	: ent.NasabahDetail{No_rekening: rekeningDebit},
		Nasabah_kredit 	: ent.NasabahDetail{No_rekening: rekeningKredit},
		Nominal			: nominal,
		Berita 			: berita,
	}

	for !(proses == "y") && !(proses == "n") {
		fmt.Print("\n=== Konfirmasi ===\n")
		fmt.Println("Rek.Debit\t:", rekeningDebit)
		fmt.Println("Nama\t\t:", response.Nama)
		fmt.Println("Nominal\t\t:", nominal)
		fmt.Println("Berita\t\t:", berita)
		fmt.Println("Rek.Kredit\t:", rekeningKredit)
		fmt.Println("Nama\t\t:", response2.Nama)
		fmt.Print("\nProses Pindah Buku? (y/n) : ")
		fmt.Scan(&proses)

		switch proses {
			case "y":
				_, err := PindahBuku(transaksiPB)
				//jika error
				if (err != nil) {
					fmt.Println("\nPindah buku gagal...")
					fmt.Println(err)
				}				
				//berhasil namun saldo tidak cukup
				if transaksiPB.Berita == "Saldo Tidak Cukup"{
					fmt.Println("\nStatus :", transaksiPB.Berita)				
				} else { //berhasil
					// fmt.Println(transaksi)
					fmt.Println("\nStatus\t\t: Pindah Buku Berhasil")
					fmt.Println("Id Teller\t:", transaksiPB.Id_user)
					fmt.Println("Tanggal\t\t:", transaksiPB.Tanggal)
					fmt.Println("Rekening Debit\t:", transaksiPB.Nasabah_debit.No_rekening)
					fmt.Println("Nama\t\t:", response.Nama)
					fmt.Println("Rekening Kredit\t:", transaksiPB.Nasabah_debit.No_rekening)
					fmt.Println("Nama\t\t:", response2.Nama)
					fmt.Println("Kode Transaksi\t:", "ob")
					fmt.Println("Nominal\t\t:", transaksiPB.Nominal)
					fmt.Println("Berita\t\t:", transaksiPB.Berita)
				}				
			case "n":
				fmt.Println("Batal")
			default:
				fmt.Println("Pilihan salah")			
			}
	}
	return
}





