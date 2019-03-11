package main
import (
    "fmt"
    "encoding/json"
)
type ColorGroup struct{
    id int
    name string
    colors []string
}


func main(){
	group:=ColorGroup{
		id:1,
		name:"reds",
		colors: []string{"crimson","red","ruby","maroon"},
	}

b,err:=json.Marshal(group)

fmt.Println(b,err)
}