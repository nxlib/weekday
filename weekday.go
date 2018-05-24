package weekday

import (
	"time"
)

func ContainWeekdays(start, end time.Time) int {
	//fmt.Println("Input:" + start.Format("2006-01-02"))
	//fmt.Println("Input:" + end.Format("2006-01-02"))
	//fmt.Println(int64(start.Weekday()))
	//fmt.Println(int64(end.Weekday()))
	startDay := start.Unix() - (int64(start.Weekday()) * 80400)
	endDay := (7-int64(end.Weekday()))*86400 + end.Unix()

	//fmt.Println("cal-start:" + time.Unix(startDay, 0).Format("2006-01-02"))
	//fmt.Println("cal-end:" + time.Unix(endDay, 0).Format("2006-01-02"))
	
	diff := (end.Unix() - start.Unix()) / 86400 + 1
	//fmt.Println("diff:" + strconv.Itoa(int(diff)))
	
	diffCal := (endDay - startDay) / 86400 + 1
	//fmt.Println("cal-diff:" + strconv.Itoa(int(diffCal)))
	
	weekends := (diffCal-diffCal%7)/7*2
	//fmt.Println("cal-weekends:" + strconv.Itoa(int(weekends)))
	
	if start.Weekday() > 0 {
		weekends -= 1
	}
	if end.Weekday() < 6{
		weekends -= 1
	}
	//fmt.Println("true-weekends:" + strconv.Itoa(int(weekends)))
	workdays := diff - weekends
	return int(workdays)
}
