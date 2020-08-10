package cs

import (
	"fmt"
	// "time"
	// "golang.org/x/net/context"

	ent "bds/entities"
	// conf "bds/config"
	bank "bds/proto"
)

func MenuCS(user ent.User){
	//fmt.Println("\n=== Menu CS ===")
	// fmt.Println(user)
	// fmt.Println("CS Masuk pak ekoooooooooo")
	
	var pil int
	menu := ("=== Menu Customer Service ===\n" +
		"1.  Pembuatan Rekening\n" +
		"2.  Manajemen Nasabah\n" +
		"99. Logout")

	for pil != 99 {
		fmt.Println("\nId User =", user.Id_user)
		fmt.Println(menu)
		fmt.Print("Pilihan\t: ")
		fmt.Scan(&pil)

		switch pil {
		case 1:
			//fmt.Println("Menu Pembuatan Rekening")
			MenuPembuatanRekening(user.Id_user)
		case 2:
			//fmt.Println("Menu Manajemen Nasabah")
			MenuManajemenNasabah(user.Id_user)
		case 99:
			fmt.Println("Berhasil Logout")
		default:
			fmt.Println("Pilihan salah")			
		}
	}
	return
}

func MenuPembuatanRekening(idUser int64) {	

	//Masuk ke menu pembuatan rekening
	var pil int
	menu := ("\n=== Menu Pembuatan Rekening ===\n" +
		"1.  Pencarian CIF\n" +
		"2.  Pembuatan CIF\n" +
		"3.  Pembuatan Rekening Tabungan\n" +
		"99. Kembali ke Menu CS")

	for pil != 99 {
		fmt.Println(menu)
		fmt.Print("Pilihan\t: ")
		fmt.Scan(&pil)

		switch pil {
		case 1:
			var cif int64
			fmt.Println("\n=== Pencarian CIF ===")
			fmt.Print("Masukan CIF/NIK\t:")
			fmt.Scan(&cif)
			response, err := PencarianCIF(cif)
			if err != nil {
				fmt.Println("Pencarian Gagal...")
			}

			if response == -1 {
				fmt.Println("NIK atau CIF tidak ditemukan...")
			}
		case 2:
			var (
				nik 			int64
				nama			string
				tempatLahir		string
				tanggalLahir	string
				alamat			string
				noTelepon		string
			)
			fmt.Println("\n=== Pembuatan CIF ===")
			fmt.Print("Masukan NIK\t\t:")
			fmt.Scan(&nik)
			
			//cek nik apakah sudah ada atau belum
			response, err := PencarianCIF(nik)
			if err != nil {
				fmt.Println("Pencarian Gagal...")
				break
			}

			if response == 1 {
				fmt.Println("NIK sudah ada dalam database...")
				break
			}

			fmt.Print("Masukan Nama\t\t:")
			fmt.Scan(&nama)
			fmt.Print("Masukan Tempat Lahir\t:")
			fmt.Scan(&tempatLahir)
			fmt.Print("Masukan Tanggal Lahir\t:")
			fmt.Scan(&tanggalLahir)
			fmt.Print("Masukan Alamat\t\t:")
			fmt.Scan(&alamat)
			fmt.Print("Masukan No Telepon\t:")
			fmt.Scan(&noTelepon)

			nasabah := bank.Nasabah{
				Nik				: nik,
				Nama			: nama,
				TempatLahir		: tempatLahir,
				TanggalLahir	: tanggalLahir,
				Alamat			: alamat,
				NoTelepon		: noTelepon,
			}
			
			response2, err2 := PembuatanCif(&nasabah)
			if err2 != nil {
				fmt.Println("Pembuatan CIF gagal... : ", err2)
				break
			}

			fmt.Println("\nPembuatan CIF Berhasil...")
			fmt.Println("Cif\t\t:", response2.Cif)
			fmt.Println("Nik\t\t:", response2.Nik)
			fmt.Println("Nama\t\t:", response2.Nama)
			fmt.Println("Tempat Lahir\t:", response2.TempatLahir)
			fmt.Println("Tanggal Lahir\t:", response2.TanggalLahir)
			fmt.Println("Alamat\t\t:", response2.Alamat)
			fmt.Println("NoTelepon\t:", response2.NoTelepon)

		case 3:
			var (
				cif		int64
				saldo 	int64
			)

			fmt.Println("\n=== Pembuatan Rekening Tabungan ===")
			fmt.Print("Masukan CIF\t\t:")
			fmt.Scan(&cif)

			//Cek CIF sudah ada atau belum
			response, err := PencarianCIF(cif)
			if err != nil {
				fmt.Println("Pencarian Gagal...")
				break
			}
			if response == -1 {
				fmt.Println("CIF tidak ditemukan, silahkan buat cif terlebih dahulu...")
				break
			}

			//Melanjutkan pembuatan rekening
			fmt.Print("\nMasukan Saldo Awal\t:")
			fmt.Scan(&saldo)

			nasabahDetail := bank.NasabahDetail{
				Cif		: cif,
				Saldo	: saldo,
			}

			//Memanggil PembuatanTabungan()
			response2, err2 := PembuatanTabungan(&nasabahDetail)
			if err2 != nil {
				fmt.Println("Pencarian Gagal...")
				break
			}

			//Mengambiil data dari response2
			nasabahDetail.Nama = response2.Nama
			nasabahDetail.NoRekening = response2.NoRekening

			fmt.Println("\nPembuatan Rekening Berhasil...")
			fmt.Println("Cif\t\t:", nasabahDetail.Cif)
			fmt.Println("Nama\t\t:", nasabahDetail.Nama)
			fmt.Println("No Rekening\t:", nasabahDetail.NoRekening)
			fmt.Println("Saldo\t\t:", nasabahDetail.Saldo)

		case 99:
			fmt.Println("Berhasil Logout")
		default:
			fmt.Println("Pilihan salah")			
		}
	}
	return
}

