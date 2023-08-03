package main
import(
	"fmt"
)
func main(){
	sample:="Yathee"
	samplerune:=[]rune(sample)
	Generatepermutation:=(samplerune, 0,len(samplerune)-1)
}
func Generatepermutation(samplerune []rune,left,right int){
	if left == right{
		fmt.Println(string(samplerune))
	}else{
		for i := left;i
	 }
} 