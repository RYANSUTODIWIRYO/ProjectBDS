package teller

import (
	"fmt"
	"golang.org/x/net/context"

	ent "bds/entities"
	conf "bds/config"
	bank "bds/proto"
)

func SetorTunai(transaksi ent.Transaksi) (ent.Transaksi, error) {
	//Koneksi ke grpc
	conn, err := conf.KoneksiGrpc()
	if err != nil {
		fmt.Println(err)
		return ent.Transaksi{}, err
	}
	defer conn.Close()

	//Memanggil funtcion SetorTunai() dari server
	s := bank.NewBankServiceClient(conn)
	response, err := s.SetorTunai(context.Background(), &bank.Transaksi{
		IdUser: transaksi.Id_user,
		NoRekening: transaksi.No_rekening,
		Tanggal: transaksi.Tanggal,
		JenisTransaksi: transaksi.Jenis_transaksi,
		Nominal: transaksi.Nominal,
		Saldo: transaksi.Saldo,
		Berita: transaksi.Berita,
	})

	//Memasukan nilai yang didapat
	if response.Saldo > 0 {
		transaksi.Saldo = response.Saldo
		return transaksi, nil
	} else {
		return ent.Transaksi{}, err
	}
}

func TarikTunai(transaksi ent.Transaksi) (ent.Transaksi, error) {
	//Koneksi ke grpc
	conn, err := conf.KoneksiGrpc()
	if err != nil {
		fmt.Println(err)
		return ent.Transaksi{}, err
	}
	defer conn.Close()

	//Memanggil funtcion TarikTunai() dari server
	s := bank.NewBankServiceClient(conn)
	response, err := s.TarikTunai(context.Background(), &bank.Transaksi{
		IdUser: transaksi.Id_user,
		NoRekening: transaksi.No_rekening,
		Tanggal: transaksi.Tanggal,
		JenisTransaksi: transaksi.Jenis_transaksi,
		Nominal: transaksi.Nominal,
		Saldo: transaksi.Saldo,
		Berita: transaksi.Berita,
	})

	//Memasukan nilai yang didapat
	if response.Saldo > 0 { //berhasil
		transaksi.Saldo = response.Saldo
		return transaksi, nil
	} else if response.Berita == "Saldo Tidak Cukup" { //berhasil, saldo tidak cukup
		return ent.Transaksi{Berita:response.Berita}, nil
	} else {
		return ent.Transaksi{}, err //gagal
	}
}

func CetakBuku(transaksi ent.Transaksi) (error) {
	//Koneksi ke grpc
	conn, err := conf.KoneksiGrpc()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer conn.Close()

	//Memanggil funtcion CetakBuku() dari server
	s := bank.NewBankServiceClient(conn)
	response, err := s.CetakBuku(context.Background(), &bank.Transaksi{
		NoRekening: transaksi.No_rekening,
	})	
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("No\tId\tTanggal\t\tJenis Transaksi\tNominal\tSaldo\tBerita")
	no := 1
	for _ , value := range response.Transaksi{
		fmt.Print(no, "\t")
    	fmt.Print(value.IdTransaksi, "\t")
		//fmt.Println(value.NoRekening)
		fmt.Print(value.Tanggal, "\t")
		fmt.Print(value.JenisTransaksi, "\t")
		fmt.Print(value.Nominal, "\t")
		fmt.Print(value.Saldo, "\t")
		fmt.Print(value.Berita, "\n")
		no++
	}
	return nil
}

func PindahBuku(transaksiPB ent.TransaksiPB) (ent.TransaksiPB, error) {
	//Koneksi ke grpc
	conn, err := conf.KoneksiGrpc()
	if err != nil {
		fmt.Println(err)
		return ent.TransaksiPB{}, err
	}
	defer conn.Close()

	//Memanggil funtcion TarikTunai() dari server
	s := bank.NewBankServiceClient(conn)
	response, err := s.PindahBuku(context.Background(), &bank.TransaksiPB{
		IdUser				: transaksiPB.Id_user,
		Tanggal				: transaksiPB.Tanggal,
		NasabahDebit		: &bank.NasabahDetail{NoRekening : transaksiPB.Nasabah_debit.No_rekening},
		NasabahKredit		: &bank.NasabahDetail{NoRekening : transaksiPB.Nasabah_kredit.No_rekening},
		Nominal: transaksiPB.Nominal,
		Berita: transaksiPB.Berita,
	})

	fmt.Println("respone", response.Nominal)

	//Memasukan nilai yang didapat
	if response.IdUser > 0 { //berhasil
		return transaksiPB, nil
	} else if response.Berita == "Saldo Tidak Cukup" { //berhasil, saldo tidak cukup
		return ent.TransaksiPB{Berita:response.Berita}, nil
	} else {
		return ent.TransaksiPB{}, err //gagal
	}
}

