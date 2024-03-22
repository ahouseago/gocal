package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	prev := flag.Bool("prev", false, "Show previous month")
	next := flag.Bool("next", false, "Show next month")
	flag.Parse()
	if *prev && *next {
		panic("Must only specify one of --prev or --next")
	}

	now := time.Now()
	if *prev {
		now = now.AddDate(0, -1, 0)
	} else if *next {
		now = now.AddDate(0, 1, 0)
	}

	out := monthStr(now.Year(), now.Month())
	_, err := io.WriteString(os.Stdout, out)
	if err != nil {
		fmt.Println("Failed to write to stdout")
		fmt.Println(out)
	}
}

func monthStr(year int, month time.Month) string {
	now := time.Now()
	date := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	monthYear := date.Format("January 2006")
	var days [7]string
	for i := range 7 {
		days[i] = time.Weekday((i + 1) % 7).String()[:2]
	}
	daysString := strings.Join(days[:], " ")

	monthYear = strings.Repeat(" ", (len(daysString)-len(monthYear))/2) + monthYear

	firstOfMonth := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	calPadding := strings.Repeat("   ", int(firstOfMonth.Weekday()-1))

	datesStr := calPadding
	for i := 1; i <= lastOfMonth.Day(); i++ {
		if now.Year() == year && now.Month() == month && now.Day() == i {
			datesStr += fmt.Sprintf("(%2d)", i)
		} else if now.Year() == year && now.Month() == month && now.Day()+1 == i {
			datesStr += fmt.Sprintf("%2d", i)
		} else {
			datesStr += fmt.Sprintf("%3d", i)
		}
	}
	cal := ""
	for i := 0; i < len(datesStr); i += len(daysString) + 1 {
		cal += datesStr[i:min(len(datesStr), i+len(daysString)+1)] + "\n"
	}

	return fmt.Sprintf(" %s\n %s\n%s", monthYear, daysString, cal)
}