func MenuManajemenNasabah(idUser int64) {
	var (
		cif				int64
		nik				int64
		nama			string
		tempatLahir		string
		tanggalLahir	string
		alamat			string
		noTelepon		string
	)
	
	fmt.Println("\n=== Menu Manajemen Nasabah ===")
	fmt.Print("Masukan CIF/NIK\t:")
	fmt.Scan(&cif)
	response, err := PencarianCIF(cif)
	if err != nil {
		fmt.Println("Pencarian Gagal...")
		return
	}

	if response == -1 {
		fmt.Println("NIK atau CIF tidak ditemukan...")
		return
	}

	//Input data baru
	fmt.Println("\n=== Data Baru ===")
	fmt.Print("NIK\t\t:")
	fmt.Scan(&nik)
	fmt.Print("Nama\t\t:")
	fmt.Scan(&nama)
	fmt.Print("Tempat Lahir\t:")
	fmt.Scan(&tempatLahir)
	fmt.Print("Tanggal Lahir\t:")
	fmt.Scan(&tanggalLahir)
	fmt.Print("Alamat\t\t:")
	fmt.Scan(&alamat)
	fmt.Print("No Telepon\t:")
	fmt.Scan(&noTelepon)

	nasabah := bank.Nasabah{
		Cif				: cif,
		Nik				: nik,
		Nama			: nama,
		TempatLahir		: tempatLahir,
		TanggalLahir	: tanggalLahir,
		Alamat			: alamat,
		NoTelepon		: noTelepon,
	}

	//Melakukan proses update nasabah
	_, err2 := ManajemenNasabah(&nasabah)
	if err2 != nil {
		fmt.Println("Update nasabah gagal...")
		return
	}

	//Menampilkan data
	fmt.Println("\nUpdate nasabah berhasil!")
	fmt.Println("CIF\t\t:", nasabah.Cif)
	fmt.Println("NIK\t\t:", nasabah.Nik)
	fmt.Println("Nama\t\t:", nasabah.Nama)
	fmt.Println("Tempat Lahir\t:", nasabah.TempatLahir)
	fmt.Println("Tanggal Lahir\t:", nasabah.TanggalLahir)
	fmt.Println("Alamat\t\t:", nasabah.Alamat)
	fmt.Println("No Telepon\t:", nasabah.NoTelepon)
}