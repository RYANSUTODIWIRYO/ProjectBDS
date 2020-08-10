package main

import (
	"fmt"
	"golang.org/x/net/context"

	ent "bds/entities"
	teller "bds/client/teller"
	conf "bds/config"
	bank "bds/proto"
)

func main()  {
	menu:
	//fmt.Println("\n=== Branch ===\n===  Delivery ===\n===   System\t===")
	
	fmt.Println("\n=== Branch ===")
	fmt.Println(" === Delivery ===")
	fmt.Println("  === System  =====")
	
	user, err := menuLogin()

	if err != nil {
		fmt.Println(err)
		goto menu
	}

	//Masuk menu sesuai role nya
	switch user.Role {
		case "teller":
			teller.MenuTeller(user)
		case "cs":
			// client.MenuCs()
		default:
			fmt.Println("Id User atau Password salah...")	
	}
	goto menu
}

func menuLogin() (ent.User, error){
	var (
		id_user	int64
		password string
	)

	//Input id user dan password
	fmt.Print("\n=== Menu Login ===\n")
	fmt.Print("Id User\t: ")
	fmt.Scan(&id_user)
	fmt.Print("Password: ")
	fmt.Scan(&password)

	//Memanggil function login()
	user, err := login(id_user, password)
	if err != nil {
		fmt.Println(err)
		return ent.User{}, err
	}
	return user, nil
}

func login(id_user int64, password string) (ent.User, error){

	//Koneksi ke grpc
	conn, err := conf.KoneksiGrpc()
	if err != nil {
		fmt.Println(err)
		return ent.User{}, err
	}
	defer conn.Close()

	//Memanggil funcion login() dari server
	service := bank.NewBankServiceClient(conn)
	response, err := service.Login(context.Background(), &bank.User{
		IdUser: id_user,
		Password: password,
	})

	//Memasukan nilai yang didapat
	user := ent.User{
		Id_user : response.IdUser,
		Password : response.Password,
		Nama_user : response.NamaUser,
		Role : response.Role,
		Cabang : response.Cabang,
	}

	return user, nil
}
