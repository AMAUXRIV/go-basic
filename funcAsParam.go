package main 

import (
	"fmt"
	"strings"
)

func filter (data []string,callback func(string) bool) []string{
	var result []string
	for _,each := range data {
		if filtered := callback(each); filtered {
			result = append(result,each)
		}
	}
	return result
}
/*
Fungsi filter di atas digunakan untuk memfilter elemen-elemen pada sebuah slice data berdasarkan kondisi tertentu 
yang diberikan oleh fungsi callback. Fungsi callback merupakan sebuah fungsi yang diberikan sebagai parameter pada fungsi filter 
dan akan dipanggil untuk setiap elemen pada slice data. Fungsi callback akan mengembalikan nilai true atau false untuk menunjukkan 
apakah sebuah elemen harus disertakan dalam hasil filter atau tidak.
*/

func main (){
	var data = []string {"budi","joko","udin","anjing"}
	var dataO= filter(data,func(each string) bool {
		return strings.Contains(each,"o")
	})

	var datalen5 = filter(data,func(each string) bool {
		return len(each)== 4
	})

	var jorok = filter(data,func(each string) bool {
		
		return strings.Contains(each,"anjing") 
	})
	fmt.Println("Data Yang berisikan O",dataO)
	fmt.Println("Data Yang Jorok",jorok)
	fmt.Println("Filter jumlah huruf 4",datalen5)
}