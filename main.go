package main

import "fmt"

func main() {
	var lsn float32 = 4
	var lnn float32 = lsn / 365
	var start float32 = 1000000
	var ans float32 = start

	for i := 0; i < 365; i++ {
		ln := ((ans * lnn) / 100)
		fmt.Println("Loi nhuan ngay ", i, ": ", ln)
		ans += ln
	}
	fmt.Println("Tien cuoi: ", ans)
	fmt.Println("Loi nhuan tren nam %: ", ((ans-start)/start)*100)
	fmt.Println("Tien moi ngay: ", ((ans-start)/365)*100)
}
