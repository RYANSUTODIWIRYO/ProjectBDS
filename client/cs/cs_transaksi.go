package cs

import (
	"fmt"
	"golang.org/x/net/context"

	// ent "bds/entities"
	conf "bds/config"
	bank "bds/proto"
)

func PencarianCIF(cif int64) (int, error){
	//Koneksi ke grpc
	conn, err := conf.KoneksiGrpc()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer conn.Close()

	//Memanggil funtcion findByNIKOrCIF() dari server
	s := bank.NewBankServiceClient(conn)
	nasabah, err2 := s.FindByCifOrNik(context.Background(), &bank.Nasabah{
		Cif: cif,
	})	
	if err2 != nil {
		fmt.Println(err)
		return 0, err
	}

	if nasabah.Cif == 0 {
		//fmt.Println("NIK atau CIF tidak ditemukan...")
		return -1, nil
	}

	fmt.Println("\nCif\t\t:", nasabah.Cif)
	fmt.Println("Nik\t\t:", nasabah.Nik)
	fmt.Println("Nama\t\t:", nasabah.Nama)
	fmt.Println("Tempat Lahir\t:", nasabah.TempatLahir)
	fmt.Println("Tanggal Lahir\t:", nasabah.TanggalLahir)
	fmt.Println("Alamat\t\t:", nasabah.Alamat)
	fmt.Println("NoTelepon\t:", nasabah.NoTelepon)
	
	return 1, nil
}

func PembuatanCif(nasabah *bank.Nasabah) (*bank.Nasabah, error){
	//Koneksi ke grpc
	conn, err := conf.KoneksiGrpc()
	if err != nil {
		fmt.Println(err)
		return &bank.Nasabah{}, err
	}
	defer conn.Close()

	//Memanggil funtcion findByNIKOrCIF() dari server
	s := bank.NewBankServiceClient(conn)
	response, err2 := s.BuatCif(context.Background(), nasabah)	
	if err2 != nil {
		fmt.Println(err)
		return &bank.Nasabah{}, err
	}

	return response, nil
}

func PembuatanTabungan(nasabah *bank.NasabahDetail) (*bank.NasabahDetail, error){
	//Koneksi ke grpc
	conn, err := conf.KoneksiGrpc()
	if err != nil {
		fmt.Println(err)
		return &bank.NasabahDetail{}, err
	}
	defer conn.Close()

	//Memanggil funtcion BuatTabungan() dari server
	s := bank.NewBankServiceClient(conn)
	response, err2 := s.BuatTabungan(context.Background(), nasabah)	
	if err2 != nil {
		fmt.Println(err)
		return &bank.NasabahDetail{}, err
	}

	return response, nil
}

func ManajemenNasabah(nasabah *bank.Nasabah) (*bank.Nasabah, error) {
	//Koneksi ke grpc
	conn, err := conf.KoneksiGrpc()
	if err != nil {
		fmt.Println(err)
		return &bank.Nasabah{}, err
	}
	defer conn.Close()

	//Memanggil funtcion UpdateTabungan() dari server
	s := bank.NewBankServiceClient(conn)
	response, err2 := s.UpdateNasabah(context.Background(), nasabah)	
	if err2 != nil {
		fmt.Println(err)
		return &bank.Nasabah{}, err
	}

	return response, nil
}

