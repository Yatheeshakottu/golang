package main
import(
 "fmt"
)
func main(){
 var  s interface{}="yatheesha"
 a:=s.(string)
 fmt.Println(a)
 a,ok:=s.(string)
 fmt.Println(a,ok)
 b,ok:=s.(float64)
 fmt.Println(b,ok)
 c:=s.(int)
 fmt.Println(c)
}