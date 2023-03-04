package main

import "fmt"

type Gaji interface {
	gaji() int
}

type pegawaiTetap struct {
	nama  string
	gapok int
	tunj  int
}

type pegawaiKontrak struct {
	nama  string
	gapok int
}

type freelancer struct {
	nama     string
	rateJam  int
	totalJam int
}

func (pt pegawaiTetap) gaji() int {
	return pt.tunj + pt.gapok
}

func (pk pegawaiKontrak) gaji() int {
	return pk.gapok
}

func (fl freelancer) gaji() int {
	return fl.rateJam * fl.totalJam

}

func main() {
	total_pengeluaran := 0
	pk1 := pegawaiTetap{"A", 30000, 2000}
	fmt.Println(pk1.nama, pk1.gaji())package main

import "fmt"

type Employee interface {
    Salary() int
}

type Permanent struct {
    Name      string
    BasicPay  int
    Allowance int
}

type Contract struct {
    Name     string
    BasicPay int
}

type Freelance struct {
    Name      string
    HourlyPay int
    Hours     int
}

func (p Permanent) Salary() int {
    return p.BasicPay + p.Allowance
}

func (c Contract) Salary() int {
    return c.BasicPay
}

func (f Freelance) Salary() int {
    return f.HourlyPay * f.Hours
}

func main() {
    employees := []Employee{
        Permanent{Name: "A", BasicPay: 30000, Allowance: 2000},
        Contract{Name: "B", BasicPay: 30000},
        Freelance{Name: "C", HourlyPay: 100000, Hours: 40},
    }

    totalSalary := 0
    for _, e := range employees {
        salary := e.Salary()
        fmt.Printf("%s: %d\n", e.(interface{ Name() string }).Name(), salary)
        totalSalary += salary
    }

    fmt.Printf("Total Pengeluaran Perusahaan: %d\n", totalSalary)
}

	pt1 := pegawaiKontrak{"B", 30000}
	fmt.Println(pt1.nama, pt1.gaji())
	total_pengeluaran += pt1.gaji()
	free := freelancer{"amau", 100000, 40}
	fmt.Println(free.nama, free.gaji())
	total_pengeluaran += free.gaji()

	fmt.Println("Total Pengeluaran Perusahaan", total_pengeluaran)

}
