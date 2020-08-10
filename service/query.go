package service

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	// . "github.com/ragilmaulana/restapi/tugas-golang/BranchDeliverySistem/entities"
	. "bds/entities"
	bank "bds/proto"
	//"time"
)

type UserService struct {
	DB *sql.DB
}

func (us UserService) LoginUser(id_user int64, password string) (User, error) {
	rows, err := us.DB.Query("SELECT * FROM user WHERE id_user = ? AND password = ? ",id_user, password)
	if err != nil {
		return User{}, err
	} else {
		var user User
		for rows.Next() {
			var id_user int
			var password string
			var nama string
			var role string
			var cabang string
			err2 := rows.Scan(&id_user, &password, &nama, &role, &cabang)
			if err2 != nil {
				return User{}, err
			} else {
				user = User{
					Id_user: int64 (id_user),
					Password: password,
					Nama_user: nama,
					Role: role,
					Cabang: cabang,
				}
			}
		}
		return user, nil
	}
}

func (us UserService) CariNasabah(rekeningTujuan int64) (NasabahDetail, error) {
	//rows, err := us.DB.Query("SELECT * FROM user WHERE id_user = ? AND password = ? ",id_user, password)
	rows, err := us.DB.Query(
		"SELECT nasabah_detail.cif, nasabah.nama, nasabah_detail.no_rekening, nasabah_detail.saldo " +
		"FROM bank.nasabah_detail " +
		"INNER JOIN bank.nasabah " +
		"ON (nasabah_detail.cif = nasabah.cif AND nasabah_detail.no_rekening = ?)", rekeningTujuan)

		if err != nil {
		return NasabahDetail{}, err
	} else {
		var nasabahDetail NasabahDetail
		for rows.Next() {
			var (
				cif			int
				nama		string
				no_rekening int
				saldo	 	int
			)
			err2 := rows.Scan(&cif, &nama, &no_rekening, &saldo)
			if err2 != nil {
				return NasabahDetail{}, err
			} else {
				nasabahDetail = NasabahDetail{
					Cif: 			int64 (cif),
					Nama: 			nama,
					No_rekening:	int64 (no_rekening),
					Saldo: 			int64 (saldo),
				}
			}
		}
		return nasabahDetail, nil
	}
}

// insert setor tunai to db
func (us UserService) SetorTunaiService(transaksi Transaksi, nasabah NasabahDetail) (int32, Transaksi, error) {
	// tanggal := time.Now().Format("2006-01-02 15:04:05")
	// saldo := nasabah.Saldo
	// jenisTransaksi := "st"
	// check saldo overload

	nasabah.Saldo = nasabah.Saldo + transaksi.Nominal
	rows, err := us.DB.Exec(
		"INSERT INTO transaksi (id_user, no_rekening, tanggal, jenis_transaksi, nominal, saldo, berita) values (?,?,?,?,?,?,?)",
		transaksi.Id_user,
		nasabah.No_rekening,
		transaksi.Tanggal,
		transaksi.Jenis_transaksi,
		transaksi.Nominal,
		nasabah.Saldo,
		transaksi.Berita)
	_, err = us.DB.Exec(
		"UPDATE nasabah_detail SET saldo = ? where no_rekening = ?", nasabah.Saldo, nasabah.No_rekening)
	if err != nil {
		return 0, Transaksi{}, err
	} else {
		// status, _ := rows.RowsAffected()
		fmt.Println("nihhhhh", transaksi)
		transaksi.Saldo = nasabah.Saldo
		rows.RowsAffected()
		return 1, transaksi, nil
	}
}

