package main 

import "fmt"

type Gaji interface {
	gaji() int
}

type pegawaiTetap struct {
	nama string
	gapok int
	tunj int
}

type pegawaiKontrak struct {
	nama string
	gapok int
}

type freelancer struct {
	nama string
	rateJam int
	totalJam int
	
}
func (pt pegawaiTetap) gaji() int{
	return pt.tunj + pt.gapok
}

func (pk pegawaiKontrak) gaji() int {
	return pk.gapok
}

func (fl freelancer) gaji() int {
	return fl.rateJam * fl.totalJam

}

func main () {
	total_pengeluaran := 0
	pk1 := pegawaiTetap{"A",30000,2000}
	fmt.Println(pk1.nama,pk1.gaji())
	pt1 := pegawaiKontrak{"B",30000}
	fmt.Println(pt1.nama,pt1.gaji())
	total_pengeluaran += pt1.gaji()
	free := freelancer{"amau",100000,40}
	fmt.Println(free.nama,free.gaji())
	total_pengeluaran += free.gaji()

	fmt.Println("Total Pengeluaran Perusahaan",total_pengeluaran)

}