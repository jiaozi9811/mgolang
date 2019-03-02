package main
import (
	//"bytes"
	//"io"
	//"encoding/json"
	//"net/http"
    //"os"
    "fmt"
    "sort"
)


func main(){
		numbers := []int{-2, -1, 0, 6, 2}
		sort.Sort(sort.IntSlice(numbers))
		fmt.Println("IsSorted:", numbers)
}