// insert setor tunai to db
func (us UserService) TarikTunaiService(transaksi Transaksi, nasabah NasabahDetail) (int32, Transaksi, error) {
	// tanggal := time.Now().Format("2006-01-02 15:04:05")
	// saldo := nasabah.Saldo
	// jenisTransaksi := "st"
	// check saldo overload

	nasabah.Saldo = nasabah.Saldo - transaksi.Nominal
	if nasabah.Saldo < 0 {
		return -1, Transaksi{}, nil
	}

	rows, err := us.DB.Exec(
		"INSERT INTO transaksi (id_user, no_rekening, tanggal, jenis_transaksi, nominal, saldo, berita) values (?,?,?,?,?,?,?)",
		transaksi.Id_user,
		nasabah.No_rekening,
		transaksi.Tanggal,
		transaksi.Jenis_transaksi,
		transaksi.Nominal,
		nasabah.Saldo,
		transaksi.Berita)
	_, err = us.DB.Exec(
		"UPDATE nasabah_detail SET saldo = ? where no_rekening = ?", nasabah.Saldo, nasabah.No_rekening)
	if err != nil {
		return 0, Transaksi{}, err
	} else {
		// status, _ := rows.RowsAffected()
		fmt.Println("nihhhhh", transaksi)
		transaksi.Saldo = nasabah.Saldo
		rows.RowsAffected()
		return 1, transaksi, nil
	}
}

func (us UserService) CetakBuku(no_rekening int) ([]Transaksi, error) {
	rows, err := us.DB.Query("SELECT * FROM transaksi WHERE no_rekening = ?", no_rekening)
	
	if err != nil {
		return []Transaksi{}, err
	} else {
		var transaksi []Transaksi
		for rows.Next() {
			var id_transaksi 	int
			var id_user 		int
			var no_rekening		int
			var tanggal 		string
			var jenis_transaksi string
			var nominal 		float64
			var saldo 			float64
			var berita 			string
			err2 := rows.Scan(&id_transaksi, &id_user, &no_rekening, &tanggal, &jenis_transaksi, &nominal, &saldo, &berita)
			if err2 != nil {	
				return []Transaksi{}, err2
			} else {
				trx := Transaksi{
					Id_transaksi:		int64 (id_transaksi),
					Id_user:      		int64 (id_user),
					No_rekening: 		int64 (no_rekening),
					Tanggal:         	tanggal,
					Jenis_transaksi:	jenis_transaksi,
					Nominal:         	int64 (nominal),
					Saldo:           	int64 (saldo),
					Berita:          	berita,
				}
				transaksi = append(transaksi, trx)
			}
		}
		return transaksi, nil
	}
}

func (us UserService) PindahBukuService(idUser int64, tanggal string, rekeningAwal, rekeingTujuan NasabahDetail, nominal int64, berita string) (int64, error) {
	// check saldo overload
	if rekeningAwal.Saldo < nominal {
		return 0, nil
	} else {
		saldoRekAwal := int(rekeningAwal.Saldo - nominal)
		//fmt.Println(saldoRekAwal)
		saldoRekTujuan := int(rekeingTujuan.Saldo + nominal)
		//fmt.Println(saldoRekTujuan)
		_, err := us.DB.Exec(
			"INSERT INTO transaksi (id_user, no_rekening, tanggal, jenis_transaksi, nominal, saldo, berita) VALUES (?,?,?,?,?,?,?)",
			idUser, rekeningAwal.No_rekening, tanggal, "pb (d)", nominal, saldoRekAwal, berita)
		if err != nil {
			panic(err)
		}
		//fmt.Println(insertRekUtama)
		_, err2 := us.DB.Exec("UPDATE nasabah_detail SET saldo = ? WHERE no_rekening = ?", saldoRekAwal, rekeningAwal.No_rekening)
		if err2 != nil {
			panic(err)
		}
		//fmt.Println(update)
		_, err3 := us.DB.Exec("INSERT INTO transaksi (id_user, no_rekening, tanggal, jenis_transaksi, nominal, saldo, berita) VALUES (?,?,?,?,?,?,?)",
			idUser, rekeingTujuan.No_rekening, tanggal, "pb (k)", nominal, saldoRekTujuan, berita)
		if err3 != nil {
			panic(err)
		}
		//fmt.Println(insertRekKedua)
		_, err4 := us.DB.Exec("UPDATE nasabah_detail SET saldo = ? where no_rekening = ?", saldoRekTujuan, rekeingTujuan.No_rekening)
		if err4 != nil {
			panic(err)
		}
		//fmt.Println(updateRekTujuan)

	}
	return idUser, nil
}

