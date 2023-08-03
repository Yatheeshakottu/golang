package main

import "fmt"

func main() {
	var search = 9
	arr := []int{1, 2, 3, 67, 9, 843, 7}
	for k, v := range arr {
		if v == search {
			fmt.Println("we found a search value", search, "at index", k)
			break
		}
	}
}
/*func main(){
	x:=[]int{9,8,7,6,5,3}
	search:=7
	for i:=0;i<len(x);i++{
if arr[i]==search{
	fm.Println("we found the search value",search)
	
}
	}

}*/