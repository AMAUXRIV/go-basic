package main 

import (
	"fmt"
	"runtime"

)

func print(till int,message string) {
	for i := 0 ; i < till; i++ {
		fmt.Println ((i +1 ),message)
	}
}


func main() {
	runtime.GOMAXPROCS(2)
	go print(5,"halo apa kabar")
	print(5,"halo")

	var input string
	fmt.Scanln(&input)	
}