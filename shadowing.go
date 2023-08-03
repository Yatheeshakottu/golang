package main
import(
	"fmt"
)
var numOfCars = 2    // Line 1
type Car struct{
	name string
	model string
	color string
}
 func main(){
car := [{
            name:"Toyota",
            model:"Corolla",
            color:"red"
        },
        {
	    name:"Toyota",
            model:"Innova",
            color:"gray"
	}]
								
								
func countRedCars(){
    for i:=0; i<numOfCars; i++{
	if cars[i].color == "red" {
		numOfCars +=1    // Line 2
		fmt.Println("Inside countRedCars method ", numOfCars)    //Line 3
	}
    }	    
}