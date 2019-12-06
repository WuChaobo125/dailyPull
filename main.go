package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			// the year is a leap year if it is divisible by 400.
			if year%400 == 0 {
				return true
			} else {
				return false
			}
		} else {
			return true
		}

	} else {
		return false
	}
}

func dateGenerator(year int, cnt int) string {
	d2w := make(map[int]string)
	d2w[0] = "Mon"
	d2w[1] = "Tue"
	d2w[2] = "Wed"
	d2w[3] = "Thu"
	d2w[4] = "Fri"
	d2w[5] = "Sat"
	d2w[6] = "Sun"

	//m2d stores the dates of each month
	m2d := make(map[int]int)
	mon_30 := [4]int{4, 6, 9, 11}
	mon_31 := [7]int{1, 3, 5, 7, 8, 10, 12}

	for _, v := range mon_30 {
		m2d[v] = 30
	}

	for _, v := range mon_31 {
		m2d[v] = 31
	}

	if isLeapYear(year) {
		m2d[2] = 29
	} else {
		m2d[2] = 28
	}
	var b bytes.Buffer
	for m := 1; m <= 12; m++ {
		for d := 1; d <= m2d[m]; d++ {
			dateStr := strconv.Itoa(year) + "/" + strconv.Itoa(m) + "/" + strconv.Itoa(d) + " (" + d2w[cnt%7] + ")\r\n\r\n"
			for t := 6; t <= 11; t++ {
				dateStr += strconv.Itoa(t) + ":00AM\r\n\r\n"
			}
			dateStr += "12:00PM\r\n\r\n"
			for t := 1; t <= 11; t++ {
				dateStr += strconv.Itoa(t) + ":00PM\r\n\r\n"
			}
			// dateStr += "To do list:\r\n"
			// for t := 1; t <= 10; t++ {
			// 	dateStr += strconv.Itoa(t)
			// 	dateStr += ".\r\n"
			// }
			dateStr += "\r\n"
			b.WriteString(dateStr)
			cnt++
		}
	}
	return b.String()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var year, fday int
	var year2str, fday2str string
	for {
		fmt.Println("Enter year:")
		fmt.Print("->")
		year2str, _ = reader.ReadString('\n')
		year2str = strings.Replace(year2str, "\r\n", "", -1)
		year, _ = strconv.Atoi(year2str)
		if year >= 0 {
			break
		} else {
			fmt.Println("Error: Year must be a positive integer!\r\n")
		}
	}

	for {
		fmt.Printf("Enter the first day of %s (1 for Mon, 2 for Tue...) \r\n", year2str)
		fmt.Print("->")
		fday2str, _ = reader.ReadString('\n')
		fday2str = strings.Replace(fday2str, "\r\n", "", -1)
		fday, _ = strconv.Atoi(fday2str)
		if fday > 0 && fday < 8 {
			break
		} else {
			fmt.Println("Error: Input must be between 1 to 7 (1 for Mon, 2 for Tue...)\r\n")
		}
	}

	userFile := year2str + ".txt"

	fout, err := os.Create(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fout.Close()

	fout.WriteString(dateGenerator(year, fday-1))
	fmt.Print("Your note is generated.")
	return

}
