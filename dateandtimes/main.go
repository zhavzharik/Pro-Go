package main

import (
	"fmt"
	"time"
)

//func PrintTime(label string, t *time.Time) {
//	Printfln("%s: Day: %v: Month: %v Year: %v", label, t.Day(), t.Month(), t.Year())
//}

func PrintTime(label string, t *time.Time) {
	//layout := "Day: 02 Month: Jan Year: 2006"
	//fmt.Println(label, t.Format(layout))
	fmt.Println(label, t.Format(time.RFC822Z))
}

func main() {
	current := time.Now()
	specific := time.Date(1995, time.June, 9, 0, 0, 0, 0, time.Local)
	unix := time.Unix(1433228090, 0)
	PrintTime("Current", &current)
	PrintTime("Specific", &specific)
	PrintTime("UNIX", &unix)
	fmt.Println()

	layout := "2006-Jan-02"
	dates := []string{
		"1995-Jun-09",
		"2015-Jun-02",
	}
	for _, d := range dates {
		time, err := time.Parse(layout, d)
		if err == nil {
			PrintTime("Parsed", &time)
		} else {
			Printfln("Error: %s", err.Error())
		}
	}
	fmt.Println()

	datesX := []string{
		"09 Jun 95 00:00 GMT",
		"02 Jun 15 00:00 GMT",
		"2015-Jun-02",
	}
	Printfln("time.RFC822 = %v", time.RFC822)
	for _, d := range datesX {
		time, err := time.Parse(time.RFC822, d)
		if err == nil {
			PrintTime("Parsed", &time)
		} else {
			Printfln("Error: %s", err.Error())
		}
	}
	fmt.Println()

	layoutX := "02 Jan 06 15:04"
	dateX := "09 Jun 95 19:30"
	london, lonerr := time.LoadLocation("Europe/London")
	newyork, nycerr := time.LoadLocation("America/New_York")
	local, _ := time.LoadLocation("Local")
	if lonerr == nil && nycerr == nil {
		nolocation, _ := time.Parse(layoutX, dateX)
		londonTime, _ := time.ParseInLocation(layoutX, dateX, london)
		newyorkTime, _ := time.ParseInLocation(layoutX, dateX, newyork)
		localTime, _ := time.ParseInLocation(layoutX, dateX, local)
		PrintTime("No location:", &nolocation)
		PrintTime("London:", &londonTime)
		PrintTime("New York:", &newyorkTime)
		PrintTime("Local:", &localTime)

	} else {
		fmt.Println(lonerr.Error(), nycerr.Error())
	}
	fmt.Println()

	londonFixed := time.FixedZone("BST", 1*60*60)
	newyorkFixed := time.FixedZone("EDT", -4*60*60)
	localFixed := time.FixedZone("Local", 3*60*60)

	nolocationF, _ := time.Parse(layoutX, dateX)
	londonTimeF, _ := time.ParseInLocation(layoutX, dateX, londonFixed)
	newyorkTimeF, _ := time.ParseInLocation(layoutX, dateX, newyorkFixed)
	localTimeF, _ := time.ParseInLocation(layoutX, dateX, localFixed)
	PrintTime("No location:", &nolocationF)
	PrintTime("London:", &londonTimeF)
	PrintTime("New York:", &newyorkTimeF)
	PrintTime("Local:", &localTimeF)
	fmt.Println()

	t, err := time.Parse(time.RFC822, "09 Jun 95 04:59 BST")
	if err == nil {
		Printfln("Example: %v,", t)
		Printfln("Time Now %v", time.Now().Format(time.RFC822))
		Printfln("After: %v", t.After(time.Now()))
		Printfln("Round: %v", t.Round(time.Hour))
		Printfln("Truncate: %v", t.Truncate(time.Hour))
	} else {
		fmt.Println(err.Error())
	}
	fmt.Println()

	t1, _ := time.Parse(time.RFC822Z, "09 Jun 95 04:59 +0100")
	t2, _ := time.Parse(time.RFC822Z, "08 Jun 95 23:59 -0400")
	Printfln("Equal Method: %v", t1.Equal(t2))
	Printfln("Equality Operator: %v", t1 == t2)
	fmt.Println()

	var d time.Duration = time.Hour + (30 * time.Minute)
	Printfln("Hours: %v", d.Hours())
	Printfln("Mins: %v", d.Minutes())
	Printfln("Seconds: %v", d.Seconds())
	Printfln("Millseconds: %v", d.Milliseconds())
	rounded := d.Round(time.Hour)
	Printfln("Rounded Hours: %v", rounded.Hours())
	Printfln("Rounded Mins: %v", rounded.Minutes())
	trunc := d.Truncate(time.Hour)
	Printfln("Truncated Hours: %v", trunc.Hours())
	Printfln("Rounded Mins: %v", trunc.Minutes())
	fmt.Println()

	toYears := func(d time.Duration) int {
		return int(d.Hours() / (24 * 365))
	}
	future := time.Date(2051, 0, 0, 0, 0, 0, 0, time.Local)
	past := time.Date(1965, 0, 0, 0, 0, 0, 0, time.Local)
	Printfln("Future: %v", toYears(time.Until(future)))
	Printfln("Past: %v", toYears(time.Since(past)))
	fmt.Println()

	d1, err1 := time.ParseDuration("1h30m")
	if err1 == nil {
		Printfln("Hours: %v", d1.Hours())
		Printfln("Minutes: %v", d1.Minutes())
		Printfln("Seconds: %v", d1.Seconds())
		Printfln("Milliseconds: %v", d1.Milliseconds())
	} else {
		fmt.Println(err.Error())
	}
}

//база данных часовых поясов
//https://www.iana.org/time-zones
//https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
