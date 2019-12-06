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
			}
			return false
		}
		return true
	}
	return false
}

//m2d stores the dates of each month
var m2d = make(map[int]int)
var mon30 = [4]int{4, 6, 9, 11}
var mon31 = [7]int{1, 3, 5, 7, 8, 10, 12}

func isValidatedDate(y int, m int, d int) bool {
	if y > 1752 && m < 13 && m > 0 {
		if d < 0 || d > m2d[m] {
			return false
		}
		return true
	}
	return false
}

//only work for year after 1752
//0 = Sunday, 1 = monday ...
func dayOfWeek(y int, m int, d int) int {
	t := [12]int{0, 3, 2, 5, 0, 3, 5, 1, 4, 6, 2, 4}
	if m < 3 {
		y--
	}
	return (y + y/4 - y/100 + y/400 + t[m-1] + d) % 7
}

//0 for todo list
//1 for time-based diary
func contentOfDay(option int) string {
	var output string
	if option == 0 {
		output += "To do list:\r\n"
		for t := 1; t <= 10; t++ {
			output += strconv.Itoa(t)
			output += ".\r\n"
		}
	}
	if option == 1 {
		for t := 6; t <= 11; t++ {
			output += strconv.Itoa(t) + ":00AM\r\n\r\n"
		}
		output += "12:00PM\r\n\r\n"
		for t := 1; t <= 11; t++ {
			output += strconv.Itoa(t) + ":00PM\r\n\r\n"
		}
	}
	return output
}

func dateGenerator(year int, month int, day int, period int, option int) string {
	cnt := dayOfWeek(year, month, day)

	d2w := make(map[int]string)
	d2w[0] = "Sun"
	d2w[1] = "Mon"
	d2w[2] = "Tue"
	d2w[3] = "Wed"
	d2w[4] = "Thu"
	d2w[5] = "Fri"
	d2w[6] = "Sat"

	//generate daily contents based on option
	contents := contentOfDay(option)

	var b bytes.Buffer

	for d := day; d <= m2d[month] && period > 0; d++ {
		dateStr := strconv.Itoa(year) + "/" + strconv.Itoa(month) + "/" + strconv.Itoa(d) + " (" + d2w[cnt%7] + ")\r\n\r\n"
		dateStr += contents
		dateStr += "\r\n"
		b.WriteString(dateStr)
		cnt++
		period--
	}

	for m := month + 1; m <= 12; m++ {
		for d := 1; d <= m2d[m] && period > 0; d++ {
			dateStr := strconv.Itoa(year) + "/" + strconv.Itoa(m) + "/" + strconv.Itoa(d) + " (" + d2w[cnt%7] + ")\r\n\r\n"
			dateStr += contents
			dateStr += "\r\n"
			b.WriteString(dateStr)
			cnt++
			period--
		}
	}
	return b.String()
}

func main() {
	for _, v := range mon30 {
		m2d[v] = 30
	}
	for _, v := range mon31 {
		m2d[v] = 31
	}

	reader := bufio.NewReader(os.Stdin)
	var year, month, day, period, option int
	var year2str, month2str, day2str, period2str, option2str string
	fmt.Println("<-dummy diary->")
	fmt.Println("--------------------------")
	for {
		fmt.Println("Enter option: 0 for todolist; 1 for time based diary")
		fmt.Print("->")
		option2str, _ = reader.ReadString('\n')
		option2str = strings.Replace(option2str, "\r\n", "", -1)
		option, _ = strconv.Atoi(option2str)
		if option == 0 || option == 1 {
			break
		}
	}

	for {
		fmt.Println("Enter year:")
		fmt.Print("->")
		year2str, _ = reader.ReadString('\n')
		year2str = strings.Replace(year2str, "\r\n", "", -1)
		year, _ = strconv.Atoi(year2str)

		if isLeapYear(year) {
			m2d[2] = 29
		} else {
			m2d[2] = 28
		}

		fmt.Println("Enter month:")
		fmt.Print("->")
		month2str, _ = reader.ReadString('\n')
		month2str = strings.Replace(month2str, "\r\n", "", -1)
		month, _ = strconv.Atoi(month2str)

		fmt.Println("Enter date:")
		fmt.Print("->")
		day2str, _ = reader.ReadString('\n')
		day2str = strings.Replace(day2str, "\r\n", "", -1)
		day, _ = strconv.Atoi(day2str)

		if isValidatedDate(year, month, day) {
			break
		}
		fmt.Println("The input date doesn't exist!")
	}

	fmt.Println("Enter period:")
	fmt.Print("->")
	period2str, _ = reader.ReadString('\n')
	period2str = strings.Replace(period2str, "\r\n", "", -1)
	period, _ = strconv.Atoi(period2str)

	userFile := year2str + ".txt"
	fout, err := os.Create(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fout.Close()

	fout.WriteString(dateGenerator(year, month, day, period, option))
	fmt.Print("Your note is generated.")
	return
}
