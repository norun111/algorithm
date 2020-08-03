package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	list := strings.Split(sc.Text(), " ")
	a,_ := strconv.Atoi(list[0])
	b, _ := strconv.Atoi(list[1])
	c, _ := strconv.Atoi(list[2])

	sc.Scan()
	trainNum, _ := strconv.Atoi(sc.Text())

	digitList := []int{}

	li := []string{}

	for i:=0; i<trainNum; i++ {
		sc.Scan()
		li = append(li, sc.Text())
		timeList := strings.Split(sc.Text(), " ")
		hour, _ := strconv.Atoi(timeList[0])
		minute, _ := strconv.Atoi(timeList[1])

		minute-=a
		if minute < 0 {
			hour--
			minute = 60 + minute
		}
		minute+=a
		if minute >= 60 {
			minute-=60
			hour++
		}
		minute+=b
		if minute >= 60 {
			minute-=60
			hour++
		}
		minute+=c
		if minute >= 60 {
			minute-=60
			hour++
		}

		if hour <= 8 && minute <= 59 {
			digitList = append(digitList, i)
		}
	}

	correctDigit := digitList[len(digitList)-1]

	correctTime := li[correctDigit]

	newcorrectTime := strings.Split(correctTime, " ")
	cHour, _ := strconv.Atoi(newcorrectTime[0])
	cMinute, _ := strconv.Atoi(newcorrectTime[1])

	cMinute-=a
	if cMinute < 0 {
		cHour--
		cMinute = 60 + cMinute
	}

	strHour := strconv.Itoa(cHour)
	strMinute := strconv.Itoa(cMinute)

	if len(strMinute) == 2 {
		fmt.Println("0"+strHour+":"+strMinute)
	} else {
		fmt.Println("0"+strHour+":"+"0"+strMinute)
	}

}