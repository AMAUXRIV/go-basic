package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue //
		}
		fmt.Println(n)
	}
	// membuat segitiga bintang *
	var b int
	var c int
	for b = 1; b <= 5; b++ {
		for c = 1; c <= b; c++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func rangE() {
	fmt.Println("==============")
	nama := []string{"eko", "amau","amu","asu"}
	for i,value := range nama {
		if i == 1 {
			continue
			// break
		}else {
			fmt.Println("index",i,"nama",value)
		}
	}

}
