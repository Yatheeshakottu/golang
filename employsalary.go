package main

import (
	"fmt"
)

func main() {
	var grossSA, HRA, Basicsalary, DA float64
	fmt.Println("Enter the employ Basicsalary:")
	fmt.Scan(&Basicsalary)
	fmt.Print("Enter the employ HRA salary:")
	fmt.Scan(&HRA)
	fmt.Print("Enter the employ DA salary:")
	fmt.Scan(&DA)
	grossSA = Basicsalary + HRA + DA
	fmt.Println("The employ gross salary:", grossSA)
}
