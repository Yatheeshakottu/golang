package main
import(
	"fmt"
)
func main(){
	str1:="BC"
	str2:="BANGLORE"
	fmt.Println(findDiff(str1,str2))
}
func findDiff(str1,str2 string)(string ,string){
	mp:=make(map[rune] bool)
	mp2:=make(map[rune]bool)
	res1:=""
	res2:=""
	for i := 0;i < len(str2);i++{
		mp[rune(str2[i])]=true
	}
	for j := 0;j < len(str1);j++{
		if_,ok:=mp[rune(str1[j])];
		!ok{
			res1 += string(str1[j])
		
		}
	}
	for i := 0;i < len(str1);i++{
		mp2[rune(str1[i])]=true
	}
	for j := 0;j < len(str2);j++{
		if _,ok:=mp2[rune(str2[j])];
		!ok{
			res2 += string(str2[j])
		}
	}
	return res1, res2
	}

