/*package main

import "fmt"
//Fungsi yang dapat disimpan dalam variabel itulah closure dan dia bekerja di dalam scope pribadi dan bisa mengubah diluar scope nya
// akan tetapi var yang ada di dalamnya tidak dapat di akses diluar scope pribadi 
// jika ada kesamaan antar variabel diluar scope pribadi maka harus di deklarasi ulang (cek yt pzn :closure)
// berupa anonymous func yakni tanpa nama
func main () {
	var getMinMax = func (n  []int) (int,int){
		var min,max int
		for i,value := range n {
			switch {
			case i == 0 :
					max, min = value,value
			case value > max :
				max = value
			case value < min :
				min = value
			}
			
		
		}
		return min,max
	}

	var numbers = []int{1,2,3,4,5,6,7}
	var min,max = getMinMax(numbers)
	fmt.Printf("data :%v\nmax:%v\nmin:%v\n",numbers,max,min)
}*/


/*
Immediately-Invoked Function Expression (IIFE)
Closure jenis ini dieksekusi langsung pada saat deklarasinya. 
Biasa digunakan untuk membungkus proses yang hanya dilakukan sekali, bisa mengembalikan nilai, bisa juga tidak.
*/


// IIFE function
package main

import "fmt"

func main() {
    var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}

    var newNumbers = func(min int) []int /*ini diisi oleh r*/ {
        var r []int
        for _, e := range numbers {
            if e < min {
                continue
            }
            r = append(r, e)
        }
        return r
    }(2) // ini adalh paramater suatu nilai balikan yang ditampung nantinya oleh newNumbers

    fmt.Println("original number :", numbers)
    fmt.Println("filtered number :", newNumbers)
}