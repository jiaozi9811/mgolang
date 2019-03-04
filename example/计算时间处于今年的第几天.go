//计算时间处于今年的第几天
package main
import (
	"fmt"
	"time"
)
func main(){
fmt.Println(GetDaysOfYear(time.Now()))
}

func GetDaysOfYear(t time.Time) (days int) {
	year := t.Year()
	mont := int(t.Month())
	day := t.Day()
	feb := 28
	if (year % 100 != 0 && year % 4 == 0) || year % 400 == 0 {
		feb = 29
	}
	monthDays := []int{31, feb, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	for index, monthDays := range monthDays {
		if index < mont - 1  {
			days += monthDays
		}
	}
	return days + day
}