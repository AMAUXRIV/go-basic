package main

import "fmt"

type Employee interface {
	Gaji() int
	Name() string
}

type Permanent struct {
	nama  string
	gapok int
	tunj int
}

type Contract struct {
	nama  string
	gapok int
}

type Freelancer struct {
	nama     string
	rateJam  int
	totalJam int
}

func (p Permanent) Gaji() int {
	return p.tunj + p.gapok
}

func (p Permanent) Name() string {
	return p.nama
}

func (c Contract) Gaji() int {
	return c.gapok
}

func (c Contract) Name() string {
	return c.nama
}

func (f Freelancer) Gaji() int {
	return f.rateJam * f.totalJam
}

func (f Freelancer) Name() string {
	return f.nama
}

func main() {
	employees := []Employee{
		Permanent{"A", 30000, 2000},
		Contract{"B", 30000},
		Freelancer{"amau", 100000, 40},
	}

	total_pengeluaran := 0

	for _, e := range employees {
		fmt.Println(e.Name(), e.Gaji())
		total_pengeluaran += e.Gaji()
	}

	fmt.Println("Total Pengeluaran Perusahaan:", total_pengeluaran)
}