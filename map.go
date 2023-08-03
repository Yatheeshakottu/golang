package main

 
import(
	"fmt"

)  
func main(){
	m:=map[string]int{
		"james":           42,
		"miss moneypenny": 32,
	}
	fmt.Println(m)
	fmt.Println(m["james"])
   v,ok := (m["barnamas"])
fmt.Println(v)
	fmt.Println(ok)  
	if v,ok:=m["miss moneypenny"];ok{ 
		fmt.Println("this is the print ",v)
}


}