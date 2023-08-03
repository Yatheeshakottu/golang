/*
package main

import "fmt"

	func main() {
		var i, j, num int
		var add [10][10]int
		fmt.Print("enter the size of the matrix")
		fmt.Scan(&num)
		fmt.Print("enter the elements in to the matrix")
		for i = 0; i < num; i++ {
			for j = 0; j < num; j++ {
				fmt.Scan(&add[i][j])
			}
		}
		y := 1
		for i = 0; i < num; i++ {
			for j = 0; j < num; j++ {
				if add[i][j] != 1 && add[j][i] != 0 {
					y = 0
					break

				}
			}
		}

		if y == 1 {
			fmt.Println("it is a identity matrix")
		} else {
			fmt.Println("It is not a identity matrix")
		}
	}
*/
package main

import "fmt"

func main() {
	var num, i, j int

	var identMat [10][10]int

	fmt.Print("Enter the Matrix Size = ")
	fmt.Scan(&num)

	fmt.Print("Enter the Matrix Items = ")
	for i = 0; i < num; i++ {
		for j = 0; j < num; j++ {
			fmt.Scan(&identMat[i][j])
		}
	}
	flag := 1
	for i = 0; i < num; i++ {
		for j = 0; j < num; j++ {
			if identMat[i][j] != 1 && identMat[j][i] != 0 {
				flag = 0
				break
			}
		}
	}
	if flag == 1 {
		fmt.Println("It is an Idenetity Matrix")
	} else {
		fmt.Println("It is Not an Idenetity Matrix")
	}
}