package main


import "fmt"

type Config struct {

	Env string 
	Proxy Proxy
}


type Proxy struct {
	Address string
	Port string
}


func main () {

		c := &Config{
			Env : "DEBUG : TRUE",
			Proxy: Proxy{
				Address: "addr",
				Port: "port",
			},
		}

		fmt.Println ("Connect",c)
		fmt.Println("alamat",c.Proxy.Address)

}