func (us UserService) FindByCifOrNikService(cif int64) (Nasabah, error) {
	//Mencari nasabah berdasarkan cif atau nik
	rows, err := us.DB.Query("SELECT * FROM nasabah WHERE cif = ? OR nik = ?", cif, cif)
	if err != nil {
		panic(err)
	} else {
		var nasabah Nasabah
		for rows.Next() {
			var (
				cif int64
				nik int64
				nama string
				tempat_lahir string
				tanggal_lahir string
				alamat string
				no_telepon string
			)

			//Menampung hasil query
			err2 := rows.Scan(&cif, &nik, &nama, &tempat_lahir, &tanggal_lahir, &alamat, &no_telepon)
			if err2 != nil {
				panic(err2)
			}

			nasabah = Nasabah{
				Cif:           cif,
				Nik:           nik,
				Nama:          nama,
				Tempat_lahir:  tempat_lahir,
				Tanggal_lahir: tanggal_lahir,
				Alamat:        alamat,
				No_telepon:    no_telepon,
			}
			//fmt.Println("called")
			//fmt.Println(nasabah)
		}
		return nasabah, nil
	}
}

func (us UserService) BuatCifService(nasabah *bank.Nasabah) (*bank.Nasabah, error) {
	//Memasukan data ke table nasabah
	rows, err := us.DB.Exec("INSERT INTO nasabah (nik,nama,tempat_lahir,tanggal_lahir,alamat,no_telepon) values (?,?,?,?,?,?)",
		nasabah.Nik, nasabah.Nama, nasabah.TempatLahir, nasabah.TanggalLahir, nasabah.Alamat, nasabah.NoTelepon)
	//fmt.Println("nik", nasabah.Nik)
	if err != nil {
		panic(err)
	} else {
		status, _ := rows.RowsAffected()
		if status > 0 {
			response, err := us.FindByCifOrNikService(nasabah.Nik)
			nasabah.Cif = response.Cif
			return nasabah, err
		} else {
			return &bank.Nasabah{}, err
		}
	}
}

func (us UserService) FindLastRekService() (int, error) {
	rows, err := us.DB.Query("SELECT no_rekening FROM nasabah_detail ORDER BY no_rekening DESC LIMIT 1")
	if err != nil {
		panic(err)
	} else {
		// looping data
		var no_rekening int
		for rows.Next() {

			err2 := rows.Scan(&no_rekening)
			if err2 != nil {
				panic(err2)
			}
		}
		return no_rekening, nil
	}
}

func (us UserService) BuatTabunganService(nasabah *bank.NasabahDetail) (*bank.NasabahDetail, error) {
	//Mencari nomor rekening yang dimasukan terakhir
	last_no_rekening, _ := us.FindLastRekService()
		last_no_rekening += 1

	//Memasukan ke database nasabah_detail
	//fmt.Println(last_no_rekening)
	rows, err := us.DB.Exec("INSERT INTO nasabah_detail (cif, no_rekening, saldo) VALUES (?,?,?)",
		nasabah.Cif, last_no_rekening, nasabah.Saldo)
	if err != nil {
		panic(err)
	}

	//Mencari nama dari method FindByCifOrNikService()
	status, _ := rows.RowsAffected()
	if status > 0 {
		// call method print nasabah by cif
		response, _ := us.FindByCifOrNikService(nasabah.Cif)
		nasabah.Nama = response.Nama
		nasabah.NoRekening = int64 (last_no_rekening)
		return nasabah, nil
	} else {
		return &bank.NasabahDetail{}, nil
	}
}

func (us UserService) UpdateNasabahService(nasabah *bank.Nasabah) (*bank.Nasabah, error) {
	
	rows, err := us.DB.Exec("UPDATE nasabah SET nik = ?, nama = ?, tempat_lahir = ?, tanggal_lahir = ?, alamat = ?, no_telepon = ? WHERE cif = ?",
	nasabah.Nik, nasabah.Nama, nasabah.TempatLahir, nasabah.TanggalLahir, nasabah.Alamat, nasabah.NoTelepon, nasabah.Cif)
	// fmt.Println(nasabah)
	if err != nil {
		panic(err)
	}
	
	//var Respons *BranchDeliverySystem.NASABAH_INFO
	status, _ := rows.RowsAffected()
	//fmt.Println(idUser)
	if status > 0 {
		return nasabah, nil
	} else {
		return &bank.Nasabah{}, nil
	}
}
