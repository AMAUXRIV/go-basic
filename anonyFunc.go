package main 


import "fmt"



type Blacklist func(string)bool

func registerBlacklist(name string, blacklist Blacklist){
	if blacklist(name){
		fmt.Println("You Blocked Bitch")
	}else {
		fmt.Println("Welcome",name)
	}
}


func main () {

	//Anonymous FUnction ialah Fungsi Tanpa nama
	blacklist := func(name string) bool{
		return name == "admin"
	}
	
	registerBlacklist("admin", blacklist)
}