package main

import "fmt"

type Halo interface {
	login() (string, error)
	logout() error
}

type sapa struct {
	Nama string
}

func (s *sapa) login() (string, error) {
	if s == nil {
		return "", fmt.Errorf("Kosong")
	}
	return fmt.Sprintf("halo %s", s.Nama), nil
}

func (s *sapa) logout() error {
	// implementation of logout method
	return nil
}


func NewMap (nama string) (map[string]string){
	if nama == "" {
		return nil
	} else {
		return map[string]string{
			"nama" : nama,
		}
	}

}

func main() {
	// code to test the sapa struct and Halo interface
	var helo Halo

	helo = &sapa{Nama: "Ujang"}

	// see a different between the two this step 
	fmt.Println (helo.login())
	fmt.Println("=============")
	greeting, err := helo.login()
	if err != nil {
		fmt.Println(err)

	}

	fmt.Println(greeting)
	//===============

	var coba = NewMap("Ujang")
	fmt.Println(coba)


	//cek data

	var data = NewMap("eko")
	if data == nil {
		fmt.Println("Data Kosong")

	}else {
		fmt.Println(data)
	}
	
}

