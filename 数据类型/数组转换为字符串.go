package main
import (
	"fmt"
	"strings"
)
func main(){
	array:=[]string{"a","A","b","B","c","C"}
	fmt.Println(array)
	trim:=strings.Trim(fmt.Sprint(array),"[]")
	fmt.Println(trim)
	str:=strings.Replace(trim," ",",",-1)
	fmt.Println(str)
}